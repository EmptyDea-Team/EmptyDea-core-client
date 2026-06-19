package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Renaming 指示基于铁砧的物品重命名操作。
type Renaming struct {
	// Path 是要重命名物品所在位置。
	Path resources_control.SlotLocation
	// NewName 是物品新名称。
	NewName string
}

// ID 返回重命名操作编号。
func (Renaming) ID() uint8 {
	return IDItemStackOperationHighLevelRenaming
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (Renaming) CanInline() bool {
	return false
}

// RenameInventoryItem 指示重命名背包槽位物品的操作。
type RenameInventoryItem struct {
	// Slot 是背包内要重命名物品的槽位。
	Slot resources_control.SlotID
	// NewName 是物品新名称。
	NewName string
}

// ID 返回重命名操作编号。
func (RenameInventoryItem) ID() uint8 {
	return IDItemStackOperationHighLevelRenaming
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (RenameInventoryItem) CanInline() bool {
	return false
}
