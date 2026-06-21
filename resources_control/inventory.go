package resources_control

import (
	"context"

	"github.com/EmptyDea-Team/EmptyDea-core-api/define"
	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

type (
	// SlotID 是单个物品栏槽位的索引，它是从 0 开始索引的。
	SlotID = define.SlotID
	// WindowID 是机器人已打开(或持有)的库存的窗口 ID。
	WindowID = define.WindowID
	// DynamicContainerID 是机器人已打开(或持有)的动态库存的容器 ID。
	DynamicContainerID = define.DynamicContainerID
	// WindowName 唯一标识一个普通窗口或动态容器窗口。
	WindowName = define.WindowName
	// SlotLocation 描述一个物品的所在的位置。
	SlotLocation = define.SlotLocation
)

const (
	// WindowIDInventory 是玩家背包窗口 ID。
	WindowIDInventory = define.WindowIDInventory
	// WindowIDOffHand 是副手窗口 ID。
	WindowIDOffHand = define.WindowIDOffHand
	// WindowIDArmour 是盔甲栏窗口 ID。
	WindowIDArmour = define.WindowIDArmour
	// WindowIDCrafting 是合成栏窗口 ID。
	WindowIDCrafting = define.WindowIDCrafting
	// WindowIDUI 是 UI 窗口 ID。
	WindowIDUI = define.WindowIDUI
	// WindowIDDynamic 是动态容器窗口 ID。
	WindowIDDynamic = define.WindowIDDynamic
)

var (
	// WindowNameInventory 是玩家背包窗口名。
	WindowNameInventory = define.WindowNameInventory
	// WindowNameOffHand 是副手窗口名。
	WindowNameOffHand = define.WindowNameOffHand
	// WindowNameArmour 是盔甲栏窗口名。
	WindowNameArmour = define.WindowNameArmour
	// WindowNameCrafting 是合成栏窗口名。
	WindowNameCrafting = define.WindowNameCrafting
	// WindowNameUI 是 UI 窗口名。
	WindowNameUI = define.WindowNameUI
)

// NewDynamicContainerWindowName 基于动态容器 ID 构造动态容器窗口名。
func NewDynamicContainerWindowName(dynamicContainerID DynamicContainerID) WindowName {
	return define.NewDynamicContainerWindowName(dynamicContainerID)
}

// NewWindowName 基于窗口 ID 和动态容器 ID 构造窗口名。
func NewWindowName(windowID WindowID, dynamicContainerID DynamicContainerID) WindowName {
	return define.NewWindowName(windowID, dynamicContainerID)
}

// Inventory 描述远程单个库存。
type Inventory struct {
	windowName  WindowName
	inventories *Inventories
}

// GetItemStack 返回当前库存中物品栏编号为 slotID 的物品堆栈信息。
func (i *Inventory) GetItemStack(ctx context.Context, slotID SlotID) (item *protocol_pb.ItemInstance, err error) {
	item, _, err = i.inventories.GetItemStack(ctx, i.windowName, slotID)
	return item, err
}

// GetAllItemStack 列出当前库存的全部物品堆栈实例。
func (i *Inventory) GetAllItemStack(ctx context.Context) (mapping map[SlotID]*protocol_pb.ItemInstance, err error) {
	mapping, _, err = i.inventories.GetAllItemStack(ctx, i.windowName)
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
	return &Inventory{windowName: windowName, inventories: i}, true, nil
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
