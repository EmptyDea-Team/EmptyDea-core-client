package uqholder

import resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"

// UQHolder 聚合由数据包维护的机器人、玩家、世界和实体状态。
type UQHolder struct {
	client resources_control_pb.UQHolderClient
}

// New 基于 UQHolderClient 创建远程 UQHolder。
func New(client resources_control_pb.UQHolderClient) (uqholder *UQHolder) {
	return &UQHolder{client: client}
}

// Bot 返回机器人自身状态的只读视图。
func (u *UQHolder) Bot() (bot *Bot) {
	return &Bot{client: u.client}
}

// Players 返回在线玩家状态索引。
func (u *UQHolder) Players() (players *Players) {
	return &Players{client: u.client}
}

// World 返回世界状态查询接口。
func (u *UQHolder) World() (world *World) {
	return &World{client: u.client}
}

// Entities 返回实体状态索引。
func (u *UQHolder) Entities() (entities *Entities) {
	return &Entities{client: u.client}
}
