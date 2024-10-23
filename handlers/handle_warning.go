package handlers

import (
	"fmt"
	"hackarena2-0-mono-tanks-go/bot"
	"hackarena2-0-mono-tanks-go/packet/warning"
)

func HandleWarning(botInstance *bot.Bot, warn warning.Warning, message *string) error {

	if botInstance == nil {
		return fmt.Errorf("bot not initialized")
	}

	botInstance.OnWarningReceived(warn, message)

	return nil
}
