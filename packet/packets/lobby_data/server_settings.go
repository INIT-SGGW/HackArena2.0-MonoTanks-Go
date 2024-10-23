package lobby_data

// ServerSettings represents the configuration settings for the server.
type ServerSettings struct {
	// GridDimension is the dimensions of the grid. The grid is a square with sides of equal length.
	GridDimension uint32 `json:"gridDimension"`

	// NumberOfPlayers is the number of players participating in the game. Minimum is 2. Maximum is 4.
	NumberOfPlayers uint32 `json:"numberOfPlayers"`

	// Seed is the seed value used for random number generation, ensuring consistency in results.
	// It is used to generate the grid and player starting positions.
	Seed uint32 `json:"seed"`

	// BroadcastInterval is the interval at which broadcast messages are sent to bots, in milliseconds.
	BroadcastInterval uint32 `json:"broadcastInterval"`

	// EagerBroadcast is a flag that determines whether broadcasts should happen
	// immediately after all players have made their action (true)
	// or at regular intervals (false).
	EagerBroadcast bool `json:"eagerBroadcast"`

	// SandboxMode is a flag that determines whether the game is in sandbox mode.
	// If sandbox mode is enabled, the game will not progress and will remain in the same state indefinitely.
	SandboxMode bool `json:"sandboxMode"`

	// Ticks is the number of ticks to run the game for. This is nil if sandbox mode is enabled.
	Ticks *int `json:"ticks"`

	// MatchName is the name of the match.
	MatchName *string `json:"matchName"`

	// Version is the version of the game running on the server.
	Version string `json:"version"`
}
