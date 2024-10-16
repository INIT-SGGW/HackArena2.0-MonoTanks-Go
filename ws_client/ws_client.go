package ws_client

import (
	"context"
	"encoding/json"
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/handlers"
	"hack-arena-2024-h2-go/packet"
	"hack-arena-2024-h2-go/packet/packets/game_end"
	"hack-arena-2024-h2-go/packet/packets/game_state"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
	"log"
	"net/url"
	"sync"

	"hack-arena-2024-h2-go/packet/warning"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	readTask   *sync.WaitGroup
	writeTask  *sync.WaitGroup
	conn       *websocket.Conn
	tx         chan []byte
	agentMutex sync.Mutex
	agent      *agent.Agent
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
	case packet.ConnectionAccepted:
		fmt.Println("[System] 🎉 Connection accepted")
	case packet.ConnectionRejected:
		fmt.Printf("[System] 🚨 Connection rejected -> %s\n", p.Payload)
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

		client.agentMutex.Lock()
		err = handlers.HandlePrepareToGame(&client.agent, &lobbyData)
		client.agentMutex.Unlock()
		if err != nil {
			log.Printf("[System] 🚨 Error handling prepare to game: %v", err)
		}
	case packet.LobbyDeleted:
		fmt.Println("[System] 🚪 Lobby deleted")

	case packet.GameStarting:
		client.agentMutex.Lock()
		if client.agent != nil {
			handlers.HandleGameStarting(client.tx, client.agent)
		} else {
			log.Println("[System] 🚨 Received GameStarting but agent is not initialized")
		}
		client.agentMutex.Unlock()

	case packet.GameStarted:
		fmt.Println("[System] 🎲 Game started")

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

		client.agentMutex.Lock()
		if client.agent != nil {
			handlers.HandleNextMove(client.tx, client.agent, gameState)
		} else {
			log.Println("[System] 🚨 Received GameStatePacket but agent is not initialized")
		}
		client.agentMutex.Unlock()

	case packet.GameEndPacket:
		fmt.Println("[System] 🏁 Game ended")

		var gameEnd game_end.GameEnd
		payloadBytes, _ := json.Marshal(p.Payload)
		if err := json.Unmarshal(payloadBytes, &gameEnd); err != nil {
			log.Printf("[System] 🚨 Error unmarshalling GameEnd payload: %v", err)
			return
		}

		client.agentMutex.Lock()
		err := handlers.HandleGameEnded(client.agent, gameEnd)
		client.agentMutex.Unlock()
		if err != nil {
			log.Printf("[System] 🚨 Error handling game ended: %v", err)
		}

	// Warnings
	case packet.CustomWarning:
		message := p.Payload.(map[string]interface{})["message"].(string)
		client.agentMutex.Lock()
		handlers.HandleWarning(client.agent, warning.CustomWarning, &message)
		client.agentMutex.Unlock()
	case packet.PlayerAlreadyMadeActionWarning:
		client.agentMutex.Lock()
		handlers.HandleWarning(client.agent, warning.PlayerAlreadyMadeActionWarning, nil)
		client.agentMutex.Unlock()
	case packet.ActionIgnoredDueToDeadWarning:
		client.agentMutex.Lock()
		handlers.HandleWarning(client.agent, warning.ActionIgnoredDueToDeadWarning, nil)
		client.agentMutex.Unlock()
	case packet.SlowResponseWarning:
		client.agentMutex.Lock()
		handlers.HandleWarning(client.agent, warning.SlowResponseWarning, nil)
		client.agentMutex.Unlock()

	// Errors
	case packet.InvalidPacketTypeError:
		fmt.Println("[System] 🚨 Client sent an invalid packet type error")
	case packet.InvalidPacketUsageError:
		fmt.Println("[System] 🚨 Client used packet in invalid way")

	default:
		log.Printf("[System] 🚨 Unknown packet type -> %s", p.Type)
	}
}
