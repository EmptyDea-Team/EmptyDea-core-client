package game_interface

import (
	"context"
	"fmt"

	game_interface_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface"
	resources_uqholder "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control/uqholder"

	game_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	"google.golang.org/grpc"
)

// Player 扩展自 resources_control/uqholder.Player，添加玩家交互方法。
type Player struct {
	*resources_uqholder.Player
	client game_interface_pb.PlayerKitServiceClient
	ref    *resources_control_pb.PlayerRef
	kit    *PlayerKit
}

func (p *PlayerKit) newPlayer(ref *resources_control_pb.PlayerRef) Player {
	return Player{
		Player: resources_uqholder.NewPlayer(p.uqholderClient, ref),
		client: p.client,
		ref:    ref,
		kit:    p,
	}
}

// SendChat 向玩家发送普通聊天栏消息。
func (p Player) SendChat(ctx context.Context, msg string) error {
	_, err := p.client.SendChat(ctx, &game_interface_pb.PlayerChatRequest{Player: p.ref, Message: msg})
	return err
}

// SendRawChat 向玩家发送原始 rawtext 聊天消息。
func (p Player) SendRawChat(ctx context.Context, rawText string) error {
	_, err := p.client.SendRawChat(ctx, &game_interface_pb.PlayerRawChatRequest{Player: p.ref, RawText: rawText})
	return err
}

// SendTitle 向玩家发送标题文本。
func (p Player) SendTitle(ctx context.Context, title string) error {
	_, err := p.client.SendTitle(ctx, &game_interface_pb.PlayerTitleRequest{Player: p.ref, Title: title})
	return err
}

// SendRawTitle 向玩家发送原始 rawtext 标题文本。
func (p Player) SendRawTitle(ctx context.Context, rawTitle string) error {
	_, err := p.client.SendRawTitle(ctx, &game_interface_pb.PlayerRawTitleRequest{Player: p.ref, RawTitle: rawTitle})
	return err
}

// SendSubTitle 向玩家发送副标题文本。
func (p Player) SendSubTitle(ctx context.Context, subtitle string) error {
	_, err := p.client.SendSubTitle(ctx, &game_interface_pb.PlayerSubTitleRequest{Player: p.ref, Subtitle: subtitle})
	return err
}

// SendRawSubTitle 向玩家发送原始 rawtext 副标题文本。
func (p Player) SendRawSubTitle(ctx context.Context, rawSubtitle string) error {
	_, err := p.client.SendRawSubTitle(ctx, &game_interface_pb.PlayerRawSubTitleRequest{Player: p.ref, RawSubtitle: rawSubtitle})
	return err
}

// SendActionBar 向玩家发送 actionbar 文本。
func (p Player) SendActionBar(ctx context.Context, actionBar string) error {
	_, err := p.client.SendActionBar(ctx, &game_interface_pb.PlayerActionBarRequest{Player: p.ref, ActionBar: actionBar})
	return err
}

// SendRawActionBar 向玩家发送原始 rawtext actionbar 文本。
func (p Player) SendRawActionBar(ctx context.Context, rawText string) error {
	_, err := p.client.SendRawActionBar(ctx, &game_interface_pb.PlayerRawActionBarRequest{Player: p.ref, RawText: rawText})
	return err
}

// OpenAbility 创建玩家能力变更构建器。
func (p Player) OpenAbility(ctx context.Context) game_interface_api.AbilityBuilder {
	return &AbilityBuilder{player: p, ctx: ctx}
}

// AbilityBuilder 暂存玩家能力变更，并在 Commit 后提交到服务端。
type AbilityBuilder struct {
	player Player
	ctx    context.Context
	err    error
}

func (b *AbilityBuilder) setAbility(action func(context.Context, *game_interface_pb.PlayerAbilityRequest, ...grpc.CallOption) (*game_interface_pb.PlayerActionResponse, error), allow bool) game_interface_api.AbilityBuilder {
	if b.err != nil {
		return b
	}
	_, b.err = action(b.ctx, &game_interface_pb.PlayerAbilityRequest{Player: b.player.ref, Allow: allow})
	return b
}

// SetBuildAbility 暂存玩家是否可以放置方块。
func (b *AbilityBuilder) SetBuildAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetBuildAbility, allow)
}

// SetMineAbility 暂存玩家是否可以挖掘方块。
func (b *AbilityBuilder) SetMineAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetMineAbility, allow)
}

// SetDoorsAndSwitchesAbility 暂存玩家是否可以使用门和开关。
func (b *AbilityBuilder) SetDoorsAndSwitchesAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetDoorsAndSwitchesAbility, allow)
}

// SetOpenContainersAbility 暂存玩家是否可以打开容器。
func (b *AbilityBuilder) SetOpenContainersAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetOpenContainersAbility, allow)
}

// SetAttackPlayersAbility 暂存玩家是否可以攻击玩家。
func (b *AbilityBuilder) SetAttackPlayersAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetAttackPlayersAbility, allow)
}

// SetAttackMobsAbility 暂存玩家是否可以攻击生物。
func (b *AbilityBuilder) SetAttackMobsAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetAttackMobsAbility, allow)
}

// SetOperatorCommandsAbility 暂存玩家是否可以使用操作员命令。
func (b *AbilityBuilder) SetOperatorCommandsAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetOperatorCommandsAbility, allow)
}

// SetTeleportAbility 暂存玩家是否可以传送。
func (b *AbilityBuilder) SetTeleportAbility(allow bool) game_interface_api.AbilityBuilder {
	return b.setAbility(b.player.client.SetTeleportAbility, allow)
}

// Commit 提交玩家能力变更。
func (b *AbilityBuilder) Commit() error {
	if b.err != nil {
		return b.err
	}
	if b.ctx == nil {
		return fmt.Errorf("Commit: context is nil")
	}
	_, err := b.player.client.CommitAbility(b.ctx, &game_interface_pb.PlayerCommitAbilityRequest{Player: b.player.ref})
	return err
}

// PlayerKit 只负责获取 Player 实例。
type PlayerKit struct {
	client         game_interface_pb.PlayerKitServiceClient
	uqholderClient resources_control_pb.UQHolderClient
}

// ListOnlinePlayers 返回当前在线玩家列表。
func (p *PlayerKit) ListOnlinePlayers(ctx context.Context) ([]game_interface_api.Player, error) {
	resp, err := p.client.ListOnlinePlayers(ctx, &game_control_pb.Empty{})
	if err != nil {
		return nil, err
	}
	players := make([]game_interface_api.Player, 0, len(resp.Players))
	for _, ref := range resp.Players {
		if ref != nil {
			players = append(players, p.newPlayer(ref))
		}
	}
	return players, nil
}

// GetPlayerByName 根据玩家名查询玩家。
func (p *PlayerKit) GetPlayerByName(ctx context.Context, name string) (game_interface_api.Player, bool, error) {
	resp, err := p.client.GetPlayerByName(ctx, &game_interface_pb.GetPlayerByNameRequest{Name: name})
	return p.playerLookup(resp, err)
}

// GetPlayerByUUIDString 根据玩家 UUID 字符串查询玩家。
func (p *PlayerKit) GetPlayerByUUIDString(ctx context.Context, uuidString string) (game_interface_api.Player, bool, error) {
	resp, err := p.client.GetPlayerByUUIDString(ctx, &game_interface_pb.GetPlayerByUUIDRequest{UUID: uuidString})
	return p.playerLookup(resp, err)
}

// GetPlayerByUniqueID 根据实体唯一 ID 查询玩家。
func (p *PlayerKit) GetPlayerByUniqueID(ctx context.Context, id int64) (game_interface_api.Player, bool, error) {
	resp, err := p.client.GetPlayerByUniqueID(ctx, &game_interface_pb.GetPlayerByUniqueIDRequest{UniqueID: id})
	return p.playerLookup(resp, err)
}

// GetPlayerByRuntimeID 根据实体运行时 ID 查询玩家。
func (p *PlayerKit) GetPlayerByRuntimeID(ctx context.Context, id uint64) (game_interface_api.Player, bool, error) {
	resp, err := p.client.GetPlayerByRuntimeID(ctx, &game_interface_pb.GetPlayerByRuntimeIDRequest{RuntimeID: id})
	return p.playerLookup(resp, err)
}

func (p *PlayerKit) playerLookup(resp *game_interface_pb.PlayerLookupResponse, err error) (game_interface_api.Player, bool, error) {
	if err != nil {
		return Player{}, false, err
	}
	if resp == nil || !resp.Found || resp.Player == nil {
		return Player{}, false, nil
	}
	return p.newPlayer(resp.Player), true, nil
}
