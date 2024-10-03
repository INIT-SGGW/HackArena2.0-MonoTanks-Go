package handlers

import (
	"encoding/json"
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet/packets/agent_response"
	"hack-arena-2024-h2-go/packet/packets/game_state"
	"sync"
)

func HandleNextMove(tx chan []byte, agentMutex *sync.Mutex, agent *agent.Agent, gameState game_state.GameState) error {
	gameStateID := gameState.ID

	// Lock the agent and get the next move
	var agentResponse *agent_response.AgentResponse
	agentMutex.Lock()
	defer agentMutex.Unlock()

	if agent == nil {
		return fmt.Errorf("agent not initialized")
	}

	agentResponse = agent.NextMove(&gameState)

	// Convert agent response to packet
	responsePacket := agentResponse.ToPacket(gameStateID)
	responseString, err := json.Marshal(responsePacket)
	if err != nil {
		return fmt.Errorf("failed to serialize response packet: %v", err)
	}

	// Send the response
	select {
	case tx <- responseString:
		return nil
	default:
		return fmt.Errorf("failed to send message")
	}
}
