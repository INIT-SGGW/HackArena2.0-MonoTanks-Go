package packets

import "hack-arena-2024-h2-go/packet"

type ReadyToReceiveGameState struct {
}

func (r *ReadyToReceiveGameState) ToPacket() packet.Packet {
	return packet.Packet{
		Type: packet.ReadyToReceiveGameState,
	}
}
