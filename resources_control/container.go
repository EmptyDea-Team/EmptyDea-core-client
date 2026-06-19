package resources_control

import (
	"context"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// ContainerID 是容器 ID。
type ContainerID uint8

// ContainerManager 描述远程容器资源。
type ContainerManager struct {
	client resources_control_pb.ContainerResourceServiceClient
}

// States 返回已打开容器的状态。
func (c *ContainerManager) States(ctx context.Context) (state uint8, err error) {
	resp, err := c.client.GetContainerState(ctx, &resources_control_pb.GetContainerStateRequest{})
	if err != nil {
		return 0, err
	}
	return uint8(resp.State), nil
}

// ContainerData 获取当前已打开容器的状态。
func (c *ContainerManager) ContainerData(ctx context.Context) (data *packet_pb.ContainerOpen, containerID ContainerID, existed bool, err error) {
	resp, err := c.client.GetContainerData(ctx, &resources_control_pb.GetContainerDataRequest{})
	if err != nil {
		return nil, 0, false, err
	}
	if !resp.Found || resp.Data == nil {
		return nil, 0, false, nil
	}
	return resp.Data, ContainerID(resp.ContainerID), true, nil
}
