package uqholder

import (
	"context"
	"time"

	client_convertutil "github.com/EmptyDea-Team/EmptyDea-core-client/convertutil"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
)

// Player 暴露单个玩家状态的只读视图。
type Player struct {
	client resources_control_pb.UQHolderClient
	ref    *resources_control_pb.PlayerRef
}

// NewPlayer 基于远程 UQHolder client 和玩家引用创建玩家只读视图。
func NewPlayer(client resources_control_pb.UQHolderClient, ref *resources_control_pb.PlayerRef) *Player {
	return &Player{client: client, ref: ref}
}

// GetUUIDString 返回玩家 UUID 字符串。
func (p *Player) GetUUIDString(ctx context.Context) (uuid string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerUUID(ctx, p.ref))
}

// GetName 返回玩家名称。
func (p *Player) GetName(ctx context.Context) (name string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerName(ctx, p.ref))
}

// GetXUID 返回玩家 XUID。
func (p *Player) GetXUID(ctx context.Context) (xuid string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerXUID(ctx, p.ref))
}

// GetEntityUniqueID 返回玩家实体唯一 ID。
func (p *Player) GetEntityUniqueID(ctx context.Context) (entityUniqueID int64, ok bool, err error) {
	return client_convertutil.Int64Value(p.client.GetPlayerEntityUniqueID(ctx, p.ref))
}

// GetEntityRuntimeID 返回玩家实体运行时 ID。
func (p *Player) GetEntityRuntimeID(ctx context.Context) (entityRuntimeID uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(p.client.GetPlayerEntityRuntimeID(ctx, p.ref))
}

// GetNeteaseUID 返回玩家网易 UID。
func (p *Player) GetNeteaseUID(ctx context.Context) (neteaseUID int64, ok bool, err error) {
	return client_convertutil.Int64Value(p.client.GetPlayerNeteaseUID(ctx, p.ref))
}

// GetLoginTime 返回玩家登录时间。
func (p *Player) GetLoginTime(ctx context.Context) (loginTime time.Time, ok bool, err error) {
	resp, err := p.client.GetPlayerLoginTime(ctx, p.ref)
	if err != nil {
		return time.Time{}, false, err
	}
	if !resp.Ok || resp.Value == nil {
		return time.Time{}, false, nil
	}
	return resp.Value.AsTime(), true, nil
}

// GetPlatformChatID 返回玩家平台聊天 ID。
func (p *Player) GetPlatformChatID(ctx context.Context) (platformChatID string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerPlatformChatID(ctx, p.ref))
}

// GetBuildPlatform 返回玩家构建平台。
func (p *Player) GetBuildPlatform(ctx context.Context) (buildPlatform int32, ok bool, err error) {
	return client_convertutil.Int32Value(p.client.GetPlayerBuildPlatform(ctx, p.ref))
}

// GetSkinID 返回玩家皮肤 ID。
func (p *Player) GetSkinID(ctx context.Context) (skinID string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerSkinID(ctx, p.ref))
}

// GetDeviceID 返回玩家设备 ID。
func (p *Player) GetDeviceID(ctx context.Context) (deviceID string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerDeviceID(ctx, p.ref))
}

// GetEntityMetadata 返回玩家实体元数据。
func (p *Player) GetEntityMetadata(ctx context.Context) (metadata *protocol_pb.EntityMetadata, ok bool, err error) {
	resp, err := p.client.GetPlayerEntityMetadata(ctx, p.ref)
	if err != nil {
		return nil, false, err
	}
	return resp.Value, resp.Ok, nil
}

// GetPosition 返回玩家位置。
func (p *Player) GetPosition(ctx context.Context) (pos *mgl32_pb.Vec3, ok bool, err error) {
	return client_convertutil.Vec3Value(p.client.GetPlayerPosition(ctx, p.ref))
}

// GetPitch 返回玩家俯仰角。
func (p *Player) GetPitch(ctx context.Context) (pitch float32, ok bool, err error) {
	return client_convertutil.Float32Value(p.client.GetPlayerPitch(ctx, p.ref))
}

// GetYaw 返回玩家偏航角。
func (p *Player) GetYaw(ctx context.Context) (yaw float32, ok bool, err error) {
	return client_convertutil.Float32Value(p.client.GetPlayerYaw(ctx, p.ref))
}

// GetHeadYaw 返回玩家头部偏航角。
func (p *Player) GetHeadYaw(ctx context.Context) (headYaw float32, ok bool, err error) {
	return client_convertutil.Float32Value(p.client.GetPlayerHeadYaw(ctx, p.ref))
}

// GetMoveMode 返回玩家移动模式。
func (p *Player) GetMoveMode(ctx context.Context) (moveMode byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(p.client.GetPlayerMoveMode(ctx, p.ref))
	return byte(value), ok, err
}

// GetOnGround 返回玩家是否在地面上。
func (p *Player) GetOnGround(ctx context.Context) (onGround bool, ok bool, err error) {
	return client_convertutil.BoolValue(p.client.GetPlayerOnGround(ctx, p.ref))
}

// GetRiddenEntityRuntimeID 返回玩家正在骑乘实体的运行时 ID。
func (p *Player) GetRiddenEntityRuntimeID(ctx context.Context) (riddenEntityRuntimeID uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(p.client.GetPlayerRiddenEntityRuntimeID(ctx, p.ref))
}

// GetTick 返回玩家状态最后更新时间刻。
func (p *Player) GetTick(ctx context.Context) (tick uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(p.client.GetPlayerTick(ctx, p.ref))
}

// GetAbilities 返回玩家权限和能力查询接口。
func (p *Player) GetAbilities() (abilities *PlayerAbilities) {
	return &PlayerAbilities{client: p.client, ref: p.ref}
}

// GetOnline 返回玩家是否在线。
func (p *Player) GetOnline(ctx context.Context) (online bool, ok bool, err error) {
	return client_convertutil.BoolValue(p.client.GetPlayerOnline(ctx, p.ref))
}

// PlayerAbilities 暴露玩家权限和能力层状态。
type PlayerAbilities struct {
	client resources_control_pb.UQHolderClient
	ref    *resources_control_pb.PlayerRef
}

// GetCommandPermissions 返回玩家命令权限等级。
func (a *PlayerAbilities) GetCommandPermissions(ctx context.Context) (commandPermissions byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(a.client.GetPlayerCommandPermissions(ctx, a.ref))
	return byte(value), ok, err
}

// GetPlayerPermissions 返回玩家权限等级。
func (a *PlayerAbilities) GetPlayerPermissions(ctx context.Context) (playerPermissions byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(a.client.GetPlayerPermissions(ctx, a.ref))
	return byte(value), ok, err
}

// GetFlySpeed 返回玩家飞行速度。
func (a *PlayerAbilities) GetFlySpeed(ctx context.Context) (flySpeed float32, ok bool, err error) {
	return client_convertutil.Float32Value(a.client.GetPlayerFlySpeed(ctx, a.ref))
}

// GetWalkSpeed 返回玩家行走速度。
func (a *PlayerAbilities) GetWalkSpeed(ctx context.Context) (walkSpeed float32, ok bool, err error) {
	return client_convertutil.Float32Value(a.client.GetPlayerWalkSpeed(ctx, a.ref))
}

// GetCanBuild 返回玩家是否可以放置方块。
func (a *PlayerAbilities) GetCanBuild(ctx context.Context) (canBuild bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanBuild(ctx, a.ref))
}

// GetCanMine 返回玩家是否可以挖掘方块。
func (a *PlayerAbilities) GetCanMine(ctx context.Context) (canMine bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanMine(ctx, a.ref))
}

// GetCanUseDoorsAndSwitches 返回玩家是否可以使用门和开关。
func (a *PlayerAbilities) GetCanUseDoorsAndSwitches(ctx context.Context) (canUseDoorsAndSwitches bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanUseDoorsAndSwitches(ctx, a.ref))
}

// GetCanOpenContainers 返回玩家是否可以打开容器。
func (a *PlayerAbilities) GetCanOpenContainers(ctx context.Context) (canOpenContainers bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanOpenContainers(ctx, a.ref))
}

// GetCanAttackPlayers 返回玩家是否可以攻击玩家。
func (a *PlayerAbilities) GetCanAttackPlayers(ctx context.Context) (canAttackPlayers bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanAttackPlayers(ctx, a.ref))
}

// GetCanAttackMobs 返回玩家是否可以攻击生物。
func (a *PlayerAbilities) GetCanAttackMobs(ctx context.Context) (canAttackMobs bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanAttackMobs(ctx, a.ref))
}

// GetCanUseOperatorCommands 返回玩家是否可以使用操作员命令。
func (a *PlayerAbilities) GetCanUseOperatorCommands(ctx context.Context) (canUseOperatorCommands bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanUseOperatorCommands(ctx, a.ref))
}

// GetCanTeleport 返回玩家是否可以传送。
func (a *PlayerAbilities) GetCanTeleport(ctx context.Context) (canTeleport bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanTeleport(ctx, a.ref))
}

// GetIsInvulnerable 返回玩家是否无敌。
func (a *PlayerAbilities) GetIsInvulnerable(ctx context.Context) (isInvulnerable bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsInvulnerable(ctx, a.ref))
}

// GetIsFlying 返回玩家是否正在飞行。
func (a *PlayerAbilities) GetIsFlying(ctx context.Context) (isFlying bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsFlying(ctx, a.ref))
}

// GetCanFly 返回玩家是否可以飞行。
func (a *PlayerAbilities) GetCanFly(ctx context.Context) (canFly bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanFly(ctx, a.ref))
}

// GetCanInstantBuild 返回玩家是否可以瞬间建造。
func (a *PlayerAbilities) GetCanInstantBuild(ctx context.Context) (canInstantBuild bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanInstantBuild(ctx, a.ref))
}

// GetCanUseLightning 返回玩家是否可以使用闪电。
func (a *PlayerAbilities) GetCanUseLightning(ctx context.Context) (canUseLightning bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanUseLightning(ctx, a.ref))
}

// GetIsMuted 返回玩家是否被禁言。
func (a *PlayerAbilities) GetIsMuted(ctx context.Context) (isMuted bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsMuted(ctx, a.ref))
}

// GetIsWorldBuilder 返回玩家是否为世界建造者。
func (a *PlayerAbilities) GetIsWorldBuilder(ctx context.Context) (isWorldBuilder bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsWorldBuilder(ctx, a.ref))
}

// GetHasNoClip 返回玩家是否启用无碰撞。
func (a *PlayerAbilities) GetHasNoClip(ctx context.Context) (hasNoClip bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerHasNoClip(ctx, a.ref))
}

// GetIsPrivilegedBuilder 返回玩家是否为特权建造者。
func (a *PlayerAbilities) GetIsPrivilegedBuilder(ctx context.Context) (isPrivilegedBuilder bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsPrivilegedBuilder(ctx, a.ref))
}
