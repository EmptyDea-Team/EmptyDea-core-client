package uqholder

import (
	"context"

	uqholder_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control/uqholder"
	client_convertutil "github.com/EmptyDea-Team/EmptyDea-core-client/convertutil"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
)

// Players 暴露在线玩家索引的只读查询接口。
type Players struct {
	client resources_control_pb.UQHolderClient
}

// GetOnline 返回当前在线玩家列表。
func (p *Players) GetOnline(ctx context.Context) (players []uqholder_api.Player, err error) {
	resp, err := p.client.GetOnlinePlayers(ctx, client_convertutil.Empty())
	if err != nil {
		return nil, err
	}
	players = make([]uqholder_api.Player, 0, len(resp.Players))
	for _, ref := range resp.Players {
		if ref != nil {
			players = append(players, NewPlayer(p.client, ref))
		}
	}
	return players, nil
}

// GetByUUIDString 根据玩家 UUID 字符串查询玩家。
func (p *Players) GetByUUIDString(ctx context.Context, id string) (player uqholder_api.Player, existed bool, err error) {
	resp, err := p.client.GetPlayerByUUID(ctx, &resources_control_pb.GetPlayerByUUIDRequest{UUID: id})
	return p.playerLookup(resp, err)
}

// GetByName 根据玩家名称查询玩家。
func (p *Players) GetByName(ctx context.Context, name string) (player uqholder_api.Player, existed bool, err error) {
	resp, err := p.client.GetPlayerByName(ctx, &resources_control_pb.GetPlayerByNameRequest{Name: name})
	return p.playerLookup(resp, err)
}

// GetByUniqueID 根据玩家实体唯一 ID 查询玩家。
func (p *Players) GetByUniqueID(ctx context.Context, id int64) (player uqholder_api.Player, existed bool, err error) {
	resp, err := p.client.GetPlayerByUniqueID(ctx, &resources_control_pb.GetPlayerByUniqueIDRequest{UniqueID: id})
	return p.playerLookup(resp, err)
}

// GetByRuntimeID 根据玩家实体运行时 ID 查询玩家。
func (p *Players) GetByRuntimeID(ctx context.Context, id uint64) (player uqholder_api.Player, existed bool, err error) {
	resp, err := p.client.GetPlayerByRuntimeID(ctx, &resources_control_pb.GetPlayerByRuntimeIDRequest{RuntimeID: id})
	return p.playerLookup(resp, err)
}

func (p *Players) playerLookup(resp *resources_control_pb.PlayerLookupResponse, err error) (player uqholder_api.Player, existed bool, retErr error) {
	if err != nil {
		return nil, false, err
	}
	if resp == nil || !resp.Found || resp.Player == nil {
		return nil, false, nil
	}
	return NewPlayer(p.client, resp.Player), true, nil
}
