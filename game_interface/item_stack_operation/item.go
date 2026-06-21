package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-api/define"

type (
	// ExpectedNewItem 描述物品堆栈操作完成后的期望物品信息。
	ExpectedNewItem = define.ExpectedNewItem
	// ItemNewType 描述物品基本类型字段应如何更新。
	ItemNewType = define.ItemNewType
	// ItemNewBlockRuntimeID 描述物品方块运行时 ID 应如何更新。
	ItemNewBlockRuntimeID = define.ItemNewBlockRuntimeID
	// ItemNewNBTData 描述物品 NBT 数据应如何更新。
	ItemNewNBTData = define.ItemNewNBTData
	// ItemNewComponent 描述物品 Legacy 组件应如何更新。
	ItemNewComponent = define.ItemNewComponent
)
