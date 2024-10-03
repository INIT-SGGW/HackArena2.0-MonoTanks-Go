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

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	readTask   *sync.WaitGroup
	writeTask  *sync.WaitGroup
	conn       *websocket.Conn
	tx         chan []byte
	agentMutex sync.Mutex
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

	agent := &agent.Agent{}
	client.readTask.Add(1)
	go client.createReaderTask(agent)

	client.writeTask.Add(1)
	go client.createWriterTask()

	return nil
}

func (client *WebSocketClient) Run(ctx context.Context) error {
	<-ctx.Done()
	close(client.tx)
	client.readTask.Wait()
	client.writeTask.Wait()
	return client.conn.Close()
}

func (client *WebSocketClient) constructURL(host string, port int, code string, nickname string) string {
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   "/",
	}
	q := u.Query()
	q.Set("nickname", nickname)
	q.Set("typeOfPacketType", "string")
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

func (client *WebSocketClient) createReaderTask(agent *agent.Agent) {
	defer client.readTask.Done()
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			log.Printf("[System] 🌋 WebSocket receive error -> %v", err)
			return
		}
		go client.processMessage(message, agent)
	}
}

func (client *WebSocketClient) processMessage(message []byte, agent *agent.Agent) {

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
		client.processTextMessage(p, agent)
	}
}

func (client *WebSocketClient) processTextMessage(p packet.Packet, agent *agent.Agent) {
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
		handlers.HandlePrepareToGame(&client.agentMutex, &agent, &lobbyData)
	case packet.LobbyDeleted:
		fmt.Println("[System] 🚪 Lobby deleted")
	case packet.GameStart:
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
			return
		}

		handlers.HandleNextMove(client.tx, &client.agentMutex, agent, gameState)
	case packet.GameEndPacket:
		fmt.Println("[System] 🏁 Game ended")
		if gameEnd, ok := p.Payload.(game_end.GameEnd); ok {
			handlers.HandleGameEnded(&client.agentMutex, agent, gameEnd)
		} else {
			log.Printf("[System] 🚨 Invalid game end type")
		}
	case packet.PlayerAlreadyMadeActionWarning:
		fmt.Println("[System] 🚨 Player already made action warning")
	case packet.MissingGameStateIdWarning:
		fmt.Println("[System] 🚨 Missing game state id warning")
	case packet.SlowResponseWarning:
		fmt.Println("[System] 🚨 Slow response warning")
	case packet.InvalidPacketTypeError:
		fmt.Println("[System] 🚨 Client sent an invalid packet type error")
	case packet.InvalidPacketUsageError:
		fmt.Println("[System] 🚨 Client used packet in invalid way")
	default:
		log.Printf("[System] 🚨 Unknown packet type -> %s", p.Type)
	}
}
