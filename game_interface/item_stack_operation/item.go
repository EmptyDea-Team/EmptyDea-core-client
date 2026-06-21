package item_stack_operation

import item_stack_operation_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_operation"

type (
	// ExpectedNewItem 描述物品堆栈操作完成后的期望物品信息。
	ExpectedNewItem = item_stack_operation_api.ExpectedNewItem
	// ItemNewType 描述物品基本类型字段应如何更新。
	ItemNewType = item_stack_operation_api.ItemNewType
	// ItemNewBlockRuntimeID 描述物品方块运行时 ID 应如何更新。
	ItemNewBlockRuntimeID = item_stack_operation_api.ItemNewBlockRuntimeID
	// ItemNewNBTData 描述物品 NBT 数据应如何更新。
	ItemNewNBTData = item_stack_operation_api.ItemNewNBTData
	// ItemNewComponent 描述物品 Legacy 组件应如何更新。
	ItemNewComponent = item_stack_operation_api.ItemNewComponent
)
