package resources_control

import protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"

// SlotID 是单个物品栏槽位的索引，它是从 0 开始索引的。
type SlotID uint8

// WindowID 是机器人已打开(或持有)的库存的窗口 ID。
type WindowID uint32

// DynamicContainerID 是机器人已打开(或持有)的动态库存的容器 ID。
type DynamicContainerID uint32

// ContainerID 是容器 ID。
type ContainerID uint8

// WindowName 唯一标识一个普通窗口或动态容器窗口。
type WindowName struct {
	WindowID           WindowID
	DynamicContainerID DynamicContainerID
}

// SlotLocation 描述一个物品的所在的位置。
type SlotLocation struct {
	WindowName
	SlotID SlotID
}

const (
	// WindowIDDynamic 是动态容器窗口的本地窗口 ID。
	WindowIDDynamic WindowID = 125
	// WindowIDCrafting 是合成栏窗口 ID。
	WindowIDCrafting WindowID = 121
)

var (
	// WindowNameInventory 是玩家背包窗口名。
	WindowNameInventory = WindowName{WindowID: WindowID(protocol_pb.WindowIDEnum_WindowIDInventory)}
	// WindowNameOffHand 是副手窗口名。
	WindowNameOffHand = WindowName{WindowID: WindowID(protocol_pb.WindowIDEnum_WindowIDOffHand)}
	// WindowNameArmour 是盔甲栏窗口名。
	WindowNameArmour = WindowName{WindowID: WindowID(protocol_pb.WindowIDEnum_WindowIDArmour)}
	// WindowNameCrafting 是合成栏窗口名。
	WindowNameCrafting = WindowName{WindowID: WindowIDCrafting}
	// WindowNameUI 是 UI 窗口名。
	WindowNameUI = WindowName{WindowID: WindowID(protocol_pb.WindowIDEnum_WindowIDUI)}
)

// NewWindowName 基于窗口 ID 和动态容器 ID 构造窗口名。
func NewWindowName(windowID WindowID, dynamicContainerID DynamicContainerID) WindowName {
	if windowID != WindowIDDynamic {
		dynamicContainerID = 0
	}
	return WindowName{WindowID: windowID, DynamicContainerID: dynamicContainerID}
}

// BotInfo 是当前机器人基础登录信息。
type BotInfo struct {
	BotName         string
	XUID            string
	EntityUniqueID  int64
	EntityRuntimeID uint64
}
