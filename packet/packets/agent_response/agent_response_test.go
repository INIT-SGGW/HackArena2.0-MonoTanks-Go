package agent_response

import (
	"encoding/json"
	"hack-arena-2024-h2-go/packet"
	"testing"
)

func TestAgentResponseSerialization(t *testing.T) {
	tests := []struct {
		name         string
		response     *AgentResponse
		gameStateID  string
		expectedJSON string
		expectedType packet.PacketType
	}{
		{
			name:         "Movement Forward",
			response:     NewMovement(0),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"movement","payload":{"direction":0,"gameStateId":"test-id"}}`,
			expectedType: packet.MovementPacket,
		},
		{
			name:         "Movement Backward",
			response:     NewMovement(1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"movement","payload":{"direction":1,"gameStateId":"test-id"}}`,
			expectedType: packet.MovementPacket,
		},
		{
			name:         "Rotation Left Right",
			response:     NewRotation(0, 1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"rotation","payload":{"gameStateId":"test-id","tankRotation":0,"turretRotation":1}}`,
			expectedType: packet.RotationPacket,
		},
		{
			name:         "AbilityUse",
			response:     NewAbilityUse(0),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"abilityUse","payload":{"abilityType":0,"gameStateId":"test-id"}}`,
			expectedType: packet.AbilityUsePacket,
		},
		{
			name:         "Rotation Left None",
			response:     NewRotation(0, -1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"rotation","payload":{"gameStateId":"test-id","tankRotation":0}}`,
			expectedType: packet.RotationPacket,
		},
		{
			name:         "Rotation None Right",
			response:     NewRotation(-1, 1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"rotation","payload":{"gameStateId":"test-id","turretRotation":1}}`,
			expectedType: packet.RotationPacket,
		},
		{
			name:         "Rotation None None",
			response:     NewRotation(-1, -1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"rotation","payload":{"gameStateId":"test-id"}}`,
			expectedType: packet.RotationPacket,
		},
		{
			name:         "Pass",
			response:     NewPass(),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"pass","payload":{"gameStateId":"test-id"}}`,
			expectedType: packet.PassPacket,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkt := tt.response.ToPacket(tt.gameStateID)

			if pkt.Type != tt.expectedType {
				t.Errorf("Expected packet type %v, got %v", tt.expectedType, pkt.Type)
			}

			data, err := json.Marshal(pkt)
			if err != nil {
				t.Fatalf("Failed to marshal packet: %v", err)
			}

			if string(data) != tt.expectedJSON {
				t.Errorf("Expected JSON:\n%s\nGot:\n%s", tt.expectedJSON, string(data))
			}
		})
	}
}
