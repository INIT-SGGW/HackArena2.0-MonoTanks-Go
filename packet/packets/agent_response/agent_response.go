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
	// "forward" or "backward", empty if not applicable
	Direction string `json:"direction,omitempty"`

	// TankRotation specifies the rotation of the tank body
	// "left" or "right", empty if not applicable
	TankRotation string `json:"tankRotation,omitempty"`

	// TurretRotation specifies the rotation of the tank's turret
	// "left" or "right", empty if not applicable
	TurretRotation string `json:"turretRotation,omitempty"`

	// AbilityType represents the type of ability to use
	AbilityType string `json:"abilityType,omitempty"`
}

// ResponseType is an enumeration of the types of responses an agent can have.
type ResponseType string

const (
	Movement         ResponseType = "movement"
	Rotation         ResponseType = "rotation"
	AbilityUse       ResponseType = "abilityUse"
	Pass             ResponseType = "pass"
	FireBullet                    = "fireBullet"
	UseLaser                      = "useLaser"
	FireDoubleBullet              = "fireDoubleBullet"
	UseRadar                      = "useRadar"
	DropMine                      = "dropMine"
	Forward                       = "forward"
	Backward                      = "backward"
	Left                          = "left"
	Right                         = "right"
)

// NewMovement creates a new AgentResponse for tank movement.
// direction: "forward" or "backward"
func NewMovement(direction string) *AgentResponse {
	return &AgentResponse{
		Type:      Movement,
		Direction: direction,
	}
}

// NewRotation creates a new AgentResponse for tank rotation.
// Both tankRotation and turretRotation use the following values:
// "left" or "right", empty if not applicable
func NewRotation(tankRotation, turretRotation string) *AgentResponse {
	return &AgentResponse{
		Type:           Rotation,
		TankRotation:   tankRotation,
		TurretRotation: turretRotation,
	}
}

// NewAbilityUse creates a new AgentResponse for ability use.
func NewAbilityUse(abilityType string) *AgentResponse {
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
			Direction string `json:"direction"`
		}{ar.Direction})

	case Rotation:
		rotations := make(map[string]string)
		if ar.TankRotation != "" {
			rotations["tankRotation"] = ar.TankRotation
		}
		if ar.TurretRotation != "" {
			rotations["turretRotation"] = ar.TurretRotation
		}
		return json.Marshal(rotations)

	case AbilityUse:
		return json.Marshal(struct {
			AbilityType string `json:"abilityType"`
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
		if ar.TankRotation != "" {
			payload["tankRotation"] = ar.TankRotation
		}
		if ar.TurretRotation != "" {
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
