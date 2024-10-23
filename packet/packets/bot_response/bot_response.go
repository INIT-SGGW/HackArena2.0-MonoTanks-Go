package bot_response

import (
	"encoding/json"
	"errors"
	"hackarena2-0-mono-tanks-go/packet"
)

// BotResponse represents the various responses an bot can have in the system.
type BotResponse struct {
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

// ResponseType is an enumeration of the types of responses an bot can have.
type ResponseType string

const (
	Movement   ResponseType = "movement"
	Rotation   ResponseType = "rotation"
	AbilityUse ResponseType = "abilityUse"
	Pass       ResponseType = "pass"
)

// NewMovement creates a new BotResponse for tank movement.
// direction: "forward" or "backward"
func NewMovement(direction string) *BotResponse {
	return &BotResponse{
		Type:      Movement,
		Direction: direction,
	}
}

// NewRotation creates a new BotResponse for tank rotation.
// Both tankRotation and turretRotation use the following values:
// "left" or "right", empty if not applicable
func NewRotation(tankRotation, turretRotation string) *BotResponse {
	return &BotResponse{
		Type:           Rotation,
		TankRotation:   tankRotation,
		TurretRotation: turretRotation,
	}
}

// NewAbilityUse creates a new BotResponse for ability use.
func NewAbilityUse(abilityType string) *BotResponse {
	return &BotResponse{
		Type:        AbilityUse,
		AbilityType: abilityType,
	}
}

// NewPass creates a new BotResponse for response pass.
func NewPass() *BotResponse {
	return &BotResponse{
		Type: Pass,
	}
}

// MarshalJSON customizes the JSON representation of the BotResponse type.
func (ar *BotResponse) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON customizes the JSON deserialization of the BotResponse type.
func (ar *BotResponse) UnmarshalJSON(data []byte) error {
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

func (ar *BotResponse) unmarshalMovement(temp map[string]json.RawMessage) error {
	ar.Type = Movement
	return json.Unmarshal(temp["direction"], &ar.Direction)
}

func (ar *BotResponse) unmarshalRotation(temp map[string]json.RawMessage) error {
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

func (ar *BotResponse) unmarshalAbilityUse(temp map[string]json.RawMessage) error {
	ar.Type = AbilityUse
	return json.Unmarshal(temp["abilityType"], &ar.AbilityType)
}

func hasKey(m map[string]json.RawMessage, key string) bool {
	_, ok := m[key]
	return ok
}

func (ar BotResponse) ToPacket(gameStateID string) packet.Packet {
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
