package handlers

import (
	"fmt"
	"hackarena2-0-mono-tanks-go/bot"
	"hackarena2-0-mono-tanks-go/packet/packets/game_end"
)

func HandleGameEnded(botInstance *bot.Bot, gameEnd game_end.GameEnd) error {
	if botInstance == nil {
		return fmt.Errorf("bot not initialized")
	}

	botInstance.OnGameEnded(&gameEnd)
	return nil
}
