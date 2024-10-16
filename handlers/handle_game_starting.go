package handlers

import (
	"encoding/json"
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet/packets"
)

func HandleGameStarting(tx chan []byte, agent *agent.Agent) error {

	if agent == nil {
		return fmt.Errorf("agent not initialized")
	}

	agent.OnGameStarting()

	// Convert agent response to packet
	responsePacket := packets.ReadyToReceiveGameState{}
	responseString, err := json.Marshal(responsePacket.ToPacket())
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
