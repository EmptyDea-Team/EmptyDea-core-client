package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Trimming 指示锻造台纹饰操作。
type Trimming struct {
	// TrimItem 是待纹饰物品所在位置。
	TrimItem resources_control.SlotLocation
	// Material 是纹饰材料所在位置。
	Material resources_control.SlotLocation
	// Template 是纹饰模板所在位置。
	Template resources_control.SlotLocation
	// ResultItem 是结果物品期望更新后的数据。
	ResultItem ExpectedNewItem
}

// ID 返回锻造台纹饰操作编号。
func (Trimming) ID() uint8 {
	return IDItemStackOperationHighLevelTrimming
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (Trimming) CanInline() bool {
	return false
}

// TrimmingFromInventory 指示使用背包内物品执行锻造台纹饰的操作。
type TrimmingFromInventory struct {
	// TrimItemSlot 是背包内待纹饰物品槽位。
	TrimItemSlot resources_control.SlotID
	// MaterialSlot 是背包内纹饰材料槽位。
	MaterialSlot resources_control.SlotID
	// TemplateSlot 是背包内纹饰模板槽位。
	TemplateSlot resources_control.SlotID
	// ResultItem 是结果物品期望更新后的数据。
	ResultItem ExpectedNewItem
}

// ID 返回锻造台纹饰操作编号。
func (TrimmingFromInventory) ID() uint8 {
	return IDItemStackOperationHighLevelTrimming
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (TrimmingFromInventory) CanInline() bool {
	return false
}
