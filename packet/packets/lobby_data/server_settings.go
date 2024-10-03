package lobby_data

import (
	"encoding/json"
)

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

// NewServerSettings is a constructor for ServerSettings.
func NewServerSettings(gridDimension, numberOfPlayers, seed, broadcastInterval uint32, eagerBroadcast bool) *ServerSettings {
	return &ServerSettings{
		GridDimension:     gridDimension,
		NumberOfPlayers:   numberOfPlayers,
		Seed:              seed,
		BroadcastInterval: broadcastInterval,
		EagerBroadcast:    eagerBroadcast,
	}
}

// String returns the JSON representation of the ServerSettings.
func (ss *ServerSettings) String() string {
	data, err := json.Marshal(ss)
	if err != nil {
		return "Error: " + err.Error()
	}
	return string(data)
}
