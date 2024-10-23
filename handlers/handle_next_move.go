package handlers

import (
	"encoding/json"
	"fmt"
	"hackarena2-0-mono-tanks-go/bot"
	"hackarena2-0-mono-tanks-go/packet/packets/game_state"
)

func HandleNextMove(tx chan []byte, botInstance *bot.Bot, gameState game_state.GameState) error {
	gameStateID := gameState.ID

	if botInstance == nil {
		return fmt.Errorf("bot not initialized")
	}

	botResponse := botInstance.NextMove(&gameState)

	// Convert bot response to packet
	responsePacket := botResponse.ToPacket(gameStateID)
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
