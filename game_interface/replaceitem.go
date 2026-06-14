package game_interface

import (
	"context"

	resources_control_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// ReplaceitemPath 指示 replaceitem 时目标物品栏的槽位类型。
type ReplaceitemPath string

const (
	ReplacePathInventoryOnly ReplaceitemPath = "slot.inventory"
	ReplacePathHotbarOnly    ReplaceitemPath = "slot.hotbar"
	ReplacePathInventory     ReplaceitemPath = "slot.inventory | slot.hotbar"
)

// ReplaceitemInfo 指示要通过 replaceitem 生成的物品的基本信息。
type ReplaceitemInfo struct {
	Name     string
	Count    uint8
	MetaData int16
	Slot     resources_control_client.SlotID
}

// Replaceitem 是远程 replaceitem 实现。
type Replaceitem struct {
	client game_interface_pb.ReplaceitemServiceClient
}

func (r *Replaceitem) ReplaceitemInInventory(ctx context.Context, target string, path ReplaceitemPath, itemInfo ReplaceitemInfo, method string, blocked bool) error {
	_, err := r.client.ReplaceitemInInventory(ctx, &game_interface_pb.ReplaceitemInInventoryRequest{
		Target: target,
		Path:   string(path),
		ItemInfo: &game_interface_pb.ReplaceitemInfo{
			Name:     itemInfo.Name,
			Count:    uint32(itemInfo.Count),
			MetaData: int32(itemInfo.MetaData),
			Slot:     uint32(itemInfo.Slot),
		},
		Method:  method,
		Blocked: blocked,
	})
	return err
}

func (r *Replaceitem) ReplaceitemInContainerAsync(ctx context.Context, blockPos *protocol_pb.BlockPos, itemInfo ReplaceitemInfo, method string) error {
	_, err := r.client.ReplaceitemInContainerAsync(ctx, &game_interface_pb.ReplaceitemInContainerAsyncRequest{
		BlockPos: blockPos,
		ItemInfo: &game_interface_pb.ReplaceitemInfo{
			Name:     itemInfo.Name,
			Count:    uint32(itemInfo.Count),
			MetaData: int32(itemInfo.MetaData),
			Slot:     uint32(itemInfo.Slot),
		},
		Method: method,
	})
	return err
}
