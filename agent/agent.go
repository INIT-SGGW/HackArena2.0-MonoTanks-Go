package agent

import (
	"fmt"
	"math/rand"

	"hack-arena-2024-h2-go/packet/packets/agent_response"
	"hack-arena-2024-h2-go/packet/packets/agent_response/ability"
	"hack-arena-2024-h2-go/packet/packets/agent_response/movement"
	"hack-arena-2024-h2-go/packet/packets/agent_response/rotation"
	"hack-arena-2024-h2-go/packet/packets/game_end"
	"hack-arena-2024-h2-go/packet/packets/game_state"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
	"hack-arena-2024-h2-go/packet/warning"
)

// Agent represents an AI player in the game.
type Agent struct {
	MyID string
}

// OnJoiningLobby is called when the agent joins a lobby, creating a new instance of the agent.
// This method initializes the agent with the lobby's current state and other relevant details.
//
// Parameters:
//   - lobbyData: The initial state of the lobby when the agent joins.
//     Contains information like player data, game settings, etc.
//
// Returns:
// - A new instance of the agent.
func OnJoiningLobby(lobbyData *lobby_data.LobbyData) *Agent {
	return &Agent{
		MyID: lobbyData.PlayerID,
	}
}

// OnLobbyDataChanged is called whenever there is a change in the lobby data.
// This method is triggered under various circumstances, such as:
// - When a player joins or leaves the lobby.
// - When server-side game settings are updated.
//
// Parameters:
//   - lobbyData: The updated state of the lobby, containing information
//     like player details, game configurations, and other relevant data.
//     This is the same data structure as the one provided when the agent
//     first joined the lobby.
//
// Default Behavior:
// By default, this method performs no action. To add custom behavior
// when the lobby state changes, override this method in your implementation.
func (a *Agent) OnLobbyDataChanged(lobbyData *lobby_data.LobbyData) {
	// Implement the logic for handling lobby data changes
}

// NextMove is called after each game tick, when new game state data is received from the server.
// This method is responsible for determining the agent's next move based on the current game state.
//
// Parameters:
//   - gameState: The current state of the game, which includes all necessary information
//     for the agent to decide its next action, such as the entire map with walls, tanks, bullets, zones, etc.
//
// Returns:
// - AgentResponse: The action or decision made by the agent, which will be communicated back to the game server.
func (a *Agent) NextMove(gameState *game_state.GameState) *agent_response.AgentResponse {

	// Find my tank
	var myTank *game_state.Tank
	for _, tank := range gameState.Tanks {
		if tank.OwnerID == a.MyID {
			myTank = &tank
			break
		}
	}

	// If my tank is not found, it is dead
	if myTank == nil {
		return agent_response.NewPass()
	}

	switch r := rand.Float32(); {
	case r < 0.25:
		// Move the tank
		direction := movement.Forward
		if rand.Intn(2) == 1 {
			direction = movement.Backward
		}
		return agent_response.NewMovement(direction)
	case r < 0.50:
		// Rotate the tank and/or turret
		randomRotation := func() string {
			switch rand.Intn(3) {
			case 0:
				return rotation.Left
			case 1:
				return rotation.Right
			default:
				return ""
			}
		}
		return agent_response.NewRotation(randomRotation(), randomRotation())
	case r < 0.75:
		// Use ability
		abilities := []string{
			ability.FireBullet,
			ability.FireDoubleBullet,
			ability.UseLaser,
			ability.UseRadar,
			ability.DropMine,
		}
		abilityType := abilities[rand.Intn(len(abilities))]
		return agent_response.NewAbilityUse(abilityType)
	default:
		// Pass
		return agent_response.NewPass()
	}
}

// OnWarningReceived is called when a warning is received from the server.
// Please remember that if your agent is stuck processing a warning,
// the next move won't be called and vice versa.
//
// Parameters:
// - warning: The warning received from the server.
func (a *Agent) OnWarningReceived(warn warning.Warning, message *string) {
	switch warn {
	case warning.CustomWarning:
		msg := "No message"
		if message != nil {
			msg = *message
		}
		fmt.Printf("[System] ⚠️ Custom Warning: %s\n", msg)
	case warning.PlayerAlreadyMadeActionWarning:
		fmt.Println("[System] ⚠️ Player already made action warning")
	case warning.ActionIgnoredDueToDeadWarning:
		fmt.Println("[System] ⚠️ Action ignored due to dead warning")
	case warning.SlowResponseWarning:
		fmt.Println("[System] ⚠️ Slow response warning")
	}
}

// OnGameEnded is called when the game has concluded, providing the final game results.
// This method is triggered when the game ends, which is when a defined number of ticks in LobbyData has passed.
//
// Parameters:
// - gameEnd: The final state of the game, containing players' scores.
//
// Default Behavior:
// By default, this method performs no action. You can override it to implement any post-game behavior,
// such as logging, updating agent strategies, or other clean-up tasks.
//
// Notes:
// - This method is optional to override, but it can be useful for handling game result analysis and logging.
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
