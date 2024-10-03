package packet

import (
	"encoding/json"
	"fmt"
)

type PacketType string

const (
	Ping                           PacketType = "ping"
	Pong                           PacketType = "pong"
	ConnectionAccepted             PacketType = "connectionAccepted"
	ConnectionRejected             PacketType = "connectionRejected"
	LobbyDataPacket                PacketType = "lobbyData"
	LobbyDeleted                   PacketType = "lobbyDeleted"
	GameStart                      PacketType = "gameStart"
	GameStatePacket                PacketType = "gameState"
	TankMovementPacket             PacketType = "tankMovement"
	TankRotationPacket             PacketType = "tankRotation"
	TankShootPacket                PacketType = "tankShoot"
	GameEndPacket                  PacketType = "gameEnd"
	PlayerAlreadyMadeActionWarning PacketType = "playerAlreadyMadeActionWarning"
	MissingGameStateIdWarning      PacketType = "missingGameStateIdWarning"
	SlowResponseWarning            PacketType = "slowResponseWarning"
	InvalidPacketTypeError         PacketType = "invalidPacketTypeError"
	InvalidPacketUsageError        PacketType = "invalidPacketUsageError"
)

type Packet struct {
	Type    PacketType  `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewPacket(packetType PacketType, payload interface{}) *Packet {
	return &Packet{
		Type:    packetType,
		Payload: payload,
	}
}

func (p *Packet) MarshalJSON() ([]byte, error) {
	type Alias Packet
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	})
}

func (p *Packet) UnmarshalJSON(data []byte) error {
	type Alias Packet
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch aux.Type {
	case LobbyDataPacket:
		var payload map[string]interface{}
		if err := json.Unmarshal(data, &payload); err != nil {
			return err
		}
		p.Payload = payload["payload"]
	default:
		// For other packet types, we assume no specific payload structure
		p.Payload = aux.Payload
	}

	p.Type = aux.Type
	return nil
}

func (p Packet) String() string {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(data)
}
