package handlers

import (
	"fmt"
	"sync"

	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet/packets/game_end"
)

// HandleGameEnded handles the end of the game.
func HandleGameEnded(agentMutex *sync.Mutex, agentInstance *agent.Agent, gameEnd game_end.GameEnd) error {
	agentMutex.Lock()
	defer agentMutex.Unlock()

	if agentInstance == nil {
		return fmt.Errorf("agent not initialized")
	}

	agentInstance.OnGameEnded(&gameEnd)
	return nil
}
