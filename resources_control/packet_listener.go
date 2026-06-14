package resources_control

import (
	"context"
	"slices"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PacketListener 实现远程可撤销数据包监听。
type PacketListener struct {
	client resources_control_pb.PacketListenerServiceClient
}

// ListenPacket 监听数据包 ID 在 packetID 中的数据包。
func (p *PacketListener) ListenPacket(ctx context.Context, packetID []uint32, callback func(*packet_pb.Packet, error)) (uniqueID string, err error) {
	resp, err := p.client.CreatePacketListener(ctx, &resources_control_pb.CreatePacketListenerRequest{PacketIDs: packetID})
	if err != nil {
		return "", err
	}
	stream, err := p.client.ListenPacketEvents(ctx, &emptypb.Empty{})
	if err != nil {
		_, _ = p.client.DestroyPacketListener(context.Background(), &resources_control_pb.DestroyPacketListenerRequest{ListenerID: resp.ListenerID})
		return "", err
	}
	go func(listenerID string) {
		for {
			event, err := stream.Recv()
			if err != nil {
				callback(nil, err)
				return
			}
			if !slices.Contains(event.ListenerIDs, listenerID) {
				continue
			}
			callback(event.Packet, nil)
		}
	}(resp.ListenerID)
	return resp.ListenerID, nil
}

// DestroyListener 销毁唯一标识为 uniqueID 的数据包监听器。
func (p *PacketListener) DestroyListener(ctx context.Context, uniqueID string) (err error) {
	_, err = p.client.DestroyPacketListener(ctx, &resources_control_pb.DestroyPacketListenerRequest{ListenerID: uniqueID})
	return err
}
