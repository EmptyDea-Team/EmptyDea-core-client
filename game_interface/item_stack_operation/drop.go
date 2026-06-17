package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Drop 指示物品丢弃操作。
type Drop struct {
	Path  resources_control.SlotLocation
	Count uint8
}

func (Drop) ID() uint8 {
	return IDItemStackOperationDrop
}

func (Drop) CanInline() bool {
	return true
}

// DropInventoryItem 指示丢弃背包槽位物品的操作。
type DropInventoryItem struct {
	Slot  resources_control.SlotID
	Count uint8
}

func (DropInventoryItem) ID() uint8 {
	return IDItemStackOperationDrop
}

func (DropInventoryItem) CanInline() bool {
	return true
}

// DropDynamicContainerItem 指示丢弃动态容器槽位物品的操作。
type DropDynamicContainerItem struct {
	DynamicContainerID resources_control.DynamicContainerID
	Slot               resources_control.SlotID
	Count              uint8
}

func (DropDynamicContainerItem) ID() uint8 {
	return IDItemStackOperationDrop
}

func (DropDynamicContainerItem) CanInline() bool {
	return true
}

// DropContainerItem 指示丢弃已打开容器槽位物品的操作。
type DropContainerItem struct {
	Slot  resources_control.SlotID
	Count uint8
}

func (DropContainerItem) ID() uint8 {
	return IDItemStackOperationDrop
}

func (DropContainerItem) CanInline() bool {
	return true
}
