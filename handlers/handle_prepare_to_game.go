package handlers

import (
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
)

func HandlePrepareToGame(agentInstance **agent.Agent, lobbyData *lobby_data.LobbyData) error {

	fmt.Println("[System] 🤖 Creating agent")
	*agentInstance = agent.OnJoiningLobby(lobbyData)
	fmt.Println("[System] 🤖 Created agent")

	return nil
}
