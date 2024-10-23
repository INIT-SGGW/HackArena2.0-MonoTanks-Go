package packets

import "hackarena2-0-mono-tanks-go/packet"

type ReadyToReceiveGameState struct {
}

func (r *ReadyToReceiveGameState) ToPacket() packet.Packet {
	return packet.Packet{
		Type: packet.ReadyToReceiveGameState,
	}
}
