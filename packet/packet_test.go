package packet_test

import (
	"encoding/json"
	"hackarena2-0-mono-tanks-go/packet"
	"hackarena2-0-mono-tanks-go/packet/packets/game_end"
	"hackarena2-0-mono-tanks-go/packet/packets/lobby_data"
	"testing"
)

func TestPacketUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected packet.PacketType
	}{
		{
			input:    `{"type": "connectionAccepted"}`,
			expected: packet.ConnectionAccepted,
		},
		{
			input:    `{"type": "ping"}`,
			expected: packet.Ping,
		},
		{
			input:    `{"type": "pong"}`,
			expected: packet.Pong,
		},
		{
			input:    `{"type": "connectionRejected"}`,
			expected: packet.ConnectionRejected,
		},
		{
			input:    `{"type": "gameStarting"}`,
			expected: packet.GameStarting,
		},
		{
			input:    `{"type": "readyToReceiveGameState"}`,
			expected: packet.ReadyToReceiveGameState,
		},
		{
			input:    `{"type": "gameStarted"}`,
			expected: packet.GameStarted,
		},
		{
			input:    `{"type": "gameEnded"}`,
			expected: packet.GameEndedPacket,
		},
		{
			input:    `{"type": "playerAlreadyMadeActionWarning"}`,
			expected: packet.PlayerAlreadyMadeActionWarning,
		},
		{
			input:    `{"type": "slowResponseWarning"}`,
			expected: packet.SlowResponseWarning,
		},
		{
			input:    `{"type": "actionIgnoredDueToDeadWarning"}`,
			expected: packet.ActionIgnoredDueToDeadWarning,
		},
		{
			input:    `{"type": "invalidPacketTypeError"}`,
			expected: packet.InvalidPacketTypeError,
		},
		{
			input:    `{"type": "invalidPacketUsageError"}`,
			expected: packet.InvalidPacketUsageError,
		},
		{
			input:    `{"type": "movement"}`,
			expected: packet.MovementPacket,
		},
		{
			input:    `{"type": "rotation"}`,
			expected: packet.RotationPacket,
		},
		{
			input:    `{"type": "abilityUse"}`,
			expected: packet.AbilityUsePacket,
		},
		{
			input:    `{"type": "pass"}`,
			expected: packet.PassPacket,
		},
		{
			input:    `{"type": "gameNotStarted"}`,
			expected: packet.GameNotStarted,
		},
		{
			input:    `{"type": "gameInProgress"}`,
			expected: packet.GameInProgress,
		},
		{
			input:    `{"type": "gameStatusRequest"}`,
			expected: packet.GameStatusRequest,
		},
		{
			input:    `{"type": "gameState"}`,
			expected: packet.GameStatePacket,
		},
		{
			input:    `{"type": "lobbyDataRequest"}`,
			expected: packet.LobbyDataRequest,
		},
		{
			input:    `{"type": "customWarning"}`,
			expected: packet.CustomWarning,
		},
	}

	for _, test := range tests {
		var p packet.Packet
		err := json.Unmarshal([]byte(test.input), &p)
		if err != nil {
			t.Fatalf("Error unmarshalling packet: %v", err)
		}
		if p.Type != test.expected {
			t.Fatalf("Expected type %v, got %v", test.expected, p.Type)
		}
		if p.Payload != nil {
			t.Fatalf("Expected payload to be nil, got %v", p.Payload)
		}
	}
}

func TestPacketUnmarshalJSONWithPayload(t *testing.T) {
	input := `{
        "type": "lobbyData",
        "payload": {
            "playerId": "123",
            "players": [
                {"id": "player1", "nickname": "Player One", "color": 16711680},
                {"id": "player2", "nickname": "Player Two", "color": 65280}
            ],
            "serverSettings": {
                "gridDimension": 20,
                "numberOfPlayers": 4,
                "seed": 12345,
                "broadcastInterval": 1000,
                "eagerBroadcast": true
            }
        }
    }`
	var p packet.Packet
	err := json.Unmarshal([]byte(input), &p)
	if err != nil {
		t.Fatalf("Error unmarshalling packet: %v", err)
	}
	if p.Type != packet.LobbyDataPacket {
		t.Fatalf("Expected type %v, got %v", packet.LobbyDataPacket, p.Type)
	}
	if _, ok := p.Payload.(map[string]interface{}); !ok {
		t.Fatalf("Expected payload to be of type map[string]interface{}, got %T", p.Payload)
	}
}

func TestPacketUnmarshalJSONWithLobbyData(t *testing.T) {
	input := `{
        "type": "lobbyData",
        "payload": {
            "playerId": "123",
            "players": [
                {"id": "player1", "nickname": "Player One", "color": 16711680},
                {"id": "player2", "nickname": "Player Two", "color": 65280}
            ],
            "serverSettings": {
                "gridDimension": 20,
                "numberOfPlayers": 4,
                "seed": 12345,
                "broadcastInterval": 1000,
                "eagerBroadcast": true
            }
        }
    }`
	var p packet.Packet
	err := json.Unmarshal([]byte(input), &p)
	if err != nil {
		t.Fatalf("Error unmarshalling packet: %v", err)
	}
	if p.Type != packet.LobbyDataPacket {
		t.Fatalf("Expected type %v, got %v", packet.LobbyDataPacket, p.Type)
	}
	payload, ok := p.Payload.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected payload to be of type map[string]interface{}, got %T", p.Payload)
	}
	if payload["playerId"] != "123" {
		t.Fatalf("Expected payload playerId to be '123', got %v", payload["playerId"])
	}
	players, ok := payload["players"].([]interface{})
	if !ok {
		t.Fatalf("Expected players to be of type []interface{}, got %T", payload["players"])
	}
	if len(players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(players))
	}
	player1 := players[0].(map[string]interface{})
	if player1["id"] != "player1" {
		t.Fatalf("Expected player1 id to be 'player1', got %v", player1["id"])
	}
	if player1["nickname"] != "Player One" {
		t.Fatalf("Expected player1 nickname to be 'Player One', got %v", player1["nickname"])
	}
	if player1["color"] != float64(16711680) { // JSON unmarshals numbers to float64 by default
		t.Fatalf("Expected player1 color to be 16711680, got %v", player1["color"])
	}
	player2 := players[1].(map[string]interface{})
	if player2["id"] != "player2" {
		t.Fatalf("Expected player2 id to be 'player2', got %v", player2["id"])
	}
	if player2["nickname"] != "Player Two" {
		t.Fatalf("Expected player2 nickname to be 'Player Two', got %v", player2["nickname"])
	}
	if player2["color"] != float64(65280) { // JSON unmarshals numbers to float64 by default
		t.Fatalf("Expected player2 color to be 65280, got %v", player2["color"])
	}
}

func TestPacketUnmarshalJSONWithLobbyDataStruct(t *testing.T) {
	input := `{
        "type": "lobbyData",
        "payload": {
            "playerId": "123",
            "players": [
                {"id": "player1", "nickname": "Player One", "color": 16711680},
                {"id": "player2", "nickname": "Player Two", "color": 65280}
            ],
            "serverSettings": {
                "gridDimension": 20,
                "numberOfPlayers": 4,
                "seed": 12345,
                "broadcastInterval": 1000,
                "eagerBroadcast": true
            }
        }
    }`
	var p packet.Packet
	err := json.Unmarshal([]byte(input), &p)
	if err != nil {
		t.Fatalf("Error unmarshalling packet: %v", err)
	}
	if p.Type != packet.LobbyDataPacket {
		t.Fatalf("Expected type %v, got %v", packet.LobbyDataPacket, p.Type)
	}

	// Parse payload into LobbyData struct
	var lobbyData lobby_data.LobbyData
	payloadBytes, err := json.Marshal(p.Payload)
	if err != nil {
		t.Fatalf("Error marshalling payload: %v", err)
	}
	err = json.Unmarshal(payloadBytes, &lobbyData)
	if err != nil {
		t.Fatalf("Error unmarshalling payload into LobbyData: %v", err)
	}

	if lobbyData.PlayerID != "123" {
		t.Fatalf("Expected playerId to be '123', got %v", lobbyData.PlayerID)
	}
	if len(lobbyData.Players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(lobbyData.Players))
	}
	if lobbyData.Players[0].ID != "player1" {
		t.Fatalf("Expected player1 id to be 'player1', got %v", lobbyData.Players[0].ID)
	}
	if lobbyData.Players[0].Nickname != "Player One" {
		t.Fatalf("Expected player1 nickname to be 'Player One', got %v", lobbyData.Players[0].Nickname)
	}
	if lobbyData.Players[0].Color != 16711680 {
		t.Fatalf("Expected player1 color to be 16711680, got %v", lobbyData.Players[0].Color)
	}
	if lobbyData.Players[1].ID != "player2" {
		t.Fatalf("Expected player2 id to be 'player2', got %v", lobbyData.Players[1].ID)
	}
	if lobbyData.Players[1].Nickname != "Player Two" {
		t.Fatalf("Expected player2 nickname to be 'Player Two', got %v", lobbyData.Players[1].Nickname)
	}
	if lobbyData.Players[1].Color != 65280 {
		t.Fatalf("Expected player2 color to be 65280, got %v", lobbyData.Players[1].Color)
	}
	if lobbyData.ServerSettings.GridDimension != 20 {
		t.Fatalf("Expected gridDimension to be 20, got %v", lobbyData.ServerSettings.GridDimension)
	}
	if lobbyData.ServerSettings.NumberOfPlayers != 4 {
		t.Fatalf("Expected numberOfPlayers to be 4, got %v", lobbyData.ServerSettings.NumberOfPlayers)
	}
	if lobbyData.ServerSettings.Seed != 12345 {
		t.Fatalf("Expected seed to be 12345, got %v", lobbyData.ServerSettings.Seed)
	}
	if lobbyData.ServerSettings.BroadcastInterval != 1000 {
		t.Fatalf("Expected broadcastInterval to be 1000, got %v", lobbyData.ServerSettings.BroadcastInterval)
	}
	if lobbyData.ServerSettings.EagerBroadcast != true {
		t.Fatalf("Expected eagerBroadcast to be true, got %v", lobbyData.ServerSettings.EagerBroadcast)
	}
}

func TestPacketUnmarshalJSONWithGameEnd(t *testing.T) {
	input := `{
        "type": "gameEnded",
        "payload": {
            "players": [
                {"id": "player1", "score": 10},
                {"id": "player2", "score": 20}
            ]
        }
    }`
	var p packet.Packet
	err := json.Unmarshal([]byte(input), &p)
	if err != nil {
		t.Fatalf("Error unmarshalling packet: %v", err)
	}
	if p.Type != packet.GameEndedPacket {
		t.Fatalf("Expected type %v, got %v", packet.GameEndedPacket, p.Type)
	}
	payload, ok := p.Payload.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected payload to be of type map[string]interface{}, got %T", p.Payload)
	}
	var gameEnd game_end.GameEnd
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Error marshalling payload: %v", err)
	}
	err = json.Unmarshal(payloadBytes, &gameEnd)
	if err != nil {
		t.Fatalf("Error unmarshalling payload to GameEnd: %v", err)
	}
	if len(gameEnd.Players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(gameEnd.Players))
	}
	player1 := gameEnd.Players[0]
	if player1.ID != "player1" {
		t.Fatalf("Expected player1 id to be 'player1', got %v", player1.ID)
	}
	if player1.Score != 10 {
		t.Fatalf("Expected player1 score to be 10, got %v", player1.Score)
	}
	player2 := gameEnd.Players[1]
	if player2.ID != "player2" {
		t.Fatalf("Expected player2 id to be 'player2', got %v", player2.ID)
	}
	if player2.Score != 20 {
		t.Fatalf("Expected player2 score to be 20, got %v", player2.Score)
	}
}
