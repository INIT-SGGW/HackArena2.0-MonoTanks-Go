package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"hackarena2-0-mono-tanks-go/args"
	"hackarena2-0-mono-tanks-go/ws_client"
)

func main() {
	app := args.NewCLIApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	parsedArgs, ok := app.Metadata["args"].(*args.Args)
	if !ok || parsedArgs == nil {
		log.Fatal("[System] 🌋 Error: Failed to retrieve parsed arguments")
	}

	fmt.Println("[System] 🚀 Starting bot...")
	if err := startWebSocketClient(parsedArgs); err != nil {
		log.Printf("[System] 🌋 Error: %v", err)
	}
	fmt.Println("[System] 🏁 Bot stopped")
}

func startWebSocketClient(parsedArgs *args.Args) error {
	websocketClient := ws_client.NewWebSocketClient()
	err := websocketClient.Connect(parsedArgs.Host, int(parsedArgs.Port), parsedArgs.Code, parsedArgs.Nickname)
	if err != nil {
		return fmt.Errorf("connecting to the server: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		fmt.Println("\n[System] 🛑 Interrupt received, shutting down...")
		cancel()
	}()

	if err := websocketClient.Run(ctx); err != nil {
		return fmt.Errorf("running WebSocket client: %w", err)
	}
	return nil
}
