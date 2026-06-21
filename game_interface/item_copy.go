package game_interface

import (
	"context"

	game_interface_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

type (
	// ItemType 指示该物品在单次操作中的物品类型。
	ItemType = game_interface_api.ItemType
	// ItemInfo 是物品的信息。
	ItemInfo = game_interface_api.ItemInfo
	// ItemInfoWithSlot 是物品的信息，同时指示该物品位于哪个槽位。
	ItemInfoWithSlot = game_interface_api.ItemInfoWithSlot
)

// ItemCopy 是远程物品拷贝实现。
type ItemCopy struct {
	client game_interface_pb.ItemCopyServiceClient
}

// CopyItem 将 baseItems 描述的基础物品复制为 targetItems。
func (i *ItemCopy) CopyItem(ctx context.Context, containerInfo *UseItemOnBlocks, baseItems []ItemInfoWithSlot, targetItems []*ItemInfo) error {
	_, err := i.client.CopyItem(ctx, &game_interface_pb.CopyItemRequest{
		ContainerInfo: containerInfo,
		BaseItems:     itemInfoWithSlotsToProto(baseItems),
		TargetItems:   itemInfoPointersToProto(targetItems),
	})
	return err
}

func itemInfoWithSlotsToProto(src []ItemInfoWithSlot) []*game_interface_pb.ItemInfoWithSlot {
	result := make([]*game_interface_pb.ItemInfoWithSlot, 0, len(src))
	for _, item := range src {
		result = append(result, &game_interface_pb.ItemInfoWithSlot{
			Slot: uint32(item.Slot),
			Item: &game_interface_pb.ItemInfo{
				Count:    uint32(item.ItemInfo.Count),
				ItemType: uint32(item.ItemInfo.ItemType),
			},
		})
	}
	return result
}

func itemInfoPointersToProto(src []*ItemInfo) []*game_interface_pb.ItemInfo {
	result := make([]*game_interface_pb.ItemInfo, 0, len(src))
	for _, item := range src {
		if item == nil {
			result = append(result, nil)
			continue
		}
		result = append(result, &game_interface_pb.ItemInfo{
			Count:    uint32(item.Count),
			ItemType: uint32(item.ItemType),
		})
	}
	return result
}
