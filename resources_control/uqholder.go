package resources_control

import uqholder_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control/uqholder"

type (
	// UQHolder 聚合由数据包维护的机器人、玩家、世界和实体状态。
	UQHolder = uqholder_client.UQHolder
	// Bot 暴露机器人自身状态的只读查询接口。
	Bot = uqholder_client.Bot
	// Players 暴露在线玩家索引的只读查询接口。
	Players = uqholder_client.Players
	// Player 暴露单个玩家状态的只读视图。
	Player = uqholder_client.Player
	// PlayerAbilities 暴露玩家权限和能力层状态。
	PlayerAbilities = uqholder_client.PlayerAbilities
	// World 暴露世界状态的只读查询接口。
	World = uqholder_client.World
	// GameRule 暴露单条游戏规则的值和可修改性。
	GameRule = uqholder_client.GameRule
	// Entities 暴露实体状态索引。
	Entities = uqholder_client.Entities
	// Entity 暴露按运行时 ID 维护的实体状态。
	Entity = uqholder_client.Entity
	// MobEffect 暴露实体身上的单个药水效果状态。
	MobEffect = uqholder_client.MobEffect
)
