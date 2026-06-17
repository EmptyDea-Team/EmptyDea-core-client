package resources_control

import (
	"context"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// Inventory 描述远程单个库存。
type Inventory struct {
	windowName WindowName
	client     resources_control_pb.InventoryServiceClient
}

// GetItemStack 返回当前库存中物品栏编号为 slotID 的物品堆栈信息。
func (i *Inventory) GetItemStack(ctx context.Context, slotID SlotID) (item *protocol_pb.ItemInstance, err error) {
	item, _, err = (&Inventories{client: i.client}).GetItemStack(ctx, i.windowName, slotID)
	return item, err
}

// GetAllItemStack 列出当前库存的全部物品堆栈实例。
func (i *Inventory) GetAllItemStack(ctx context.Context) (mapping map[SlotID]*protocol_pb.ItemInstance, err error) {
	mapping, _, err = (&Inventories{client: i.client}).GetAllItemStack(ctx, i.windowName)
	return mapping, err
}

// Inventories 描述远程所有库存。
type Inventories struct {
	client resources_control_pb.InventoryServiceClient
}

// GetInventory 返回窗口名为 windowName 的库存。
func (i *Inventories) GetInventory(ctx context.Context, windowName WindowName) (inventory *Inventory, existed bool, err error) {
	resp, err := i.client.GetInventory(ctx, &resources_control_pb.GetInventoryRequest{WindowName: windowNameToProto(windowName)})
	if err != nil {
		return nil, false, err
	}
	if !resp.Found {
		return nil, false, nil
	}
	return &Inventory{windowName: windowName, client: i.client}, true, nil
}

// GetItemStack 加载位于 windowName 的库存中索引为 slotID 的物品。
func (i *Inventories) GetItemStack(ctx context.Context, windowName WindowName, slotID SlotID) (item *protocol_pb.ItemInstance, inventoryExisted bool, err error) {
	resp, err := i.client.GetItemStack(ctx, &resources_control_pb.GetItemStackRequest{
		WindowName: windowNameToProto(windowName),
		SlotID:     uint32(slotID),
	})
	if err != nil {
		return nil, false, err
	}
	if !resp.InventoryFound {
		return nil, false, nil
	}
	if !resp.ItemFound || resp.Item == nil {
		return NewAirItemInstance(), true, nil
	}
	return resp.Item, true, nil
}

// GetAllItemStack 列出窗口名为 windowName 的库存中的所有物品堆栈实例。
func (i *Inventories) GetAllItemStack(ctx context.Context, windowName WindowName) (mapping map[SlotID]*protocol_pb.ItemInstance, inventoryExisted bool, err error) {
	resp, err := i.client.GetAllItemStacks(ctx, &resources_control_pb.GetAllItemStacksRequest{WindowName: windowNameToProto(windowName)})
	if err != nil {
		return nil, false, err
	}
	if !resp.InventoryFound {
		return nil, false, nil
	}
	mapping = make(map[SlotID]*protocol_pb.ItemInstance, len(resp.Items))
	for _, item := range resp.Items {
		if item != nil && item.Item != nil {
			mapping[SlotID(item.SlotID)] = item.Item
		}
	}
	return mapping, true, nil
}

// GetAllWindowName 列出当前所有库存的窗口名。
func (i *Inventories) GetAllWindowName(ctx context.Context) (result []WindowName, err error) {
	resp, err := i.client.ListWindowNames(ctx, &resources_control_pb.ListWindowNamesRequest{})
	if err != nil {
		return nil, err
	}
	result = make([]WindowName, 0, len(resp.WindowNames))
	for _, windowName := range resp.WindowNames {
		result = append(result, windowNameFromProto(windowName))
	}
	return result, nil
}

// NewAirItemInstance 返回一个新的空气物品堆栈实例。
func NewAirItemInstance() (item *protocol_pb.ItemInstance) {
	return &protocol_pb.ItemInstance{}
}

func windowNameToProto(src WindowName) *resources_control_pb.WindowName {
	return &resources_control_pb.WindowName{
		WindowID:           int32(src.WindowID),
		DynamicContainerID: uint32(src.DynamicContainerID),
	}
}

func windowNameFromProto(src *resources_control_pb.WindowName) WindowName {
	if src == nil {
		return WindowName{}
	}
	return NewWindowName(WindowID(src.WindowID), DynamicContainerID(src.DynamicContainerID))
}
