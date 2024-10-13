package agent_response

import (
	"encoding/json"
	"errors"
	"hack-arena-2024-h2-go/packet"
)

// AgentResponse represents the various responses an agent can have in the system.
type AgentResponse struct {
	// Type represents the kind of response (e.g., Movement, Rotation, AbilityUse)
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

	// AbilityType represents the type of ability to use
	AbilityType int `json:"abilityType,omitempty"`
}

// ResponseType is an enumeration of the types of responses an agent can have.
type ResponseType string

const (
	Movement   ResponseType = "movement"
	Rotation   ResponseType = "rotation"
	AbilityUse ResponseType = "abilityUse"
	Pass       ResponseType = "pass"
)

// NewMovement creates a new AgentResponse for tank movement.
// direction: 0 for forward, 1 for backward
func NewMovement(direction int) *AgentResponse {
	return &AgentResponse{
		Type:      Movement,
		Direction: direction,
	}
}

// NewRotation creates a new AgentResponse for tank rotation.
// Both tankRotation and turretRotation use the following values:
// -1: no rotation
//
//	0: rotate left
//	1: rotate right
func NewRotation(tankRotation, turretRotation int) *AgentResponse {
	return &AgentResponse{
		Type:           Rotation,
		TankRotation:   tankRotation,
		TurretRotation: turretRotation,
	}
}

// NewAbilityUse creates a new AgentResponse for ability use.
func NewAbilityUse(abilityType int) *AgentResponse {
	return &AgentResponse{
		Type:        AbilityUse,
		AbilityType: abilityType,
	}
}

// NewPass creates a new AgentResponse for response pass.
func NewPass() *AgentResponse {
	return &AgentResponse{
		Type: Pass,
	}
}

// MarshalJSON customizes the JSON representation of the AgentResponse type.
func (ar *AgentResponse) MarshalJSON() ([]byte, error) {
	switch ar.Type {
	case Movement:
		return json.Marshal(struct {
			Direction int `json:"direction"`
		}{ar.Direction})

	case Rotation:
		rotations := make(map[string]int)
		if ar.TankRotation != -1 {
			rotations["tankRotation"] = ar.TankRotation
		}
		if ar.TurretRotation != -1 {
			rotations["turretRotation"] = ar.TurretRotation
		}
		return json.Marshal(rotations)

	case AbilityUse:
		return json.Marshal(struct {
			AbilityType int `json:"abilityType"`
		}{ar.AbilityType})

	case Pass:
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
		return ar.unmarshalMovement(temp)
	case hasKey(temp, "tankRotation") || hasKey(temp, "turretRotation"):
		return ar.unmarshalRotation(temp)
	case hasKey(temp, "abilityType"):
		return ar.unmarshalAbilityUse(temp)
	case len(temp) == 0:
		ar.Type = Pass
		return nil
	default:
		return errors.New("invalid response type")
	}
}

func (ar *AgentResponse) unmarshalMovement(temp map[string]json.RawMessage) error {
	ar.Type = Movement
	return json.Unmarshal(temp["direction"], &ar.Direction)
}

func (ar *AgentResponse) unmarshalRotation(temp map[string]json.RawMessage) error {
	ar.Type = Rotation
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

func (ar *AgentResponse) unmarshalAbilityUse(temp map[string]json.RawMessage) error {
	ar.Type = AbilityUse
	return json.Unmarshal(temp["abilityType"], &ar.AbilityType)
}

func hasKey(m map[string]json.RawMessage, key string) bool {
	_, ok := m[key]
	return ok
}

func (ar AgentResponse) ToPacket(gameStateID string) packet.Packet {
	switch ar.Type {
	case Movement:
		return packet.Packet{
			Type: packet.MovementPacket,
			Payload: map[string]interface{}{
				"gameStateId": gameStateID,
				"direction":   ar.Direction,
			},
		}
	case Rotation:
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
			Type:    packet.RotationPacket,
			Payload: payload,
		}
	case AbilityUse:
		return packet.Packet{
			Type: packet.AbilityUsePacket,
			Payload: map[string]interface{}{
				"gameStateId": gameStateID,
				"abilityType": ar.AbilityType,
			},
		}
	case Pass:
		return packet.Packet{
			Type: packet.PassPacket,
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
