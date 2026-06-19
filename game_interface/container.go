package game_interface

import (
	"context"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

// ContainerOpenAndClose 是远程容器打开和关闭控制系统。
type ContainerOpenAndClose struct {
	client game_interface_pb.ContainerOpenAndCloseServiceClient
}

// OpenContainer 打开 container 指定的容器。
func (c *ContainerOpenAndClose) OpenContainer(ctx context.Context, container *UseItemOnBlocks, changeToTargetSlot bool) (bool, error) {
	resp, err := c.client.OpenContainer(ctx, &game_interface_pb.OpenContainerRequest{Container: container, ChangeToTargetSlot: changeToTargetSlot})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

// OpenInventory 打开机器人背包。
func (c *ContainerOpenAndClose) OpenInventory(ctx context.Context) (bool, error) {
	resp, err := c.client.OpenInventory(ctx, &game_interface_pb.OpenInventoryRequest{})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

// CloseContainer 关闭当前已打开的容器。
func (c *ContainerOpenAndClose) CloseContainer(ctx context.Context) error {
	_, err := c.client.CloseContainer(ctx, &game_interface_pb.CloseContainerRequest{})
	return err
}
