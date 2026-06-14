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

func (b *Bot) GetName(ctx context.Context) (name string, ok bool, err error) {
	return client_convertutil.StringValue(b.client.GetBotName(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetXUID(ctx context.Context) (xuid string, ok bool, err error) {
	return client_convertutil.StringValue(b.client.GetBotXUID(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetUUID(ctx context.Context) (uuid string, ok bool, err error) {
	return client_convertutil.StringValue(b.client.GetBotUUID(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetEntityUniqueID(ctx context.Context) (entityUniqueID int64, ok bool, err error) {
	return client_convertutil.Int64Value(b.client.GetBotEntityUniqueID(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetEntityRuntimeID(ctx context.Context) (entityRuntimeID uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(b.client.GetBotEntityRuntimeID(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetPosition(ctx context.Context) (pos *mgl32_pb.Vec3, ok bool, err error) {
	resp, err := b.client.GetBotPosition(ctx, client_convertutil.Empty())
	return client_convertutil.Vec3Value(resp, err)
}

func (b *Bot) GetDimension(ctx context.Context) (dimension int32, ok bool, err error) {
	return client_convertutil.Int32Value(b.client.GetBotDimension(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetGameMode(ctx context.Context) (gameMode int32, ok bool, err error) {
	return client_convertutil.Int32Value(b.client.GetBotGameMode(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetHealth(ctx context.Context) (health float32, ok bool, err error) {
	return client_convertutil.Float32Value(b.client.GetBotHealth(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetHunger(ctx context.Context) (hunger float32, ok bool, err error) {
	return client_convertutil.Float32Value(b.client.GetBotHunger(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetSaturation(ctx context.Context) (saturation float32, ok bool, err error) {
	return client_convertutil.Float32Value(b.client.GetBotSaturation(ctx, client_convertutil.Empty()))
}

func (b *Bot) GetHotBarSlot(ctx context.Context) (hotBarSlot byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(b.client.GetBotHotBarSlot(ctx, client_convertutil.Empty()))
	return byte(value), ok, err
}
