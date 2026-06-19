package game_interface

import (
	"context"

	resources_control_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
)

// UseItemOnBlocks 是机器人在使用手持物品对方块进行操作时的通用结构体。
type UseItemOnBlocks = game_interface_pb.UseItemOnBlocks

// BotClick 是远程点击操作实现。
type BotClick struct {
	client game_interface_pb.BotClickServiceClient
}

// ChangeSelectedHotbarSlot 将机器人当前选择的物品栏槽位切换为 hotbarSlotID。
func (b *BotClick) ChangeSelectedHotbarSlot(ctx context.Context, hotbarSlotID resources_control_client.SlotID) error {
	_, err := b.client.ChangeSelectedHotbarSlot(ctx, &game_interface_pb.ChangeSelectedHotbarSlotRequest{HotbarSlotID: uint32(hotbarSlotID)})
	return err
}

// ClickBlock 使用请求描述的手持物品操作点击方块。
func (b *BotClick) ClickBlock(ctx context.Context, request *game_interface_pb.UseItemOnBlocks) error {
	_, err := b.client.ClickBlock(ctx, &game_interface_pb.ClickBlockRequest{Request: request})
	return err
}

// ClickBlockWithPosition 在指定真实位置执行方块点击，并返回服务端是否强制移动机器人。
func (b *BotClick) ClickBlockWithPosition(ctx context.Context, request *game_interface_pb.UseItemOnBlocks, position *mgl32_pb.Vec3) (bool, error) {
	resp, err := b.client.ClickBlockWithPosition(ctx, &game_interface_pb.ClickBlockWithPositionRequest{Request: request, Position: position})
	if err != nil {
		return false, err
	}
	return resp.ForceMoved, nil
}

// ClickAir 使用指定快捷栏槽位执行空气点击。
func (b *BotClick) ClickAir(ctx context.Context, hotbarSlot resources_control_client.SlotID, realPosition *mgl32_pb.Vec3) error {
	_, err := b.client.ClickAir(ctx, &game_interface_pb.ClickAirRequest{HotbarSlot: uint32(hotbarSlot), RealPosition: realPosition})
	return err
}

// PlaceBlock 使用 request 描述的位置和物品，按 blockFace 指定方向放置方块。
func (b *BotClick) PlaceBlock(ctx context.Context, request *game_interface_pb.UseItemOnBlocks, blockFace int32) error {
	_, err := b.client.PlaceBlock(ctx, &game_interface_pb.PlaceBlockRequest{Request: request, BlockFace: blockFace})
	return err
}

// PlaceBlockHighLevel 执行高层方块放置，并返回点击位置、偏移位置和机器人位置。
func (b *BotClick) PlaceBlockHighLevel(ctx context.Context, blockPos *protocol_pb.BlockPos, hotBarSlot resources_control_client.SlotID, facing uint8) (*protocol_pb.BlockPos, *protocol_pb.BlockPos, *protocol_pb.BlockPos, error) {
	resp, err := b.client.PlaceBlockHighLevel(ctx, &game_interface_pb.PlaceBlockHighLevelRequest{
		BlockPos:   blockPos,
		HotbarSlot: uint32(hotBarSlot),
		Facing:     uint32(facing),
	})
	if err != nil {
		return nil, nil, nil, err
	}
	return resp.ClickPos, resp.OffsetPos, resp.BotPos, nil
}

// PickBlock 对指定方块执行选取方块操作，并返回是否成功和目标槽位。
func (b *BotClick) PickBlock(ctx context.Context, pos *protocol_pb.BlockPos, assignNBTData bool) (bool, resources_control_client.SlotID, error) {
	resp, err := b.client.PickBlock(ctx, &game_interface_pb.PickBlockRequest{Pos: pos, AssignNBTData: assignNBTData})
	if err != nil {
		return false, 0, err
	}
	return resp.Success, resources_control_client.SlotID(resp.Slot), nil
}
