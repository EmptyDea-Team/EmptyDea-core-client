package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Swap 指示物品交换操作。
type Swap struct {
	// Source 是源物品位置。
	Source resources_control.SlotLocation
	// Destination 是目标物品位置。
	Destination resources_control.SlotLocation
}

// ID 返回物品交换操作编号。
func (Swap) ID() uint8 {
	return IDItemStackOperationSwap
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (Swap) CanInline() bool {
	return true
}

// SwapBetweenInventory 指示背包内两个槽位之间的物品交换操作。
type SwapBetweenInventory struct {
	// Source 是背包源槽位。
	Source resources_control.SlotID
	// Destination 是背包目标槽位。
	Destination resources_control.SlotID
}

// ID 返回物品交换操作编号。
func (SwapBetweenInventory) ID() uint8 {
	return IDItemStackOperationSwap
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (SwapBetweenInventory) CanInline() bool {
	return true
}

// SwapBetweenDynamicContainer 指示同一动态容器内两个槽位之间的物品交换操作。
type SwapBetweenDynamicContainer struct {
	// DynamicContainerID 是动态容器 ID。
	DynamicContainerID resources_control.DynamicContainerID
	// Source 是动态容器源槽位。
	Source resources_control.SlotID
	// Destination 是动态容器目标槽位。
	Destination resources_control.SlotID
}

// ID 返回物品交换操作编号。
func (SwapBetweenDynamicContainer) ID() uint8 {
	return IDItemStackOperationSwap
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (SwapBetweenDynamicContainer) CanInline() bool {
	return true
}

// SwapInventoryBetweenDynamicContainer 指示背包和动态容器之间的物品交换操作。
type SwapInventoryBetweenDynamicContainer struct {
	// Source 是背包源槽位。
	Source resources_control.SlotID
	// DynamicContainerID 是动态容器 ID。
	DynamicContainerID resources_control.DynamicContainerID
	// Destination 是动态容器目标槽位。
	Destination resources_control.SlotID
}

// ID 返回物品交换操作编号。
func (SwapInventoryBetweenDynamicContainer) ID() uint8 {
	return IDItemStackOperationSwap
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (SwapInventoryBetweenDynamicContainer) CanInline() bool {
	return true
}

// SwapInventoryBetweenContainer 指示背包和已打开容器之间的物品交换操作。
type SwapInventoryBetweenContainer struct {
	// Source 是背包源槽位。
	Source resources_control.SlotID
	// Destination 是已打开容器目标槽位。
	Destination resources_control.SlotID
}

// ID 返回物品交换操作编号。
func (SwapInventoryBetweenContainer) ID() uint8 {
	return IDItemStackOperationSwap
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (SwapInventoryBetweenContainer) CanInline() bool {
	return true
}
