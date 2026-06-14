package game_interface

import (
	"context"

	resources_uqholder "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control/uqholder"

	game_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
)

// Player 扩展自 resources_control/uqholder.Player，添加玩家交互方法。
type Player struct {
	*resources_uqholder.Player
	client game_interface_pb.PlayerKitServiceClient
	ref    *resources_control_pb.PlayerRef
}

func newPlayer(client game_interface_pb.PlayerKitServiceClient, uqholderClient resources_control_pb.UQHolderClient, ref *resources_control_pb.PlayerRef) Player {
	return Player{
		Player: resources_uqholder.NewPlayer(uqholderClient, ref),
		client: client,
		ref:    ref,
	}
}

func (p Player) SendChat(ctx context.Context, msg string) error {
	_, err := p.client.SendChat(ctx, &game_interface_pb.PlayerChatRequest{Player: p.ref, Message: msg})
	return err
}

func (p Player) SendRawChat(ctx context.Context, rawText string) error {
	_, err := p.client.SendRawChat(ctx, &game_interface_pb.PlayerRawChatRequest{Player: p.ref, RawText: rawText})
	return err
}

func (p Player) SendTitle(ctx context.Context, title string) error {
	_, err := p.client.SendTitle(ctx, &game_interface_pb.PlayerTitleRequest{Player: p.ref, Title: title})
	return err
}

func (p Player) SendRawTitle(ctx context.Context, rawTitle string) error {
	_, err := p.client.SendRawTitle(ctx, &game_interface_pb.PlayerRawTitleRequest{Player: p.ref, RawTitle: rawTitle})
	return err
}

func (p Player) SendSubTitle(ctx context.Context, subtitle string) error {
	_, err := p.client.SendSubTitle(ctx, &game_interface_pb.PlayerSubTitleRequest{Player: p.ref, Subtitle: subtitle})
	return err
}

func (p Player) SendRawSubTitle(ctx context.Context, rawSubtitle string) error {
	_, err := p.client.SendRawSubTitle(ctx, &game_interface_pb.PlayerRawSubTitleRequest{Player: p.ref, RawSubtitle: rawSubtitle})
	return err
}

func (p Player) SendActionBar(ctx context.Context, actionBar string) error {
	_, err := p.client.SendActionBar(ctx, &game_interface_pb.PlayerActionBarRequest{Player: p.ref, ActionBar: actionBar})
	return err
}

func (p Player) SendRawActionBar(ctx context.Context, rawText string) error {
	_, err := p.client.SendRawActionBar(ctx, &game_interface_pb.PlayerRawActionBarRequest{Player: p.ref, RawText: rawText})
	return err
}

// PlayerKit 只负责获取 Player 实例。
type PlayerKit struct {
	client         game_interface_pb.PlayerKitServiceClient
	uqholderClient resources_control_pb.UQHolderClient
}

func (p *PlayerKit) ListOnlinePlayers(ctx context.Context) ([]Player, error) {
	resp, err := p.client.ListOnlinePlayers(ctx, &game_control_pb.Empty{})
	if err != nil {
		return nil, err
	}
	players := make([]Player, 0, len(resp.Players))
	for _, ref := range resp.Players {
		if ref != nil {
			players = append(players, newPlayer(p.client, p.uqholderClient, ref))
		}
	}
	return players, nil
}

func (p *PlayerKit) GetPlayerByName(ctx context.Context, name string) (Player, bool, error) {
	resp, err := p.client.GetPlayerByName(ctx, &game_interface_pb.GetPlayerByNameRequest{Name: name})
	return p.playerLookup(resp, err)
}

func (p *PlayerKit) GetPlayerByUUIDString(ctx context.Context, uuidString string) (Player, bool, error) {
	resp, err := p.client.GetPlayerByUUIDString(ctx, &game_interface_pb.GetPlayerByUUIDRequest{UUID: uuidString})
	return p.playerLookup(resp, err)
}

func (p *PlayerKit) GetPlayerByUniqueID(ctx context.Context, id int64) (Player, bool, error) {
	resp, err := p.client.GetPlayerByUniqueID(ctx, &game_interface_pb.GetPlayerByUniqueIDRequest{UniqueID: id})
	return p.playerLookup(resp, err)
}

func (p *PlayerKit) GetPlayerByRuntimeID(ctx context.Context, id uint64) (Player, bool, error) {
	resp, err := p.client.GetPlayerByRuntimeID(ctx, &game_interface_pb.GetPlayerByRuntimeIDRequest{RuntimeID: id})
	return p.playerLookup(resp, err)
}

func (p *PlayerKit) playerLookup(resp *game_interface_pb.PlayerLookupResponse, err error) (Player, bool, error) {
	if err != nil {
		return Player{}, false, err
	}
	if resp == nil || !resp.Found || resp.Player == nil {
		return Player{}, false, nil
	}
	return newPlayer(p.client, p.uqholderClient, resp.Player), true, nil
}
