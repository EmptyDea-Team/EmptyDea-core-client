package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Looming 指示织布机操作。
type Looming struct {
	PatternName string
	PatternPath resources_control.SlotLocation
	BannerPath  resources_control.SlotLocation
	DyePath     resources_control.SlotLocation
	ResultItem  ExpectedNewItem
}

func (Looming) ID() uint8 {
	return IDItemStackOperationHighLevelLooming
}

func (Looming) CanInline() bool {
	return false
}

// LoomingFromInventory 指示使用背包内物品执行织布机合成的操作。
type LoomingFromInventory struct {
	PatternName string
	PatternSlot resources_control.SlotID
	BannerSlot  resources_control.SlotID
	DyeSlot     resources_control.SlotID
	ResultItem  ExpectedNewItem
}

func (LoomingFromInventory) ID() uint8 {
	return IDItemStackOperationHighLevelLooming
}

func (LoomingFromInventory) CanInline() bool {
	return false
}
