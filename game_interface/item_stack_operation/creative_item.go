package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// CreativeItem 指示创造物品获取操作。
type CreativeItem struct {
	CreativeItemNetworkID uint32
	Path                  resources_control.SlotLocation
	Count                 uint8
}

func (CreativeItem) ID() uint8 {
	return IDItemStackOperationCreativeItem
}

func (CreativeItem) CanInline() bool {
	return false
}

// CreativeItemToInventory 指示从创造物品栏获取物品到背包槽位的操作。
type CreativeItemToInventory struct {
	CreativeItemNetworkID uint32
	Slot                  resources_control.SlotID
	Count                 uint8
}

func (CreativeItemToInventory) ID() uint8 {
	return IDItemStackOperationCreativeItem
}

func (CreativeItemToInventory) CanInline() bool {
	return false
}
