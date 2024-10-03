package game_end

// GameEnd represents the end state of a game.
type GameEnd struct {
	// Players is the list of players at the end of the game.
	Players []GameEndPlayer `json:"players"`
}
