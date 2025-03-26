package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hashup",
		Usage: "HashUp Application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "extensions",
				Usage: "Comma separated list of file extensions to include in search results",
				Value: "",
			},
			&cli.IntFlag{
				Name:  "limit",
				Usage: "Maximum number of results to return",
				Value: 100,
			},
			&cli.IntFlag{
				Name:  "height",
				Usage: "Window height",
				Value: 800,
			},
			&cli.IntFlag{
				Name:  "width",
				Usage: "Window width",
				Value: 1300,
			},
		},
		Action: func(c *cli.Context) error {
			apiPort, err := randomPort()
			if err != nil {
				return fmt.Errorf("Error getting random port: %w", err)
			}

			fmt.Printf("API listening on http://localhost:%d\n", apiPort)
			go func() {
				if err := serveAPI(fmt.Sprintf("127.0.0.1:%d", apiPort), c); err != nil {
					panic(err)
				}
			}()

			runApp(apiPort, c)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
