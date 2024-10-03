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
		return json.Marshal(&struct {
			Direction *MoveDirection `json:"direction,omitempty"`
		}{
			Direction: ar.Direction,
		})
	case TankRotation:
		// Handle None as null in the JSON output
		var tankRotation, turretRotation *Rotation
		if ar.TankRotation != None {
			tankRotation = &ar.TankRotation
		}
		if ar.TurretRotation != None {
			turretRotation = &ar.TurretRotation
		}
		return json.Marshal(&struct {
			TankRotation   *Rotation `json:"tankRotation"`   // If None, will be serialized as null
			TurretRotation *Rotation `json:"turretRotation"` // If None, will be serialized as null
		}{
			TankRotation:   tankRotation,
			TurretRotation: turretRotation,
		})
	case TankShoot:
		return json.Marshal(&struct{}{})
	default:
		return nil, errors.New("invalid response type")
	}
}

// UnmarshalJSON customizes the JSON deserialization of the AgentResponse type.
func (ar *AgentResponse) UnmarshalJSON(data []byte) error {
	type Alias AgentResponse
	aux := &struct {
		Type string `json:"type"`
		*Alias
	}{
		Alias: (*Alias)(ar),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ar.Type = ResponseType(aux.Type)

	switch ar.Type {
	case TankMovement, TankRotation, TankShoot:
		return nil
	default:
		return errors.New("invalid response type")
	}
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
		return packet.Packet{
			Type: packet.TankRotationPacket,
			Payload: map[string]interface{}{
				"gameStateId":    gameStateID,
				"tankRotation":   ar.TankRotation,
				"turretRotation": ar.TurretRotation,
			},
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
