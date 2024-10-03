package lobby_data

type LobbyData struct {
	PlayerID       string         `json:"playerId"`
	Players        []LobbyPlayer  `json:"players"`
	ServerSettings ServerSettings `json:"serverSettings"`
}
