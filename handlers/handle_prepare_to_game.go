package handlers

import (
	"encoding/json"
	"fmt"
	"hackarena2-0-mono-tanks-go/bot"
	"hackarena2-0-mono-tanks-go/packet"
	"hackarena2-0-mono-tanks-go/packet/packets/lobby_data"
)

func HandlePrepareToGame(tx chan []byte, botInstance **bot.Bot, lobbyData *lobby_data.LobbyData) error {
	if *botInstance != nil {
		(*botInstance).OnLobbyDataChanged(lobbyData)
	} else {
		fmt.Println("[System] 🤖 Creating bot")
		*botInstance = bot.OnJoiningLobby(lobbyData)
		fmt.Println("[System] 🤖 Created bot")

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
