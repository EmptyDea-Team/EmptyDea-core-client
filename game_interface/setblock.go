package game_interface

import (
	"context"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// SetBlock 是远程 setblock 实现。
type SetBlock struct {
	client game_interface_pb.SetBlockServiceClient
}

// SetBlock 在 pos 处同步放置名称为 name、方块状态为 states 的方块。
func (s *SetBlock) SetBlock(ctx context.Context, pos *protocol_pb.BlockPos, name string, states string) error {
	_, err := s.client.SetBlock(ctx, &game_interface_pb.SetBlockRequest{Pos: pos, Name: name, States: states})
	return err
}

// SetBlockAsync 在 pos 处异步放置名称为 name、方块状态为 states 的方块。
func (s *SetBlock) SetBlockAsync(ctx context.Context, pos *protocol_pb.BlockPos, name string, states string) error {
	_, err := s.client.SetBlockAsync(ctx, &game_interface_pb.SetBlockRequest{Pos: pos, Name: name, States: states})
	return err
}

// SetAnvil 在 pos 处放置铁砧，并返回最终使用的方块状态。
func (s *SetBlock) SetAnvil(ctx context.Context, pos *protocol_pb.BlockPos, placeBaseBlock bool) (map[string]any, error) {
	resp, err := s.client.SetAnvil(ctx, &game_interface_pb.SetAnvilRequest{Pos: pos, PlaceBaseBlock: placeBaseBlock})
	if err != nil {
		return nil, err
	}
	result := make(map[string]any, len(resp.States))
	for key, value := range resp.States {
		result[key] = value.AsInterface()
	}
	return result, nil
}
