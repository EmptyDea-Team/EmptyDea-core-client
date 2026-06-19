package uqholder

import (
	"context"

	client_convertutil "github.com/EmptyDea-Team/EmptyDea-core-client/convertutil"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
)

// Entities 暴露实体状态的只读查询接口。
type Entities struct {
	client resources_control_pb.UQHolderClient
}

// GetByRuntimeID 根据实体运行时 ID 查询实体状态。
func (e *Entities) GetByRuntimeID(ctx context.Context, runtimeID uint64) (entity *Entity, existed bool, err error) {
	resp, err := e.client.GetEntityByRuntimeID(ctx, &resources_control_pb.EntityRef{RuntimeID: runtimeID})
	if err != nil {
		return nil, false, err
	}
	if !resp.Found || resp.Entity == nil {
		return nil, false, nil
	}
	return &Entity{client: e.client, ref: resp.Entity}, true, nil
}

// GetHealth 根据实体运行时 ID 返回实体生命值。
func (e *Entities) GetHealth(ctx context.Context, runtimeID uint64) (health float32, ok bool, err error) {
	return client_convertutil.Float32Value(e.client.GetEntityHealth(ctx, &resources_control_pb.EntityRef{RuntimeID: runtimeID}))
}

// GetEffectTypes 根据实体运行时 ID 返回实体当前药水效果类型列表。
func (e *Entities) GetEffectTypes(ctx context.Context, runtimeID uint64) (effectTypes []int32, existed bool, err error) {
	resp, err := e.client.GetEntityEffectTypes(ctx, &resources_control_pb.EntityRef{RuntimeID: runtimeID})
	if err != nil {
		return nil, false, err
	}
	return resp.EffectTypes, resp.Found, nil
}

// Entity 暴露按运行时 ID 维护的实体状态。
type Entity struct {
	client resources_control_pb.UQHolderClient
	ref    *resources_control_pb.EntityRef
}

// GetRuntimeID 返回实体运行时 ID。
func (e *Entity) GetRuntimeID(ctx context.Context) (runtimeID uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(e.client.GetEntityRuntimeID(ctx, e.ref))
}

// GetUniqueID 返回实体唯一 ID。
func (e *Entity) GetUniqueID(ctx context.Context) (uniqueID int64, ok bool, err error) {
	return client_convertutil.Int64Value(e.client.GetEntityUniqueID(ctx, e.ref))
}

// GetHealth 返回实体生命值。
func (e *Entity) GetHealth(ctx context.Context) (health float32, ok bool, err error) {
	return client_convertutil.Float32Value(e.client.GetEntityHealth(ctx, e.ref))
}

// GetEffectTypes 返回实体当前药水效果类型列表。
func (e *Entity) GetEffectTypes(ctx context.Context) (effectTypes []int32, existed bool, err error) {
	resp, err := e.client.GetEntityEffectTypes(ctx, e.ref)
	if err != nil {
		return nil, false, err
	}
	return resp.EffectTypes, resp.Found, nil
}

// GetEffect 根据效果类型查询实体身上的药水效果状态。
func (e *Entity) GetEffect(ctx context.Context, effectType int32) (mobEffect *MobEffect, existed bool, err error) {
	types, ok, err := e.GetEffectTypes(ctx)
	if err != nil || !ok {
		return nil, false, err
	}
	for _, value := range types {
		if value == effectType {
			return &MobEffect{client: e.client, ref: &resources_control_pb.MobEffectRef{RuntimeID: e.ref.RuntimeID, EffectType: effectType}}, true, nil
		}
	}
	return nil, false, nil
}

// MobEffect 暴露实体身上的单个药水效果状态。
type MobEffect struct {
	client resources_control_pb.UQHolderClient
	ref    *resources_control_pb.MobEffectRef
}

// GetEffectType 返回药水效果类型。
func (m *MobEffect) GetEffectType() (effectType int32) { return m.ref.EffectType }

// GetAmplifier 返回药水效果等级。
func (m *MobEffect) GetAmplifier(ctx context.Context) (amplifier int32, ok bool, err error) {
	return client_convertutil.Int32Value(m.client.GetEntityEffectAmplifier(ctx, m.ref))
}

// GetDuration 返回药水效果剩余时间。
func (m *MobEffect) GetDuration(ctx context.Context) (duration int32, ok bool, err error) {
	return client_convertutil.Int32Value(m.client.GetEntityEffectDuration(ctx, m.ref))
}

// GetParticles 返回药水效果是否显示粒子。
func (m *MobEffect) GetParticles(ctx context.Context) (particles bool, ok bool, err error) {
	return client_convertutil.BoolValue(m.client.GetEntityEffectParticles(ctx, m.ref))
}

// GetUpdatedTick 返回药水效果最后更新时间刻。
func (m *MobEffect) GetUpdatedTick(ctx context.Context) (updatedTick uint64, ok bool, err error) {
	return client_convertutil.Uint64Value(m.client.GetEntityEffectUpdatedTick(ctx, m.ref))
}
