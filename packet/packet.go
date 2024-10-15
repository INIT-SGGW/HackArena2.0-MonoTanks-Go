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
	MovementPacket                 PacketType = "movement"
	RotationPacket                 PacketType = "rotation"
	AbilityUsePacket               PacketType = "abilityUse"
	GameEndPacket                  PacketType = "gameEnd"
	CustomWarning                  PacketType = "customWarning"
	PlayerAlreadyMadeActionWarning PacketType = "playerAlreadyMadeActionWarning"
	ActionIgnoredDueToDeadWarning  PacketType = "actionIgnoredDueToDeadWarning"
	SlowResponseWarning            PacketType = "slowResponseWarning"
	InvalidPacketTypeError         PacketType = "invalidPacketTypeError"
	InvalidPacketUsageError        PacketType = "invalidPacketUsageError"
	PassPacket                     PacketType = "pass"
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
