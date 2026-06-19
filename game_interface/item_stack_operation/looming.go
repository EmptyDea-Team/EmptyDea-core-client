package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Looming 指示织布机操作。
type Looming struct {
	// PatternName 是织布机图案名称。
	PatternName string
	// PatternPath 是图案物品所在位置。
	PatternPath resources_control.SlotLocation
	// BannerPath 是旗帜物品所在位置。
	BannerPath resources_control.SlotLocation
	// DyePath 是染料物品所在位置。
	DyePath resources_control.SlotLocation
	// ResultItem 是结果物品期望更新后的数据。
	ResultItem ExpectedNewItem
}

// ID 返回织布机操作编号。
func (Looming) ID() uint8 {
	return IDItemStackOperationHighLevelLooming
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (Looming) CanInline() bool {
	return false
}

// LoomingFromInventory 指示使用背包内物品执行织布机合成的操作。
type LoomingFromInventory struct {
	// PatternName 是织布机图案名称。
	PatternName string
	// PatternSlot 是背包内图案物品槽位。
	PatternSlot resources_control.SlotID
	// BannerSlot 是背包内旗帜物品槽位。
	BannerSlot resources_control.SlotID
	// DyeSlot 是背包内染料物品槽位。
	DyeSlot resources_control.SlotID
	// ResultItem 是结果物品期望更新后的数据。
	ResultItem ExpectedNewItem
}

// ID 返回织布机操作编号。
func (LoomingFromInventory) ID() uint8 {
	return IDItemStackOperationHighLevelLooming
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (LoomingFromInventory) CanInline() bool {
	return false
}
