package game_interface

import (
	"context"

	resources_control_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
)

// ItemTransition 是远程物品状态转移实现。
type ItemTransition struct {
	client game_interface_pb.ItemTransitionServiceClient
}

// Transition 在指定源窗口和目标窗口之间执行物品状态转移。
func (h *ItemTransition) Transition(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot, srcWindowName resources_control_client.WindowName, dstWindowName resources_control_client.WindowName) (bool, error) {
	resp, err := h.client.Transition(ctx, &game_interface_pb.TransitionRequest{
		Src:           itemInfoWithSlotsToProto(src),
		Dst:           itemInfoWithSlotsToProto(dst),
		SrcWindowName: windowNameToProto(srcWindowName),
		DstWindowName: windowNameToProto(dstWindowName),
	})
	return transitionResult(resp, err)
}

// TransitionBetweenInventory 在背包内执行物品状态转移。
func (h *ItemTransition) TransitionBetweenInventory(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error) {
	resp, err := h.client.TransitionBetweenInventory(ctx, simpleTransition(src, dst))
	return transitionResult(resp, err)
}

// TransitionBetweenContainer 在已打开容器内执行物品状态转移。
func (h *ItemTransition) TransitionBetweenContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error) {
	resp, err := h.client.TransitionBetweenContainer(ctx, simpleTransition(src, dst))
	return transitionResult(resp, err)
}

// TransitionToContainer 将物品从背包转移到已打开容器。
func (h *ItemTransition) TransitionToContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error) {
	resp, err := h.client.TransitionToContainer(ctx, simpleTransition(src, dst))
	return transitionResult(resp, err)
}

// TransitionToInventory 将物品从已打开容器转移到背包。
func (h *ItemTransition) TransitionToInventory(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error) {
	resp, err := h.client.TransitionToInventory(ctx, simpleTransition(src, dst))
	return transitionResult(resp, err)
}

func simpleTransition(src []ItemInfoWithSlot, dst []ItemInfoWithSlot) *game_interface_pb.SimpleTransitionRequest {
	return &game_interface_pb.SimpleTransitionRequest{
		Src: itemInfoWithSlotsToProto(src),
		Dst: itemInfoWithSlotsToProto(dst),
	}
}

func transitionResult(resp *game_interface_pb.TransitionResponse, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func windowNameToProto(src resources_control_client.WindowName) *resources_control_pb.WindowName {
	return &resources_control_pb.WindowName{
		WindowID:           int32(src.WindowID),
		DynamicContainerID: uint32(src.DynamicContainerID),
	}
}
