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

	// BroadcastInterval is the interval at which broadcast messages are sent to clients, in milliseconds.
	BroadcastInterval uint32 `json:"broadcastInterval"`

	// EagerBroadcast is a flag that determines whether broadcasts should happen
	// immediately after all players have made their action (true)
	// or at regular intervals (false).
	EagerBroadcast bool `json:"eagerBroadcast"`
}
