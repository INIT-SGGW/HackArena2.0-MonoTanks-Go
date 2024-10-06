package game_end

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

	// Kills is the number of kills the player achieved in the game.
	Kills uint64 `json:"kills"`
}
