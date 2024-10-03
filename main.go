package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"hack-arena-2024-h2-go/args"
	"hack-arena-2024-h2-go/ws_client"

	"github.com/urfave/cli/v2"
)

func main() {
	var args args.Args

	app := &cli.App{
		Name:    "agent",
		Usage:   "Configure the agent and connect to a server",
		Version: "1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Author Name",
				Email: "author@example.com",
			},
		},
		Flags: args.GetFlags(),
		Action: func(c *cli.Context) error {
			log.Printf("Nickname: %s", args.Nickname)
			log.Printf("Host: %s", args.Host)
			log.Printf("Port: %d", args.Port)
			log.Printf("Code: %s", args.Code)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[System] 🚀 Starting client...")
	websocketClient := ws_client.NewWebSocketClient()
	err = websocketClient.Connect(args.Host, int(args.Port), args.Code, args.Nickname)
	if err != nil {
		log.Fatalf("[System] 🌋 Error connecting to the server -> %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = websocketClient.Run(ctx)
	if err != nil {
		log.Fatalf("[System] 🌋 Error running WebSocket client -> %v", err)
	}
}
