package game_interface

import (
	"context"

	resources_control_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

// ItemTransition 是远程物品状态转移实现。
type ItemTransition struct {
	client game_interface_pb.ItemTransitionServiceClient
}

func (h *ItemTransition) Transition(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot, srcWindowID resources_control_client.WindowID, dstWindowID resources_control_client.WindowID) (bool, error) {
	resp, err := h.client.Transition(ctx, &game_interface_pb.TransitionRequest{
		Src:         itemInfoWithSlotsToProto(src),
		Dst:         itemInfoWithSlotsToProto(dst),
		SrcWindowID: int32(srcWindowID),
		DstWindowID: int32(dstWindowID),
	})
	return transitionResult(resp, err)
}

func (h *ItemTransition) TransitionBetweenInventory(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error) {
	resp, err := h.client.TransitionBetweenInventory(ctx, simpleTransition(src, dst))
	return transitionResult(resp, err)
}

func (h *ItemTransition) TransitionBetweenContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error) {
	resp, err := h.client.TransitionBetweenContainer(ctx, simpleTransition(src, dst))
	return transitionResult(resp, err)
}

func (h *ItemTransition) TransitionToContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error) {
	resp, err := h.client.TransitionToContainer(ctx, simpleTransition(src, dst))
	return transitionResult(resp, err)
}

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
