package agent_response

import (
	"encoding/json"
	"hack-arena-2024-h2-go/packet"
	"hack-arena-2024-h2-go/packet/packets/agent_response/ability"
	"hack-arena-2024-h2-go/packet/packets/agent_response/movement"
	"hack-arena-2024-h2-go/packet/packets/agent_response/rotation"
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
			response:     NewMovement(movement.Forward),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"movement","payload":{"direction":"forward","gameStateId":"test-id"}}`,
			expectedType: packet.MovementPacket,
		},
		{
			name:         "Movement Backward",
			response:     NewMovement(movement.Backward),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"movement","payload":{"direction":"backward","gameStateId":"test-id"}}`,
			expectedType: packet.MovementPacket,
		},
		{
			name:         "Rotation Left Right",
			response:     NewRotation(rotation.Left, rotation.Right),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"rotation","payload":{"gameStateId":"test-id","tankRotation":"left","turretRotation":"right"}}`,
			expectedType: packet.RotationPacket,
		},
		{
			name:         "AbilityUse",
			response:     NewAbilityUse(ability.FireBullet),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"abilityUse","payload":{"abilityType":"fireBullet","gameStateId":"test-id"}}`,
			expectedType: packet.AbilityUsePacket,
		},
		{
			name:         "Rotation Left None",
			response:     NewRotation(rotation.Left, ""),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"rotation","payload":{"gameStateId":"test-id","tankRotation":"left"}}`,
			expectedType: packet.RotationPacket,
		},
		{
			name:         "Rotation None Right",
			response:     NewRotation("", rotation.Right),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"rotation","payload":{"gameStateId":"test-id","turretRotation":"right"}}`,
			expectedType: packet.RotationPacket,
		},
		{
			name:         "Rotation None None",
			response:     NewRotation("", ""),
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
