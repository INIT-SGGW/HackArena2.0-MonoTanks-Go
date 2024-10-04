package agent_response

import (
	"encoding/json"
	"errors"
	"hack-arena-2024-h2-go/packet"
)

// AgentResponse represents the various responses an agent can have in the system.
type AgentResponse struct {
	Type           ResponseType   `json:"-"`
	Direction      *MoveDirection `json:"direction,omitempty"`
	TankRotation   Rotation       `json:"tankRotation,omitempty"`
	TurretRotation Rotation       `json:"turretRotation,omitempty"`
}

// ResponseType is an enumeration of the types of responses an agent can have.
type ResponseType string

const (
	TankMovement ResponseType = "tankMovement"
	TankRotation ResponseType = "tankRotation"
	TankShoot    ResponseType = "tankShoot"
)

// Rotation represents the rotation data for a tank or turret.
type Rotation int

const (
	None  Rotation = -1 // `None` represents no rotation
	Left  Rotation = 0  // Preserve `Left` as 0
	Right Rotation = 1  // Preserve `Right` as 1
)

// MoveDirection represents the direction of movement.
type MoveDirection int

const (
	Forward MoveDirection = iota
	Backward
)

// NewTankMovement creates a new AgentResponse for tank movement.
func NewTankMovement(direction MoveDirection) *AgentResponse {
	return &AgentResponse{
		Type:      TankMovement,
		Direction: &direction,
	}
}

// NewTankRotation creates a new AgentResponse for tank rotation.
func NewTankRotation(tankRotation, turretRotation Rotation) *AgentResponse {
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

// MarshalJSON customizes the JSON representation of the AgentResponse type.
func (ar *AgentResponse) MarshalJSON() ([]byte, error) {
	switch ar.Type {
	case TankMovement:
		return json.Marshal(struct {
			Direction *MoveDirection `json:"direction,omitempty"`
		}{ar.Direction})

	case TankRotation:
		rotations := make(map[string]int)
		if ar.TankRotation != None {
			rotations["tankRotation"] = int(ar.TankRotation)
		}
		if ar.TurretRotation != None {
			rotations["turretRotation"] = int(ar.TurretRotation)
		}
		return json.Marshal(rotations)

	case TankShoot:
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
		ar.Type = TankShoot
		return nil
	default:
		return errors.New("invalid response type")
	}
}

func (ar *AgentResponse) unmarshalTankMovement(temp map[string]json.RawMessage) error {
	ar.Type = TankMovement
	var direction MoveDirection
	if err := json.Unmarshal(temp["direction"], &direction); err != nil {
		return err
	}
	ar.Direction = &direction
	return nil
}

func (ar *AgentResponse) unmarshalTankRotation(temp map[string]json.RawMessage) error {
	ar.Type = TankRotation
	ar.TankRotation = None
	ar.TurretRotation = None

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
		if ar.TankRotation != None {
			payload["tankRotation"] = ar.TankRotation
		}
		if ar.TurretRotation != None {
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
	default:
		return packet.Packet{
			Type: packet.InvalidPacketTypeError,
		}
	}
}
