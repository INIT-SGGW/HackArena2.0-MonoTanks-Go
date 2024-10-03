package lobby_data

import (
	"encoding/json"
	"fmt"
)

type LobbyData struct {
	PlayerID       string         `json:"playerId"`
	Players        []LobbyPlayer  `json:"players"`
	ServerSettings ServerSettings `json:"serverSettings"`
}

func NewLobbyData(playerID string, players []LobbyPlayer, serverSettings ServerSettings) *LobbyData {
	return &LobbyData{
		PlayerID:       playerID,
		Players:        players,
		ServerSettings: serverSettings,
	}
}

func (ld *LobbyData) String() string {
	data, err := json.Marshal(ld)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(data)
}
