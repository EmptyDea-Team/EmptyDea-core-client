package game_interface

import (
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	"google.golang.org/grpc"
)

// GameInterface 实现远程机器人与租赁服的高级交互。
type GameInterface struct {
	commands              *Commands
	structureBackup       *StructureBackup
	querytarget           *Querytarget
	movement              *Movement
	setblock              *SetBlock
	replaceitem           *Replaceitem
	botClick              *BotClick
	itemStackOperation    *ItemStackOperation
	containerOpenAndClose *ContainerOpenAndClose
	itemCopy              *ItemCopy
	itemTransition        *ItemTransition
	playerKit             *PlayerKit
}

// New 创建游戏交互层客户端集合。
func New(conn grpc.ClientConnInterface) *GameInterface {
	itemStackOperation := &ItemStackOperation{client: game_interface_pb.NewItemStackTransactionServiceClient(conn)}
	return &GameInterface{
		commands:              &Commands{client: game_interface_pb.NewCommandsServiceClient(conn)},
		structureBackup:       &StructureBackup{client: game_interface_pb.NewStructureBackupServiceClient(conn)},
		querytarget:           &Querytarget{client: game_interface_pb.NewQuerytargetServiceClient(conn)},
		movement:              &Movement{client: game_interface_pb.NewMovementServiceClient(conn)},
		setblock:              &SetBlock{client: game_interface_pb.NewSetBlockServiceClient(conn)},
		replaceitem:           &Replaceitem{client: game_interface_pb.NewReplaceitemServiceClient(conn)},
		botClick:              &BotClick{client: game_interface_pb.NewBotClickServiceClient(conn)},
		itemStackOperation:    itemStackOperation,
		containerOpenAndClose: &ContainerOpenAndClose{client: game_interface_pb.NewContainerOpenAndCloseServiceClient(conn)},
		itemCopy:              &ItemCopy{client: game_interface_pb.NewItemCopyServiceClient(conn)},
		itemTransition:        &ItemTransition{client: game_interface_pb.NewItemTransitionServiceClient(conn)},
		playerKit: &PlayerKit{
			client:         game_interface_pb.NewPlayerKitServiceClient(conn),
			uqholderClient: resources_control_pb.NewUQHolderClient(conn),
		},
	}
}

// Commands 返回机器人在 MC 命令收发上的相关实现。
func (g *GameInterface) Commands() *Commands {
	return g.commands
}

// StructureBackup 返回机器人在结构备份和恢复上的相关实现。
func (g *GameInterface) StructureBackup() *StructureBackup {
	return g.structureBackup
}

// Querytarget 返回机器人在 querytarget 命令上的实现。
func (g *GameInterface) Querytarget() *Querytarget {
	return g.querytarget
}

// Movement 返回机器人在移动上的相关实现。
func (g *GameInterface) Movement() *Movement {
	return g.movement
}

// SetBlock 返回机器人在方块放置上的相关实现。
func (g *GameInterface) SetBlock() *SetBlock {
	return g.setblock
}

// Replaceitem 返回机器人在 Replaceitem 命令上的简单包装。
func (g *GameInterface) Replaceitem() *Replaceitem {
	return g.replaceitem
}

// BotClick 返回机器人在点击操作上的相关实现。
func (g *GameInterface) BotClick() *BotClick {
	return g.botClick
}

// ItemStackOperation 返回机器人在物品堆栈操作请求上的相关实现。
func (g *GameInterface) ItemStackOperation() *ItemStackOperation {
	return g.itemStackOperation
}

// ContainerOpenAndClose 返回机器人在容器打开和关闭上的相关实现。
func (g *GameInterface) ContainerOpenAndClose() *ContainerOpenAndClose {
	return g.containerOpenAndClose
}

// ItemCopy 返回机器人在物品复制上的相关实现。
func (g *GameInterface) ItemCopy() *ItemCopy {
	return g.itemCopy
}

// ItemTransition 返回机器人在物品状态转移上的实现。
func (g *GameInterface) ItemTransition() *ItemTransition {
	return g.itemTransition
}

// PlayerKit 返回机器人在在线玩家查询和交互上的实现。
func (g *GameInterface) PlayerKit() *PlayerKit {
	return g.playerKit
}
