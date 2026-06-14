package resources_control

// SlotID 是单个物品栏槽位的索引，它是从 0 开始索引的。
type SlotID uint8

// WindowID 是机器人已打开(或持有)的库存的窗口 ID。
type WindowID uint32

// ContainerID 是容器 ID。
type ContainerID uint8

// SlotLocation 描述一个物品的所在的位置。
type SlotLocation struct {
	WindowID WindowID
	SlotID   SlotID
}

// BotInfo 是当前机器人基础登录信息。
type BotInfo struct {
	BotName         string
	XUID            string
	EntityUniqueID  int64
	EntityRuntimeID uint64
}
