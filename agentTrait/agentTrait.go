package agentTrait

import (
	"hack-arena-2024-h2-go/packet/packets/agent_response"
	"hack-arena-2024-h2-go/packet/packets/game_end"
	"hack-arena-2024-h2-go/packet/packets/game_state"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
)

// IAgent defines the behavior of an AI agent interacting with the game
// by responding to game state updates and making decisions based on the current state.
type IAgent interface {
	// OnJoiningLobby is called when the agent joins a lobby, creating a new instance of the agent.
	// This method initializes the agent with the lobby's current state and other relevant details.
	//
	// Parameters:
	// - lobbyData: The initial state of the lobby when the agent joins.
	//   Contains information like player data, game settings, etc.
	//
	// Returns:
	// - A new instance of the agent.
	OnJoiningLobby(lobbyData *lobby_data.LobbyData) IAgent

	// OnLobbyDataChanged is called whenever there is a change in the lobby data.
	//
	// This method is triggered under various circumstances, such as:
	// - When a player joins or leaves the lobby.
	// - When server-side game settings are updated.
	//
	// Parameters:
	// - lobbyData: The updated state of the lobby, containing information
	//   like player details, game configurations, and other relevant data.
	//   This is the same data structure as the one provided when the agent
	//   first joined the lobby.
	//
	// Default Behavior:
	// By default, this method performs no action. To add custom behavior
	// when the lobby state changes, override this method in your implementation.
	OnLobbyDataChanged(lobbyData *lobby_data.LobbyData)

	// NextMove is called after each game tick, when new game state data is received from the server.
	// This method is responsible for determining the agent's next move based on the current game state.
	//
	// Parameters:
	// - gameState: The current state of the game, which includes all necessary information
	//   for the agent to decide its next action, such as the entire map with walls, tanks, bullets, zones, etc.
	//
	// Returns:
	// - AgentResponse: The action or decision made by the agent, which will be communicated back to the game server.
	NextMove(gameState *game_state.GameState) agent_response.AgentResponse

	// OnGameEnded is called when the game has concluded, providing the final game results.
	//
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
	OnGameEnded(gameEnd *game_end.GameEnd)
}
