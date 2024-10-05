package packet

import (
	"encoding/json"
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
	ActionIgnoredDueToDeadWarning  PacketType = "actionIgnoredDueToDeadWarning"
)

type Packet struct {
	Type    PacketType  `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

func (p *Packet) MarshalJSON() ([]byte, error) {
	type Alias Packet
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	})
}
