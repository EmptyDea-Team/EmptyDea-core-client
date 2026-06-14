package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Trimming 指示锻造台纹饰操作。
type Trimming struct {
	TrimItem   resources_control.SlotLocation
	Material   resources_control.SlotLocation
	Template   resources_control.SlotLocation
	ResultItem ExpectedNewItem
}

func (Trimming) ID() uint8 {
	return IDItemStackOperationHighLevelTrimming
}

func (Trimming) CanInline() bool {
	return false
}

// TrimmingFromInventory 指示使用背包内物品执行锻造台纹饰的操作。
type TrimmingFromInventory struct {
	TrimItemSlot resources_control.SlotID
	MaterialSlot resources_control.SlotID
	TemplateSlot resources_control.SlotID
	ResultItem   ExpectedNewItem
}

func (TrimmingFromInventory) ID() uint8 {
	return IDItemStackOperationHighLevelTrimming
}

func (TrimmingFromInventory) CanInline() bool {
	return false
}
