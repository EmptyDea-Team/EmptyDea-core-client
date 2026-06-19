package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Move 指示物品移动操作。
type Move struct {
	// Source 是源物品位置。
	Source resources_control.SlotLocation
	// Destination 是目标物品位置。
	Destination resources_control.SlotLocation
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (Move) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (Move) CanInline() bool {
	return true
}

// MoveBetweenInventory 指示背包内两个槽位之间的物品移动操作。
type MoveBetweenInventory struct {
	// Source 是背包源槽位。
	Source resources_control.SlotID
	// Destination 是背包目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveBetweenInventory) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveBetweenInventory) CanInline() bool {
	return true
}

// MoveBetweenDynamicContainer 指示动态容器内两个槽位之间的物品移动操作。
type MoveBetweenDynamicContainer struct {
	// DynamicContainerID 是动态容器 ID。
	DynamicContainerID resources_control.DynamicContainerID
	// Source 是动态容器源槽位。
	Source resources_control.SlotID
	// Destination 是动态容器目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveBetweenDynamicContainer) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveBetweenDynamicContainer) CanInline() bool {
	return true
}

// MoveToDynamicContainer 指示从背包移动物品到动态容器的操作。
type MoveToDynamicContainer struct {
	// Source 是背包源槽位。
	Source resources_control.SlotID
	// DynamicContainerID 是动态容器 ID。
	DynamicContainerID resources_control.DynamicContainerID
	// Destination 是动态容器目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveToDynamicContainer) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveToDynamicContainer) CanInline() bool {
	return true
}

// MoveFromDynamicContainer 指示从动态容器移动物品到背包的操作。
type MoveFromDynamicContainer struct {
	// DynamicContainerID 是动态容器 ID。
	DynamicContainerID resources_control.DynamicContainerID
	// Source 是动态容器源槽位。
	Source resources_control.SlotID
	// Destination 是背包目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveFromDynamicContainer) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveFromDynamicContainer) CanInline() bool {
	return true
}

// MoveBetweenContainer 指示已打开容器内两个槽位之间的物品移动操作。
type MoveBetweenContainer struct {
	// Source 是已打开容器源槽位。
	Source resources_control.SlotID
	// Destination 是已打开容器目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveBetweenContainer) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveBetweenContainer) CanInline() bool {
	return true
}

// MoveToContainer 指示从背包移动物品到已打开容器的操作。
type MoveToContainer struct {
	// Source 是背包源槽位。
	Source resources_control.SlotID
	// Destination 是已打开容器目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveToContainer) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveToContainer) CanInline() bool {
	return true
}

// MoveToInventory 指示从已打开容器移动物品到背包的操作。
type MoveToInventory struct {
	// Source 是已打开容器源槽位。
	Source resources_control.SlotID
	// Destination 是背包目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveToInventory) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveToInventory) CanInline() bool {
	return true
}

// MoveToCraftingTable 指示从背包移动物品到合成栏的操作。
type MoveToCraftingTable struct {
	// Source 是背包源槽位。
	Source resources_control.SlotID
	// Destination 是合成栏目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveToCraftingTable) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveToCraftingTable) CanInline() bool {
	return true
}

// MoveFromCraftingTable 指示从合成栏移动物品到背包的操作。
type MoveFromCraftingTable struct {
	// Source 是合成栏源槽位。
	Source resources_control.SlotID
	// Destination 是背包目标槽位。
	Destination resources_control.SlotID
	// Count 是移动数量。
	Count uint8
}

// ID 返回物品移动操作编号。
func (MoveFromCraftingTable) ID() uint8 {
	return IDItemStackOperationMove
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (MoveFromCraftingTable) CanInline() bool {
	return true
}
