package resources_control

import (
	"context"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// Packet 是远程数据包发送与编解码接口。
type Packet struct {
	client resources_control_pb.PacketServiceClient
}

// WritePacket 用于向租赁服发送数据包 p。
func (p *Packet) WritePacket(ctx context.Context, pk *packet_pb.Packet) (err error) {
	_, err = p.client.SendPacket(ctx, &resources_control_pb.SendPacketRequest{Packet: pk})
	return err
}

// EncodePacket 将结构化数据包编码为 Minecraft payload。
func (p *Packet) EncodePacket(ctx context.Context, pk *packet_pb.Packet, protocolVersion int32, shieldID int32) (packetID uint32, payload []byte, err error) {
	resp, err := p.client.EncodePacket(ctx, &resources_control_pb.EncodePacketRequest{
		Packet:   pk,
		Protocol: protocolVersion,
		ShieldID: shieldID,
	})
	if err != nil {
		return 0, nil, err
	}
	return resp.PacketID, resp.Payload, nil
}

// DecodePacket 将 Minecraft payload 解码为结构化数据包。
func (p *Packet) DecodePacket(ctx context.Context, packetID uint32, payload []byte, protocolVersion int32, shieldID int32) (pk *packet_pb.Packet, err error) {
	resp, err := p.client.DecodePacket(ctx, &resources_control_pb.DecodePacketRequest{
		PacketID: packetID,
		Payload:  payload,
		Protocol: protocolVersion,
		ShieldID: shieldID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Packet, nil
}
