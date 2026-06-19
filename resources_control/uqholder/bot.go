package uqholder

import (
	"context"

	client_convertutil "github.com/EmptyDea-Team/EmptyDea-core-client/convertutil"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
)

// Bot 暴露机器人自身状态的只读查询接口。
type Bot struct {
	client resources_control_pb.UQHolderClient
}

// GetName 返回机器人玩家名。
func (b *Bot) GetName(ctx context.Context) (name string, ok bool, err error) {
	return client_convertutil.StringValue(b.client.GetBotName(ctx, client_convertutil.Empty()))
}

// GetXUID 返回机器人 XUID。
func (b *Bot) GetXUID(ctx context.Context) (xuid string, ok bool, err error) {
	return client_convertutil.StringValue(b.client.GetBotXUID(ctx, client_convertutil.Empty()))
}

// GetUUID 返回机器人 UUID 字符串。
func (b *Bot) GetUUID(ctx context.Context) (uuid string, ok bool, err error) {
	return client_convertutil.StringValue(b.client.GetBotUUID(ctx, client_convertutil.Empty()))
}

// GetEntityUniqueID 返回机器人实体唯一 ID。
func (b *Bot) GetEntityUniqueID(ctx context.Context) (entityUniqueID int64, ok bool, err error) {
	return client_convertutil.Int64Value(b.client.GetBotEntityUniqueID(ctx, client_convertutil.Empty()))
}

// GetEntityRuntimeID 返回机器人实体运行时 ID。
func (b *Bot) GetEntityRuntimeID(ctx context.Context) (entityRuntimeID uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(b.client.GetBotEntityRuntimeID(ctx, client_convertutil.Empty()))
}

// GetPosition 返回机器人当前位置。
func (b *Bot) GetPosition(ctx context.Context) (pos *mgl32_pb.Vec3, ok bool, err error) {
	resp, err := b.client.GetBotPosition(ctx, client_convertutil.Empty())
	return client_convertutil.Vec3Value(resp, err)
}

// GetDimension 返回机器人所在维度。
func (b *Bot) GetDimension(ctx context.Context) (dimension int32, ok bool, err error) {
	return client_convertutil.Int32Value(b.client.GetBotDimension(ctx, client_convertutil.Empty()))
}

// GetGameMode 返回机器人当前游戏模式。
func (b *Bot) GetGameMode(ctx context.Context) (gameMode int32, ok bool, err error) {
	return client_convertutil.Int32Value(b.client.GetBotGameMode(ctx, client_convertutil.Empty()))
}

// GetHealth 返回机器人当前生命值。
func (b *Bot) GetHealth(ctx context.Context) (health float32, ok bool, err error) {
	return client_convertutil.Float32Value(b.client.GetBotHealth(ctx, client_convertutil.Empty()))
}

// GetHunger 返回机器人当前饥饿值。
func (b *Bot) GetHunger(ctx context.Context) (hunger float32, ok bool, err error) {
	return client_convertutil.Float32Value(b.client.GetBotHunger(ctx, client_convertutil.Empty()))
}

// GetSaturation 返回机器人当前饱和度。
func (b *Bot) GetSaturation(ctx context.Context) (saturation float32, ok bool, err error) {
	return client_convertutil.Float32Value(b.client.GetBotSaturation(ctx, client_convertutil.Empty()))
}

// GetHotBarSlot 返回机器人当前选择的快捷栏槽位。
func (b *Bot) GetHotBarSlot(ctx context.Context) (hotBarSlot byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(b.client.GetBotHotBarSlot(ctx, client_convertutil.Empty()))
	return byte(value), ok, err
}
