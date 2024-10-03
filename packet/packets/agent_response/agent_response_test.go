package agent_response

import (
	"encoding/json"
	"testing"
)

func TestAgentResponseSerialization(t *testing.T) {
	tests := []struct {
		name     string
		response *AgentResponse
		expected string
	}{
		{
			name:     "TankMovement Forward",
			response: NewTankMovement(Forward),
			expected: `{"direction":0}`,
		},
		{
			name:     "TankMovement Backward",
			response: NewTankMovement(Backward),
			expected: `{"direction":1}`,
		},
		{
			name:     "TankRotation Left Right",
			response: NewTankRotation(Left, Right),
			expected: `{"tankRotation":0,"turretRotation":1}`,
		},
		{
			name:     "TankShoot",
			response: NewTankShoot(),
			expected: `{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.response)
			if err != nil {
				t.Fatalf("Failed to marshal response: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}
