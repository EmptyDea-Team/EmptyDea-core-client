package uqholder

import (
	"context"

	client_convertutil "github.com/EmptyDea-Team/EmptyDea-core-client/convertutil"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
)

// World 暴露世界状态的只读查询接口。
type World struct {
	client resources_control_pb.UQHolderClient
}

func (w *World) GetCurrentTick(ctx context.Context) (currentTick int64, ok bool, err error) {
	return client_convertutil.Int64Value(w.client.GetWorldCurrentTick(ctx, client_convertutil.Empty()))
}

func (w *World) GetSyncRatio(ctx context.Context) (syncRatio float32, ok bool, err error) {
	return client_convertutil.Float32Value(w.client.GetWorldSyncRatio(ctx, client_convertutil.Empty()))
}

func (w *World) GetTime(ctx context.Context) (time int32, ok bool, err error) {
	return client_convertutil.Int32Value(w.client.GetWorldTime(ctx, client_convertutil.Empty()))
}

func (w *World) GetDayTime(ctx context.Context) (dayTime int32, ok bool, err error) {
	return client_convertutil.Int32Value(w.client.GetWorldDayTime(ctx, client_convertutil.Empty()))
}

func (w *World) GetDayTimePercent(ctx context.Context) (dayTimePercent float32, ok bool, err error) {
	return client_convertutil.Float32Value(w.client.GetWorldDayTimePercent(ctx, client_convertutil.Empty()))
}

func (w *World) GetDifficulty(ctx context.Context) (difficulty uint32, ok bool, err error) {
	return client_convertutil.Uint32Value(w.client.GetWorldDifficulty(ctx, client_convertutil.Empty()))
}

func (w *World) GetWorldGameMode(ctx context.Context) (worldGameMode int32, ok bool, err error) {
	return client_convertutil.Int32Value(w.client.GetWorldGameMode(ctx, client_convertutil.Empty()))
}

func (w *World) GetGameRuleNames(ctx context.Context) (names []string, err error) {
	resp, err := w.client.GetGameRuleNames(ctx, client_convertutil.Empty())
	if err != nil {
		return nil, err
	}
	return resp.Names, nil
}

func (w *World) GetGameRule(ctx context.Context, name string) (gameRule *GameRule, existed bool, err error) {
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

func (g GameRule) GetCanBeModifiedByPlayer() (canBeModifiedByPlayer bool) {
	return g.canBeModifiedByPlayer
}
func (g GameRule) GetValue() (value string) { return g.value }
