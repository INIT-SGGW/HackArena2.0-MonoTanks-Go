package handlers

import (
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet/packets/game_end"
)

func HandleGameEnded(agentInstance *agent.Agent, gameEnd game_end.GameEnd) error {
	if agentInstance == nil {
		return fmt.Errorf("agent not initialized")
	}

	agentInstance.OnGameEnded(&gameEnd)
	return nil
}
