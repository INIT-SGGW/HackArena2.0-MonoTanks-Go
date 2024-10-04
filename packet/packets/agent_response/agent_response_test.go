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
			name:         "TankMovement Forward",
			response:     NewTankMovement(0),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"tankMovement","payload":{"direction":0,"gameStateId":"test-id"}}`,
			expectedType: packet.TankMovementPacket,
		},
		{
			name:         "TankMovement Backward",
			response:     NewTankMovement(1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"tankMovement","payload":{"direction":1,"gameStateId":"test-id"}}`,
			expectedType: packet.TankMovementPacket,
		},
		{
			name:         "TankRotation Left Right",
			response:     NewTankRotation(0, 1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"tankRotation","payload":{"gameStateId":"test-id","tankRotation":0,"turretRotation":1}}`,
			expectedType: packet.TankRotationPacket,
		},
		{
			name:         "TankShoot",
			response:     NewTankShoot(),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"tankShoot","payload":{"gameStateId":"test-id"}}`,
			expectedType: packet.TankShootPacket,
		},
		{
			name:         "TankRotation Left None",
			response:     NewTankRotation(0, -1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"tankRotation","payload":{"gameStateId":"test-id","tankRotation":0}}`,
			expectedType: packet.TankRotationPacket,
		},
		{
			name:         "TankRotation None Right",
			response:     NewTankRotation(-1, 1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"tankRotation","payload":{"gameStateId":"test-id","turretRotation":1}}`,
			expectedType: packet.TankRotationPacket,
		},
		{
			name:         "TankRotation None None",
			response:     NewTankRotation(-1, -1),
			gameStateID:  "test-id",
			expectedJSON: `{"type":"tankRotation","payload":{"gameStateId":"test-id"}}`,
			expectedType: packet.TankRotationPacket,
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
