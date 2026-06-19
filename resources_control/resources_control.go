package resources_control

import (
	"context"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
	uqholder_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control/uqholder"
	"google.golang.org/grpc"
)

// BotInfo 是当前机器人基础登录信息。
type BotInfo struct {
	// BotName 是机器人玩家名。
	BotName string
	// XUID 是机器人 Xbox 用户 ID。
	XUID string
	// EntityUniqueID 是机器人实体唯一 ID。
	EntityUniqueID int64
	// EntityRuntimeID 是机器人实体运行时 ID。
	EntityRuntimeID uint64
}

// Resources 是面向 core/resources_control 风格的远程资源中心。
type Resources struct {
	resources      *ResourcesService
	inventories    *Inventories
	container      *ContainerManager
	constantPacket *ConstantPacket
	packetListener *PacketListener
	uqholder       *UQHolder
}

// New 创建资源层客户端集合。
func New(conn grpc.ClientConnInterface) *Resources {
	return &Resources{
		resources:      &ResourcesService{client: resources_control_pb.NewResourcesServiceClient(conn)},
		inventories:    &Inventories{client: resources_control_pb.NewInventoryServiceClient(conn)},
		container:      &ContainerManager{client: resources_control_pb.NewContainerResourceServiceClient(conn)},
		constantPacket: &ConstantPacket{client: resources_control_pb.NewConstantPacketServiceClient(conn)},
		packetListener: &PacketListener{client: resources_control_pb.NewPacketListenerServiceClient(conn)},
		uqholder:       uqholder_client.New(resources_control_pb.NewUQHolderClient(conn)),
	}
}

// BotInfo 返回机器人的基本信息。
func (r *Resources) BotInfo(ctx context.Context) (info BotInfo, err error) {
	return r.resources.BotInfo(ctx)
}

// WritePacket 用于向服务端发送数据包 p。
func (r *Resources) WritePacket(ctx context.Context, p *packet_pb.Packet) (err error) {
	return r.resources.WritePacket(ctx, p)
}

// Inventories 返回库存的相关资源。
func (r *Resources) Inventories() (inventories *Inventories) {
	return r.inventories
}

// Container 返回容器的相关资源。
func (r *Resources) Container() (container *ContainerManager) {
	return r.container
}

// PacketListener 返回一个可撤销的数据包监听实现。
func (r *Resources) PacketListener() (packetListener *PacketListener) {
	return r.packetListener
}

// ConstantPacket 返回常量数据包的有关实现。
func (r *Resources) ConstantPacket() (constantPacket *ConstantPacket) {
	return r.constantPacket
}

// UQHolder 返回轻量状态集合。
func (r *Resources) UQHolder() (uqholder *UQHolder) {
	return r.uqholder
}

// ResourcesService 封装基础资源查询。
type ResourcesService struct {
	client resources_control_pb.ResourcesServiceClient
}

// BotInfo 返回机器人的基本信息。
func (r *ResourcesService) BotInfo(ctx context.Context) (info BotInfo, err error) {
	resp, err := r.client.GetBotInfo(ctx, &resources_control_pb.GetBotInfoRequest{})
	if err != nil {
		return BotInfo{}, err
	}
	return BotInfo{
		BotName:         resp.BotName,
		XUID:            resp.XUID,
		EntityUniqueID:  resp.EntityUniqueID,
		EntityRuntimeID: resp.EntityRuntimeID,
	}, nil
}

// WritePacket 用于向服务端发送数据包 p。
func (r *ResourcesService) WritePacket(ctx context.Context, p *packet_pb.Packet) (err error) {
	_, err = r.client.WritePacket(ctx, &resources_control_pb.WritePacketRequest{Packet: p})
	return err
}
