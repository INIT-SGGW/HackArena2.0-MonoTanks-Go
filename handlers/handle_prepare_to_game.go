package handlers

import (
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
	"sync"
)

func HandlePrepareToGame(agentMutex *sync.Mutex, agentInstance **agent.Agent, lobbyData *lobby_data.LobbyData) error {
	agentMutex.Lock()
	defer agentMutex.Unlock()

	if *agentInstance != nil {
		(*agentInstance).OnLobbyDataChanged(lobbyData)
	} else {
		*agentInstance = agent.OnJoiningLobby(lobbyData)
		fmt.Println("[System] 🤖 Created agent")
	}

	return nil
}
