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

func (p *Player) GetUUIDString(ctx context.Context) (uuid string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerUUID(ctx, p.ref))
}

func (p *Player) GetName(ctx context.Context) (name string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerName(ctx, p.ref))
}

func (p *Player) GetXUID(ctx context.Context) (xuid string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerXUID(ctx, p.ref))
}

func (p *Player) GetEntityUniqueID(ctx context.Context) (entityUniqueID int64, ok bool, err error) {
	return client_convertutil.Int64Value(p.client.GetPlayerEntityUniqueID(ctx, p.ref))
}

func (p *Player) GetEntityRuntimeID(ctx context.Context) (entityRuntimeID uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(p.client.GetPlayerEntityRuntimeID(ctx, p.ref))
}

func (p *Player) GetNeteaseUID(ctx context.Context) (neteaseUID int64, ok bool, err error) {
	return client_convertutil.Int64Value(p.client.GetPlayerNeteaseUID(ctx, p.ref))
}

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

func (p *Player) GetPlatformChatID(ctx context.Context) (platformChatID string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerPlatformChatID(ctx, p.ref))
}

func (p *Player) GetBuildPlatform(ctx context.Context) (buildPlatform int32, ok bool, err error) {
	return client_convertutil.Int32Value(p.client.GetPlayerBuildPlatform(ctx, p.ref))
}

func (p *Player) GetSkinID(ctx context.Context) (skinID string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerSkinID(ctx, p.ref))
}

func (p *Player) GetDeviceID(ctx context.Context) (deviceID string, ok bool, err error) {
	return client_convertutil.StringValue(p.client.GetPlayerDeviceID(ctx, p.ref))
}

func (p *Player) GetEntityMetadata(ctx context.Context) (metadata *protocol_pb.EntityMetadata, ok bool, err error) {
	resp, err := p.client.GetPlayerEntityMetadata(ctx, p.ref)
	if err != nil {
		return nil, false, err
	}
	return resp.Value, resp.Ok, nil
}

func (p *Player) GetPosition(ctx context.Context) (pos *mgl32_pb.Vec3, ok bool, err error) {
	return client_convertutil.Vec3Value(p.client.GetPlayerPosition(ctx, p.ref))
}

func (p *Player) GetPitch(ctx context.Context) (pitch float32, ok bool, err error) {
	return client_convertutil.Float32Value(p.client.GetPlayerPitch(ctx, p.ref))
}

func (p *Player) GetYaw(ctx context.Context) (yaw float32, ok bool, err error) {
	return client_convertutil.Float32Value(p.client.GetPlayerYaw(ctx, p.ref))
}

func (p *Player) GetHeadYaw(ctx context.Context) (headYaw float32, ok bool, err error) {
	return client_convertutil.Float32Value(p.client.GetPlayerHeadYaw(ctx, p.ref))
}

func (p *Player) GetMoveMode(ctx context.Context) (moveMode byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(p.client.GetPlayerMoveMode(ctx, p.ref))
	return byte(value), ok, err
}

func (p *Player) GetOnGround(ctx context.Context) (onGround bool, ok bool, err error) {
	return client_convertutil.BoolValue(p.client.GetPlayerOnGround(ctx, p.ref))
}

func (p *Player) GetRiddenEntityRuntimeID(ctx context.Context) (riddenEntityRuntimeID uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(p.client.GetPlayerRiddenEntityRuntimeID(ctx, p.ref))
}

func (p *Player) GetTick(ctx context.Context) (tick uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(p.client.GetPlayerTick(ctx, p.ref))
}

func (p *Player) GetAbilities() (abilities *PlayerAbilities) {
	return &PlayerAbilities{client: p.client, ref: p.ref}
}

func (p *Player) GetOnline(ctx context.Context) (online bool, ok bool, err error) {
	return client_convertutil.BoolValue(p.client.GetPlayerOnline(ctx, p.ref))
}

// PlayerAbilities 暴露玩家权限和能力层状态。
type PlayerAbilities struct {
	client resources_control_pb.UQHolderClient
	ref    *resources_control_pb.PlayerRef
}

func (a *PlayerAbilities) GetCommandPermissions(ctx context.Context) (commandPermissions byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(a.client.GetPlayerCommandPermissions(ctx, a.ref))
	return byte(value), ok, err
}

func (a *PlayerAbilities) GetPlayerPermissions(ctx context.Context) (playerPermissions byte, ok bool, err error) {
	value, ok, err := client_convertutil.Uint32Value(a.client.GetPlayerPermissions(ctx, a.ref))
	return byte(value), ok, err
}

func (a *PlayerAbilities) GetFlySpeed(ctx context.Context) (flySpeed float32, ok bool, err error) {
	return client_convertutil.Float32Value(a.client.GetPlayerFlySpeed(ctx, a.ref))
}

func (a *PlayerAbilities) GetWalkSpeed(ctx context.Context) (walkSpeed float32, ok bool, err error) {
	return client_convertutil.Float32Value(a.client.GetPlayerWalkSpeed(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanBuild(ctx context.Context) (canBuild bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanBuild(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanMine(ctx context.Context) (canMine bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanMine(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanUseDoorsAndSwitches(ctx context.Context) (canUseDoorsAndSwitches bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanUseDoorsAndSwitches(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanOpenContainers(ctx context.Context) (canOpenContainers bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanOpenContainers(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanAttackPlayers(ctx context.Context) (canAttackPlayers bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanAttackPlayers(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanAttackMobs(ctx context.Context) (canAttackMobs bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanAttackMobs(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanUseOperatorCommands(ctx context.Context) (canUseOperatorCommands bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanUseOperatorCommands(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanTeleport(ctx context.Context) (canTeleport bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanTeleport(ctx, a.ref))
}

func (a *PlayerAbilities) GetIsInvulnerable(ctx context.Context) (isInvulnerable bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsInvulnerable(ctx, a.ref))
}

func (a *PlayerAbilities) GetIsFlying(ctx context.Context) (isFlying bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsFlying(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanFly(ctx context.Context) (canFly bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanFly(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanInstantBuild(ctx context.Context) (canInstantBuild bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanInstantBuild(ctx, a.ref))
}

func (a *PlayerAbilities) GetCanUseLightning(ctx context.Context) (canUseLightning bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerCanUseLightning(ctx, a.ref))
}

func (a *PlayerAbilities) GetIsMuted(ctx context.Context) (isMuted bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsMuted(ctx, a.ref))
}

func (a *PlayerAbilities) GetIsWorldBuilder(ctx context.Context) (isWorldBuilder bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsWorldBuilder(ctx, a.ref))
}

func (a *PlayerAbilities) GetHasNoClip(ctx context.Context) (hasNoClip bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerHasNoClip(ctx, a.ref))
}

func (a *PlayerAbilities) GetIsPrivilegedBuilder(ctx context.Context) (isPrivilegedBuilder bool, ok bool, err error) {
	return client_convertutil.BoolValue(a.client.GetPlayerIsPrivilegedBuilder(ctx, a.ref))
}
