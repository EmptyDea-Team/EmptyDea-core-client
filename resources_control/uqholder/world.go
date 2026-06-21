package uqholder

import (
	"context"

	uqholder_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control/uqholder"
	client_convertutil "github.com/EmptyDea-Team/EmptyDea-core-client/convertutil"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
)

// World 暴露世界状态的只读查询接口。
type World struct {
	client resources_control_pb.UQHolderClient
}

// GetCurrentTick 返回世界当前 tick。
func (w *World) GetCurrentTick(ctx context.Context) (currentTick int64, ok bool, err error) {
	return client_convertutil.Int64Value(w.client.GetWorldCurrentTick(ctx, client_convertutil.Empty()))
}

// GetSyncRatio 返回世界同步比例。
func (w *World) GetSyncRatio(ctx context.Context) (syncRatio float32, ok bool, err error) {
	return client_convertutil.Float32Value(w.client.GetWorldSyncRatio(ctx, client_convertutil.Empty()))
}

// GetTime 返回世界时间。
func (w *World) GetTime(ctx context.Context) (time int32, ok bool, err error) {
	return client_convertutil.Int32Value(w.client.GetWorldTime(ctx, client_convertutil.Empty()))
}

// GetDayTime 返回世界日间时间。
func (w *World) GetDayTime(ctx context.Context) (dayTime int32, ok bool, err error) {
	return client_convertutil.Int32Value(w.client.GetWorldDayTime(ctx, client_convertutil.Empty()))
}

// GetDayTimePercent 返回世界日间时间百分比。
func (w *World) GetDayTimePercent(ctx context.Context) (dayTimePercent float32, ok bool, err error) {
	return client_convertutil.Float32Value(w.client.GetWorldDayTimePercent(ctx, client_convertutil.Empty()))
}

// GetDifficulty 返回世界难度。
func (w *World) GetDifficulty(ctx context.Context) (difficulty uint32, ok bool, err error) {
	return client_convertutil.Uint32Value(w.client.GetWorldDifficulty(ctx, client_convertutil.Empty()))
}

// GetWorldGameMode 返回世界默认游戏模式。
func (w *World) GetWorldGameMode(ctx context.Context) (worldGameMode int32, ok bool, err error) {
	return client_convertutil.Int32Value(w.client.GetWorldGameMode(ctx, client_convertutil.Empty()))
}

// GetGameRuleNames 返回当前全部游戏规则名称。
func (w *World) GetGameRuleNames(ctx context.Context) (names []string, err error) {
	resp, err := w.client.GetGameRuleNames(ctx, client_convertutil.Empty())
	if err != nil {
		return nil, err
	}
	return resp.Names, nil
}

// GetGameRule 根据名称返回单条游戏规则。
func (w *World) GetGameRule(ctx context.Context, name string) (gameRule uqholder_api.GameRule, existed bool, err error) {
	value, ok, err := client_convertutil.StringValue(w.client.GetGameRuleValue(ctx, &resources_control_pb.GameRuleKey{Name: name}))
	if err != nil || !ok {
		return nil, false, err
	}
	canModify, _, err := client_convertutil.BoolValue(w.client.GetGameRuleCanBeModifiedByPlayer(ctx, &resources_control_pb.GameRuleKey{Name: name}))
	if err != nil {
		return nil, false, err
	}
	return &GameRule{canBeModifiedByPlayer: canModify, value: value}, true, nil
}

// GameRule 暴露单条游戏规则的值和可修改性。
type GameRule struct {
	canBeModifiedByPlayer bool
	value                 string
}

// GetCanBeModifiedByPlayer 返回该游戏规则是否可由玩家修改。
func (g GameRule) GetCanBeModifiedByPlayer() (canBeModifiedByPlayer bool) {
	return g.canBeModifiedByPlayer
}

// GetValue 返回该游戏规则的字符串值。
func (g GameRule) GetValue() (value string) { return g.value }
