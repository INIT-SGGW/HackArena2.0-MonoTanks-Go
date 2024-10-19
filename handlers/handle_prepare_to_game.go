package handlers

import (
	"encoding/json"
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
)

func HandlePrepareToGame(tx chan []byte, agentInstance **agent.Agent, lobbyData *lobby_data.LobbyData) error {
	if *agentInstance != nil {
		(*agentInstance).OnLobbyDataChanged(lobbyData)
	} else {
		fmt.Println("[System] 🤖 Creating agent")
		*agentInstance = agent.OnJoiningLobby(lobbyData)
		fmt.Println("[System] 🤖 Created agent")

		if lobbyData.ServerSettings.SandboxMode {
			fmt.Println("[System] 🛠️ Sandbox mode enabled")

			readyToReceiveGameState := packet.Packet{
				Type:    packet.ReadyToReceiveGameState,
				Payload: nil,
			}
			readyToReceiveGameStateBytes, err := json.Marshal(readyToReceiveGameState)
			if err != nil {
				return fmt.Errorf("error marshalling ReadyToReceiveGameState: %w", err)
			}
			tx <- readyToReceiveGameStateBytes
			fmt.Println("[System] 🎳 Ready to receive game state sent")

			gameStatusRequest := packet.Packet{
				Type:    packet.GameStatusRequest,
				Payload: nil,
			}
			gameStatusRequestBytes, err := json.Marshal(gameStatusRequest)
			if err != nil {
				return fmt.Errorf("error marshalling GameStatusRequest: %w", err)
			}
			tx <- gameStatusRequestBytes
		}
	}

	return nil
}
