package agentTrait

import (
	"hack-arena-2024-h2-go/packet/packets/agent_response"
	"hack-arena-2024-h2-go/packet/packets/game_end"
	"hack-arena-2024-h2-go/packet/packets/game_state"
	"hack-arena-2024-h2-go/packet/packets/lobby_data"
)

type IAgent interface {
	OnLobbyDataChanged(lobbyData *lobby_data.LobbyData)
	NextMove(gameState *game_state.GameState) agent_response.AgentResponse
	OnGameEnded(gameEnd *game_end.GameEnd)
}
