package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Drop 指示物品丢弃操作。
type Drop struct {
	// Path 是要丢弃物品的位置。
	Path resources_control.SlotLocation
	// Count 是丢弃数量。
	Count uint8
}

// ID 返回物品丢弃操作编号。
func (Drop) ID() uint8 {
	return IDItemStackOperationDrop
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (Drop) CanInline() bool {
	return true
}

// DropInventoryItem 指示丢弃背包槽位物品的操作。
type DropInventoryItem struct {
	// Slot 是背包内要丢弃物品的槽位。
	Slot resources_control.SlotID
	// Count 是丢弃数量。
	Count uint8
}

// ID 返回物品丢弃操作编号。
func (DropInventoryItem) ID() uint8 {
	return IDItemStackOperationDrop
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (DropInventoryItem) CanInline() bool {
	return true
}

// DropDynamicContainerItem 指示丢弃动态容器槽位物品的操作。
type DropDynamicContainerItem struct {
	// DynamicContainerID 是动态容器 ID。
	DynamicContainerID resources_control.DynamicContainerID
	// Slot 是动态容器内要丢弃物品的槽位。
	Slot resources_control.SlotID
	// Count 是丢弃数量。
	Count uint8
}

// ID 返回物品丢弃操作编号。
func (DropDynamicContainerItem) ID() uint8 {
	return IDItemStackOperationDrop
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (DropDynamicContainerItem) CanInline() bool {
	return true
}

// DropContainerItem 指示丢弃已打开容器槽位物品的操作。
type DropContainerItem struct {
	// Slot 是已打开容器内要丢弃物品的槽位。
	Slot resources_control.SlotID
	// Count 是丢弃数量。
	Count uint8
}

// ID 返回物品丢弃操作编号。
func (DropContainerItem) ID() uint8 {
	return IDItemStackOperationDrop
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (DropContainerItem) CanInline() bool {
	return true
}
