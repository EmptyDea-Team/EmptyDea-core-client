package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Crafting 指示合成操作。
type Crafting struct {
	// RecipeNetworkID 是合成配方网络 ID。
	RecipeNetworkID uint32
	// ResultSlotID 是结果物品所在槽位。
	ResultSlotID resources_control.SlotID
	// ResultCount 是结果物品数量。
	ResultCount uint8
	// ResultItem 是结果物品期望更新后的数据。
	ResultItem ExpectedNewItem
}

// ID 返回合成操作编号。
func (Crafting) ID() uint8 {
	return IDItemStackOperationHighLevelCrafting
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (Crafting) CanInline() bool {
	return false
}
