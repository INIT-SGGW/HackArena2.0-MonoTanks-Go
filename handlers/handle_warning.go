package handlers

import (
	"fmt"
	"hack-arena-2024-h2-go/agent"
	"hack-arena-2024-h2-go/packet/warning"
)

func HandleWarning(agent *agent.Agent, warn warning.Warning, message *string) error {

	if agent == nil {
		return fmt.Errorf("agent not initialized")
	}

	agent.OnWarningReceived(warn, message)

	return nil
}
