package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// CreativeItem 指示创造物品获取操作。
type CreativeItem struct {
	// CreativeItemNetworkID 是创造物品网络 ID。
	CreativeItemNetworkID uint32
	// Path 是目标槽位位置。
	Path resources_control.SlotLocation
	// Count 是获取数量。
	Count uint8
}

// ID 返回创造物品获取操作编号。
func (CreativeItem) ID() uint8 {
	return IDItemStackOperationCreativeItem
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (CreativeItem) CanInline() bool {
	return false
}

// CreativeItemToInventory 指示从创造物品栏获取物品到背包槽位的操作。
type CreativeItemToInventory struct {
	// CreativeItemNetworkID 是创造物品网络 ID。
	CreativeItemNetworkID uint32
	// Slot 是背包目标槽位。
	Slot resources_control.SlotID
	// Count 是获取数量。
	Count uint8
}

// ID 返回创造物品获取操作编号。
func (CreativeItemToInventory) ID() uint8 {
	return IDItemStackOperationCreativeItem
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (CreativeItemToInventory) CanInline() bool {
	return false
}

// CreativeItemToDynamicContainer 指示从创造物品栏获取物品到动态容器槽位的操作。
type CreativeItemToDynamicContainer struct {
	// CreativeItemNetworkID 是创造物品网络 ID。
	CreativeItemNetworkID uint32
	// DynamicContainerID 是目标动态容器 ID。
	DynamicContainerID resources_control.DynamicContainerID
	// Slot 是动态容器目标槽位。
	Slot resources_control.SlotID
	// Count 是获取数量。
	Count uint8
}

// ID 返回创造物品获取操作编号。
func (CreativeItemToDynamicContainer) ID() uint8 {
	return IDItemStackOperationCreativeItem
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (CreativeItemToDynamicContainer) CanInline() bool {
	return false
}
