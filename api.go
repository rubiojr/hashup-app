package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rubiojr/hashup-app/internal/config"
	"github.com/rubiojr/hashup-app/internal/templates"
	"github.com/urfave/cli/v2"
)

var fileStatsJSON = `{"total_count": 0, "total_size_human": "0B", "extensions": []}`
var once sync.Once

func getFileStats() string {
	once.Do(func() {
		go func() {
			dbConn, err := dbConn("")
			if err != nil {
				log.Printf("Error connecting to database: %v", err)
				return
			}
			defer dbConn.Close()

			for {
				now := time.Now()
				stats, err := fileStats(dbConn, "file_size", true, "")
				if err != nil {
					log.Printf("Error getting file stats: %v", err)
					continue
				}

				fileStatsJSON, err = jsonStats(stats, "", 10)
				if err != nil {
					log.Printf("Error converting file stats to JSON: %v", err)
				}
				elapsed := time.Since(now)
				log.Printf("fileStats took %s", elapsed)
				time.Sleep(time.Second * 120)
			}
		}()
	})

	return fileStatsJSON
}

func healthHandlers() {
	http.HandleFunc("/health/nats", func(w http.ResponseWriter, r *http.Request) {
		config, err := config.LoadDefaultConfig()
		if err != nil {
			http.Error(w, fmt.Errorf("Error loading config: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		_, err = nats.Connect(config.Main.NatsServerURL)
		if err != nil {
			http.Error(w, fmt.Errorf("Error connecting to NATS server: %w", err).Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "NATS server is healthy")
	})
}

func fileStatsHandler() {
	http.HandleFunc("/stats/files", func(w http.ResponseWriter, r *http.Request) {
		var stats string
		for stats = getFileStats(); stats == ""; stats = getFileStats() {
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Fprintf(w, "%s", stats)
	})
}

func natsStreamInfoHandler() error {
	http.HandleFunc("/stats/nats/stream/info", func(w http.ResponseWriter, r *http.Request) {
		config, err := config.LoadDefaultConfig()
		if err != nil {
			http.Error(w, fmt.Errorf("failed to load HashUp config: %w", err).Error(), http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		nc, err := nats.Connect(config.Main.NatsServerURL)
		if err != nil {
			http.Error(w, fmt.Errorf("Error connecting to NATS server: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		js, _ := jetstream.New(nc)
		if err != nil {
			http.Error(w, fmt.Errorf("Error creating JetStream management interface: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		stream, err := js.Stream(ctx, config.Main.NatsStream)
		if err != nil {
			http.Error(w, fmt.Errorf("Error creating stream: %w", err).Error(), http.StatusInternalServerError)
			return
		}
		streamInfo, err := stream.Info(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clusterInfo := streamInfo.Cluster

		info := struct {
			StreamName    string `json:"stream_name"`
			ClusterName   string `json:"cluster_name"`
			ClusterLeader string `json:"cluster_leader"`
			Messages      int64  `json:"messages"`
			Bytes         int64  `json:"bytes"`
			ConsumerCount int64  `json:"consumer_count"`
		}{
			StreamName:    streamInfo.Config.Name,
			ClusterName:   clusterInfo.Name,
			ClusterLeader: clusterInfo.Leader,
			Messages:      int64(streamInfo.State.Msgs),
			Bytes:         int64(streamInfo.State.Bytes),
			ConsumerCount: int64(streamInfo.State.Consumers),
		}

		jsonState, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", jsonState)
	})
	return nil
}

func searchHandler(c *cli.Context) {
	extensions := strings.Split(c.String("extensions"), ",")

	// TODO
	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		_, err := config.LoadDefaultConfig()
		if err != nil {
			templates.NotConfigured().Render(r.Context(), w)
			return
		}

		_, err = getDBPath()
		if err != nil {
			templates.NotConfigured().Render(r.Context(), w)
			return
		}

		templates.Index().Render(r.Context(), w)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		dbPath, err := getDBPath()
		if err != nil {
			http.Error(w, fmt.Errorf("failed to get HashUp database path: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			http.Error(w, fmt.Errorf("failed to open HashUp database: %w", err).Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		query := r.URL.Query().Get("q")
		results, err := dbSearch(db, query, extensions, 5)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		templates.Results(results).Render(r.Context(), w)
	})
}

func serveAPI(addr string, c *cli.Context) error {
	getFileStats()
	searchHandler(c)
	err := natsStreamInfoHandler()
	if err != nil {
		return fmt.Errorf("Error handling NATS stream info: %w", err)
	}
	healthHandlers()
	fileStatsHandler()

	server := &http.Server{Addr: addr}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	<-c.Done()
	return nil
}
