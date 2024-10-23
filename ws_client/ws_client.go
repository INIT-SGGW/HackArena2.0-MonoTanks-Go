package ws_client

import (
	"context"
	"encoding/json"
	"fmt"
	"hackarena2-0-mono-tanks-go/bot"
	"hackarena2-0-mono-tanks-go/handlers"
	"hackarena2-0-mono-tanks-go/packet"
	"hackarena2-0-mono-tanks-go/packet/packets/game_end"
	"hackarena2-0-mono-tanks-go/packet/packets/game_state"
	"hackarena2-0-mono-tanks-go/packet/packets/lobby_data"
	"log"
	"net/url"
	"sync"
	"time"

	"hackarena2-0-mono-tanks-go/packet/warning"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	readTask    *sync.WaitGroup
	writeTask   *sync.WaitGroup
	conn        *websocket.Conn
	tx          chan []byte
	botMutex    sync.Mutex
	botInstance *bot.Bot
}

func NewWebSocketClient() *WebSocketClient {
	return &WebSocketClient{
		readTask:  &sync.WaitGroup{},
		writeTask: &sync.WaitGroup{},
		tx:        make(chan []byte, 100),
	}
}

func (client *WebSocketClient) Connect(host string, port int, code string, nickname string) error {
	url := client.constructURL(host, port, code, nickname)

	fmt.Printf("[System] 📞 Connecting to the server: %s\n", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("[System] 🌋 WebSocket connection error -> %v", err)
	}
	fmt.Println("[System] 🌟 Successfully connected to the server")

	client.conn = conn

	client.readTask.Add(1)
	go client.createReaderTask()

	client.writeTask.Add(1)
	go client.createWriterTask()

	return nil
}

func (client *WebSocketClient) Run(ctx context.Context) error {
	defer func() {
		if err := client.conn.Close(); err != nil {
			log.Printf("[System] 🚨 Error closing WebSocket connection: %v", err)
		}
		fmt.Println("[System] 👋 Connection closed")
	}()

	done := make(chan struct{})
	go func() {
		defer close(done)
		client.readTask.Wait()
	}()

	select {
	case <-ctx.Done():
		log.Println("[System] 🛑 Context cancelled, closing connection...")
		return client.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	case <-done:
		return nil
	}
}

func (client *WebSocketClient) constructURL(host string, port int, code string, nickname string) string {
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   "/",
	}
	q := u.Query()
	q.Set("nickname", nickname)
	q.Set("enumSerializationFormat", "string")
	q.Set("playerType", "hackathonBot")
	if code != "" {
		q.Set("joinCode", code)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (client *WebSocketClient) createWriterTask() {
	defer client.writeTask.Done()
	for message := range client.tx {
		if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("[System] 🌋 WebSocket send error -> %v", err)
		}
	}
}

func (client *WebSocketClient) createReaderTask() {
	defer client.readTask.Done()
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[System] 🌋 WebSocket unexpected close error: %v", err)
			} else {
				log.Printf("[System] 👋 WebSocket connection closed: %v", err)
			}
			return
		}
		go client.processMessage(message)
	}
}

func (client *WebSocketClient) processMessage(message []byte) {

	var p packet.Packet
	if err := json.Unmarshal(message, &p); err != nil {
		log.Printf("[System] 🚨 Error processing text message -> %v", err)
		log.Printf("[System] 🚨 Text Message -> %s", message)
		return
	}

	switch p.Type {
	case packet.Ping:
		client.tx <- []byte(`{"type":"pong"}`)
	case packet.Pong:
		fmt.Println("[System] 🏓 Received Pong")
	default:
		client.processTextMessage(p)
	}
}

func (client *WebSocketClient) processTextMessage(p packet.Packet) {
	switch p.Type {

	case packet.ConnectionRejected:
		fmt.Printf("[System] 🚨 Connection rejected -> %s\n", p.Payload)
	case packet.ConnectionAccepted:
		fmt.Println("[System] 🎉 Connection accepted")

		lobbyDataRequest := packet.Packet{
			Type:    packet.LobbyDataRequest,
			Payload: nil,
		}

		lobbyDataRequestJson, err := json.Marshal(lobbyDataRequest)
		if err != nil {
			log.Printf("[System] 🚨 Error marshalling LobbyDataRequest: %v", err)
			return
		}
		client.tx <- lobbyDataRequestJson

	case packet.LobbyDataPacket:
		fmt.Println("[System] 🎳 Lobby data received")
		var lobbyData lobby_data.LobbyData
		payloadBytes, err := json.Marshal(p.Payload)
		if err != nil {
			log.Printf("[System] 🚨 Error marshalling payload: %v", err)
			return
		}
		err = json.Unmarshal(payloadBytes, &lobbyData)
		if err != nil {
			log.Printf("[System] 🚨 Error unmarshalling payload into LobbyData: %v", err)
			return
		}

		client.botMutex.Lock()
		err = handlers.HandlePrepareToGame(client.tx, &client.botInstance, &lobbyData)
		client.botMutex.Unlock()
		if err != nil {
			log.Printf("[System] 🚨 Error handling prepare to game: %v", err)
		}

	case packet.GameNotStarted:
		fmt.Println("[System] 🎲 Game not started")

	case packet.GameStarting:
		fmt.Println("[System] 🎲 Game starting")

		// Wait until bot is not None
		for client.botInstance == nil {
			time.Sleep(100 * time.Millisecond)
		}

		readyToReceiveGameState := packet.Packet{
			Type:    packet.ReadyToReceiveGameState,
			Payload: nil,
		}
		readyToReceiveGameStateJson, err := json.Marshal(readyToReceiveGameState)
		if err != nil {
			log.Printf("[System] 🚨 Error marshalling ReadyToReceiveGameState: %v", err)
			return
		}
		client.tx <- readyToReceiveGameStateJson

	case packet.GameStarted:
		fmt.Println("[System] 🎲 Game started")

	case packet.GameInProgress:
		fmt.Println("[System] 🎲 Game in progress")

	case packet.GameStatePacket:

		var gameState game_state.GameState
		payloadBytes, err := json.Marshal(p.Payload)
		if err != nil {
			log.Printf("[System] 🚨 Error marshalling payload: %v", err)
			return
		}
		err = json.Unmarshal(payloadBytes, &gameState)
		if err != nil {
			log.Printf("[System] 🚨 Error unmarshalling payload into GameState: %v", err)
			log.Printf("[System] 🚨 Text Message -> %s", p.Payload)
			return
		}

		client.botMutex.Lock()
		if client.botInstance != nil {
			handlers.HandleNextMove(client.tx, client.botInstance, gameState)
		} else {
			log.Println("[System] 🚨 Received GameStatePacket, but bot is not initialized")
		}
		client.botMutex.Unlock()

	case packet.GameEndedPacket:
		fmt.Println("[System] 🏁 Game ended")

		var gameEnd game_end.GameEnd
		payloadBytes, _ := json.Marshal(p.Payload)
		if err := json.Unmarshal(payloadBytes, &gameEnd); err != nil {
			log.Printf("[System] 🚨 Error unmarshalling GameEnd payload: %v", err)
			return
		}

		client.botMutex.Lock()
		err := handlers.HandleGameEnded(client.botInstance, gameEnd)
		client.botMutex.Unlock()
		if err != nil {
			log.Printf("[System] 🚨 Error handling game ended: %v", err)
		}

	// Warnings
	case packet.CustomWarning:
		message := p.Payload.(map[string]interface{})["message"].(string)
		client.botMutex.Lock()
		handlers.HandleWarning(client.botInstance, warning.CustomWarning, &message)
		client.botMutex.Unlock()
	case packet.PlayerAlreadyMadeActionWarning:
		client.botMutex.Lock()
		handlers.HandleWarning(client.botInstance, warning.PlayerAlreadyMadeActionWarning, nil)
		client.botMutex.Unlock()
	case packet.ActionIgnoredDueToDeadWarning:
		client.botMutex.Lock()
		handlers.HandleWarning(client.botInstance, warning.ActionIgnoredDueToDeadWarning, nil)
		client.botMutex.Unlock()
	case packet.SlowResponseWarning:
		client.botMutex.Lock()
		handlers.HandleWarning(client.botInstance, warning.SlowResponseWarning, nil)
		client.botMutex.Unlock()

	// Errors
	case packet.InvalidPacketTypeError:
		fmt.Println("[System] 🚨 Websocket client sent an invalid packet type error")
	case packet.InvalidPacketUsageError:
		fmt.Println("[System] 🚨 Websocket client used packet in invalid way")

	default:
		log.Printf("[System] 🚨 Unknown packet type -> %s", p.Type)
	}
}
