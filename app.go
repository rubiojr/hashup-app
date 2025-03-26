package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rubiojr/gapp/pkg/glance"
	webview "github.com/rubiojr/hashup-app/_webview"
	"github.com/urfave/cli/v2"

	_ "embed"
)

//go:embed glance.yml
var glanceConfig string

func runApp(apiPort int, c *cli.Context) error {
	port, err := randomPort()
	if err != nil {
		return err
	}
	opts := []glance.Option{
		glance.WithServerPort(uint16(port)),
		glance.WithLogger(log.New(os.Stdout, "", log.LstdFlags)),
		glance.WithHost("127.0.0.1"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Replaces custom-api widget ports in config
	glanceConfig = strings.Replace(glanceConfig, "@@API_PORT@@", fmt.Sprintf("%d", apiPort), -1)
	err = glance.ServeApp(ctx, []byte(glanceConfig), opts...)
	if err != nil {
		panic(err)
	}

	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("HashUp")
	w.SetSize(c.Int("width"), c.Int("height"), webview.HintNone)
	w.Navigate(fmt.Sprintf("http://%s:%d", "127.0.0.1", port))

	w.Run()
	return nil
}
