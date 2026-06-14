package game_interface

import (
	"context"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

// Movement 是远程移动控制实现。
type Movement struct {
	client game_interface_pb.MovementServiceClient
}

func (m *Movement) StartFlying(ctx context.Context) error {
	_, err := m.client.StartFlying(ctx, &game_interface_pb.StartFlyingRequest{})
	return err
}

func (m *Movement) StopFlying(ctx context.Context) error {
	_, err := m.client.StopFlying(ctx, &game_interface_pb.StopFlyingRequest{})
	return err
}
