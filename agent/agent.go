package agent

import (
	"fmt"
	"math/rand"

	"hack-arena-2024-h2-go/packet/packets/agent_response"
	"hack-arena-2024-h2-go/packet/packets/game_end"
	"hack-arena-2024-h2-go/packet/packets/game_state"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
)

type Agent struct {
	MyID string
}

func OnJoiningLobby(lobbyData *lobby_data.LobbyData) *Agent {
	return &Agent{
		MyID: lobbyData.PlayerID,
	}
}

func (a *Agent) OnLobbyDataChanged(lobbyData *lobby_data.LobbyData) {
	// Implement the logic for handling lobby data changes
}

func (a *Agent) NextMove(gameState *game_state.GameState) *agent_response.AgentResponse {

	switch r := rand.Float32(); {
	case r < 0.33:
		direction := agent_response.Forward
		if rand.Intn(2) == 0 {
			direction = agent_response.Backward
		}
		return agent_response.NewTankMovement(direction)
	case r < 0.66:
		randomRotation := func() agent_response.Rotation {
			switch r := rand.Float32(); {
			case r < 0.33:
				return agent_response.Left
			case r < 0.66:
				return agent_response.Right
			default:
				return agent_response.Left // Default case to avoid nil return
			}
		}
		return agent_response.NewTankRotation(randomRotation(), randomRotation())
	default:
		return agent_response.NewTankShoot()
	}
}

func (a *Agent) OnGameEnded(gameEnd *game_end.GameEnd) {
	var winner game_end.GameEndPlayer
	for _, player := range gameEnd.Players {
		if player.Score > winner.Score {
			winner = player
		}
	}

	if winner.ID == a.MyID {
		fmt.Println("I won!")
	}

	for _, player := range gameEnd.Players {
		fmt.Printf("Player: %s - Score: %d\n", player.Nickname, player.Score)
	}
}
