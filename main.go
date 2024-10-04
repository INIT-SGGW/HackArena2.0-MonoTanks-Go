package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"hack-arena-2024-h2-go/args"
	"hack-arena-2024-h2-go/ws_client"
)

func main() {
	app := args.NewCLIApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// Get the parsed arguments
	parsedArgs, ok := app.Metadata["args"].(*args.Args)
	if !ok || parsedArgs == nil {
		log.Fatal("[System] 🌋 Error: Failed to retrieve parsed arguments")
	}

	fmt.Println("[System] 🚀 Starting client...")
	websocketClient := ws_client.NewWebSocketClient()
	err = websocketClient.Connect(parsedArgs.Host, int(parsedArgs.Port), parsedArgs.Code, parsedArgs.Nickname)
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
