package packet

import (
	"encoding/json"
)

type PacketType string

const (
	Ping PacketType = "ping"
	Pong PacketType = "pong"

	ConnectionRejected PacketType = "connectionRejected"
	ConnectionAccepted PacketType = "connectionAccepted"

	LobbyDataPacket  PacketType = "lobbyData"
	LobbyDataRequest PacketType = "lobbyDataRequest"

	GameNotStarted PacketType = "gameNotStarted"
	GameStarting   PacketType = "gameStarting"
	GameStarted    PacketType = "gameStarted"
	GameInProgress PacketType = "gameInProgress"

	GameStatusRequest       PacketType = "gameStatusRequest"
	ReadyToReceiveGameState PacketType = "readyToReceiveGameState"

	GameStatePacket  PacketType = "gameState"
	MovementPacket   PacketType = "movement"
	RotationPacket   PacketType = "rotation"
	AbilityUsePacket PacketType = "abilityUse"
	PassPacket       PacketType = "pass"

	GameEndedPacket PacketType = "gameEnded"

	CustomWarning                  PacketType = "customWarning"
	PlayerAlreadyMadeActionWarning PacketType = "playerAlreadyMadeActionWarning"
	ActionIgnoredDueToDeadWarning  PacketType = "actionIgnoredDueToDeadWarning"
	SlowResponseWarning            PacketType = "slowResponseWarning"

	InvalidPacketTypeError  PacketType = "invalidPacketTypeError"
	InvalidPacketUsageError PacketType = "invalidPacketUsageError"
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
