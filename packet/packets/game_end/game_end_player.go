package game_end

import (
	"encoding/json"
)

// GameEndPlayer represents a player in the game.
type GameEndPlayer struct {
	// ID is a unique identifier for the player.
	ID string `json:"id"`

	// Nickname is the player's chosen nickname or alias.
	Nickname string `json:"nickname"`

	// Color represents the player's color, used in visual representation as a color for nickname and tank.
	Color uint64 `json:"color"`

	// Score is the player's final score in the game.
	Score uint64 `json:"score"`
}

// NewGameEndPlayer is a constructor for GameEndPlayer.
func NewGameEndPlayer(id, nickname string, color, score uint64) *GameEndPlayer {
	return &GameEndPlayer{
		ID:       id,
		Nickname: nickname,
		Color:    color,
		Score:    score,
	}
}

// String returns the JSON representation of the GameEndPlayer.
func (gep *GameEndPlayer) String() string {
	data, err := json.Marshal(gep)
	if err != nil {
		return "Error: " + err.Error()
	}
	return string(data)
}
