package game_end

import (
	"encoding/json"
)

// GameEnd represents the end state of a game.
type GameEnd struct {
	// Players is the list of players at the end of the game.
	Players []GameEndPlayer `json:"players"`
}

// NewGameEnd is a constructor for GameEnd.
func NewGameEnd(players []GameEndPlayer) *GameEnd {
	return &GameEnd{
		Players: players,
	}
}

// String returns the JSON representation of the GameEnd.
func (ge *GameEnd) String() string {
	data, err := json.Marshal(ge)
	if err != nil {
		return "Error: " + err.Error()
	}
	return string(data)
}
