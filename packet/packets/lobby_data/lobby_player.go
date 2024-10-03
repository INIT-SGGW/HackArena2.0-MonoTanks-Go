package lobby_data

// LobbyPlayer represents a player in the game lobby.
type LobbyPlayer struct {
	// ID is a unique identifier for the player.
	ID string `json:"id"`

	// Nickname is the player's chosen nickname or alias.
	Nickname string `json:"nickname"`

	// Color represents the player's color, used in visual representation as a color for nickname and tank.
	Color uint64 `json:"color"`
}
