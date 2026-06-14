package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Renaming 指示基于铁砧的物品重命名操作。
type Renaming struct {
	Path    resources_control.SlotLocation
	NewName string
}

func (Renaming) ID() uint8 {
	return IDItemStackOperationHighLevelRenaming
}

func (Renaming) CanInline() bool {
	return false
}

// RenameInventoryItem 指示重命名背包槽位物品的操作。
type RenameInventoryItem struct {
	Slot    resources_control.SlotID
	NewName string
}

func (RenameInventoryItem) ID() uint8 {
	return IDItemStackOperationHighLevelRenaming
}

func (RenameInventoryItem) CanInline() bool {
	return false
}
