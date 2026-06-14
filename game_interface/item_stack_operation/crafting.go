package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Crafting 指示合成操作。
type Crafting struct {
	RecipeNetworkID uint32
	ResultSlotID    resources_control.SlotID
	ResultCount     uint8
	ResultItem      ExpectedNewItem
}

func (Crafting) ID() uint8 {
	return IDItemStackOperationHighLevelCrafting
}

func (Crafting) CanInline() bool {
	return false
}
