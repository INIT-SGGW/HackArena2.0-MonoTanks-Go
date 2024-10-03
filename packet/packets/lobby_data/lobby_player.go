package lobby_data

import (
	"encoding/json"
)

// LobbyPlayer represents a player in the game lobby.
type LobbyPlayer struct {
	// ID is a unique identifier for the player.
	ID string `json:"id"`

	// Nickname is the player's chosen nickname or alias.
	Nickname string `json:"nickname"`

	// Color represents the player's color, used in visual representation as a color for nickname and tank.
	Color uint64 `json:"color"`
}

// NewLobbyPlayer is a constructor for LobbyPlayer.
func NewLobbyPlayer(id, nickname string, color uint64) *LobbyPlayer {
	return &LobbyPlayer{
		ID:       id,
		Nickname: nickname,
		Color:    color,
	}
}

// String returns the JSON representation of the LobbyPlayer.
func (lp *LobbyPlayer) String() string {
	data, err := json.Marshal(lp)
	if err != nil {
		return "Error: " + err.Error()
	}
	return string(data)
}
