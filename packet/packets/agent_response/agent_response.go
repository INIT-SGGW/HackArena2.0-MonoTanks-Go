package agent_response

import (
	"encoding/json"
	"errors"
	"hack-arena-2024-h2-go/packet"
)

// AgentResponse represents the various responses an agent can have in the system.
type AgentResponse struct {
	// Type represents the kind of response (e.g., TankMovement, TankRotation, TankShoot)
	Type ResponseType `json:"-"`

	// Direction indicates the movement direction of the tank
	// 0 for forward, 1 for backward, nil if not applicable
	Direction int `json:"direction,omitempty"`

	// TankRotation specifies the rotation of the tank body
	// -1: no rotation, 0: rotate left, 1: rotate right
	TankRotation int `json:"tankRotation,omitempty"`

	// TurretRotation specifies the rotation of the tank's turret
	// -1: no rotation, 0: rotate left, 1: rotate right
	TurretRotation int `json:"turretRotation,omitempty"`
}

// ResponseType is an enumeration of the types of responses an agent can have.
type ResponseType string

const (
	TankMovement ResponseType = "tankMovement"
	TankRotation ResponseType = "tankRotation"
	TankShoot    ResponseType = "tankShoot"
	ResponsePass ResponseType = "responsePass"
)

// NewTankMovement creates a new AgentResponse for tank movement.
// direction: 0 for forward, 1 for backward
func NewTankMovement(direction int) *AgentResponse {
	return &AgentResponse{
		Type:      TankMovement,
		Direction: direction,
	}
}

// NewTankRotation creates a new AgentResponse for tank rotation.
// Both tankRotation and turretRotation use the following values:
// -1: no rotation
//
//	0: rotate left
//	1: rotate right
func NewTankRotation(tankRotation, turretRotation int) *AgentResponse {
	return &AgentResponse{
		Type:           TankRotation,
		TankRotation:   tankRotation,
		TurretRotation: turretRotation,
	}
}

// NewTankShoot creates a new AgentResponse for tank shooting.
func NewTankShoot() *AgentResponse {
	return &AgentResponse{
		Type: TankShoot,
	}
}

// NewResponsePass creates a new AgentResponse for response pass.
func NewResponsePass() *AgentResponse {
	return &AgentResponse{
		Type: ResponsePass,
	}
}

// MarshalJSON customizes the JSON representation of the AgentResponse type.
func (ar *AgentResponse) MarshalJSON() ([]byte, error) {
	switch ar.Type {
	case TankMovement:
		return json.Marshal(struct {
			Direction int `json:"direction"`
		}{ar.Direction})

	case TankRotation:
		rotations := make(map[string]int)
		if ar.TankRotation != -1 {
			rotations["tankRotation"] = ar.TankRotation
		}
		if ar.TurretRotation != -1 {
			rotations["turretRotation"] = ar.TurretRotation
		}
		return json.Marshal(rotations)

	case TankShoot:
		return json.Marshal(struct{}{})

	case ResponsePass:
		return json.Marshal(struct{}{})

	default:
		return nil, errors.New("invalid response type")
	}
}

// UnmarshalJSON customizes the JSON deserialization of the AgentResponse type.
func (ar *AgentResponse) UnmarshalJSON(data []byte) error {
	var temp map[string]json.RawMessage
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch {
	case hasKey(temp, "direction"):
		return ar.unmarshalTankMovement(temp)
	case hasKey(temp, "tankRotation") || hasKey(temp, "turretRotation"):
		return ar.unmarshalTankRotation(temp)
	case len(temp) == 0:
		ar.Type = ResponsePass
		return nil
	default:
		return errors.New("invalid response type")
	}
}

func (ar *AgentResponse) unmarshalTankMovement(temp map[string]json.RawMessage) error {
	ar.Type = TankMovement
	return json.Unmarshal(temp["direction"], &ar.Direction)
}

func (ar *AgentResponse) unmarshalTankRotation(temp map[string]json.RawMessage) error {
	ar.Type = TankRotation
	ar.TankRotation = -1
	ar.TurretRotation = -1

	if tankRotation, ok := temp["tankRotation"]; ok {
		if err := json.Unmarshal(tankRotation, &ar.TankRotation); err != nil {
			return err
		}
	}
	if turretRotation, ok := temp["turretRotation"]; ok {
		if err := json.Unmarshal(turretRotation, &ar.TurretRotation); err != nil {
			return err
		}
	}
	return nil
}

func hasKey(m map[string]json.RawMessage, key string) bool {
	_, ok := m[key]
	return ok
}

func (ar AgentResponse) ToPacket(gameStateID string) packet.Packet {
	switch ar.Type {
	case TankMovement:
		return packet.Packet{
			Type: packet.TankMovementPacket,
			Payload: map[string]interface{}{
				"gameStateId": gameStateID,
				"direction":   ar.Direction,
			},
		}
	case TankRotation:
		payload := map[string]interface{}{
			"gameStateId": gameStateID,
		}
		if ar.TankRotation != -1 {
			payload["tankRotation"] = ar.TankRotation
		}
		if ar.TurretRotation != -1 {
			payload["turretRotation"] = ar.TurretRotation
		}
		return packet.Packet{
			Type:    packet.TankRotationPacket,
			Payload: payload,
		}
	case TankShoot:
		return packet.Packet{
			Type: packet.TankShootPacket,
			Payload: map[string]interface{}{
				"gameStateId": gameStateID,
			},
		}
	case ResponsePass:
		return packet.Packet{
			Type: packet.ResponsePassPacket,
			Payload: map[string]interface{}{
				"gameStateId": gameStateID,
			},
		}
	default:
		return packet.Packet{
			Type: packet.InvalidPacketTypeError,
		}
	}
}
