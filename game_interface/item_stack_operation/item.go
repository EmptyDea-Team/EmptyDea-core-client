package item_stack_operation

// ExpectedNewItem 描述单个物品堆栈在经历一次物品堆栈操作后，
// 其最终应当拥有的一些数据信息。
type ExpectedNewItem struct {
	ItemType       ItemNewType
	BlockRuntimeID ItemNewBlockRuntimeID
	NBT            ItemNewNBTData
	Component      ItemNewComponent
}

// ItemNewType 描述物品的一些基本字段应如何更新。
type ItemNewType struct {
	UseNetworkID bool
	NetworkID    int32
	UseMetadata  bool
	Metadata     uint32
}

// ItemNewBlockRuntimeID 描述物品对应的方块运行时数据应该如何更新。
type ItemNewBlockRuntimeID struct {
	UseBlockRuntimeID bool
	BlockRuntimeID    int32
}

// ItemNewNBTData 描述物品的新 NBT 字段如何更新。
type ItemNewNBTData struct {
	UseNBTData       bool
	UseOriginDamage  bool
	NBTData          map[string]any
	ChangeRepairCost bool
	RepairCostDelta  int32
	ChangeDamage     bool
	DamageDelta      int32
}

// ItemNewComponent 描述物品的 Legacy 物品组件应当如何更新。
type ItemNewComponent struct {
	UseCanPlaceOn bool
	CanPlaceOn    []string
	UseCanDestroy bool
	CanDestroy    []string
}
