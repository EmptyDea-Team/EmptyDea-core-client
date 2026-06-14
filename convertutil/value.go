package convertutil

import (
	game_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control"
	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
)

// Empty 创建一个空请求。
func Empty() *game_control_pb.Empty {
	return &game_control_pb.Empty{}
}

// StringValue 解包可选 string 返回值。
func StringValue(resp *game_control_pb.StringValue, err error) (value string, ok bool, retErr error) {
	if err != nil || resp == nil {
		return "", false, err
	}
	return resp.Value, resp.Ok, nil
}

// Int32Value 解包可选 int32 返回值。
func Int32Value(resp *game_control_pb.Int32Value, err error) (value int32, ok bool, retErr error) {
	if err != nil || resp == nil {
		return 0, false, err
	}
	return resp.Value, resp.Ok, nil
}

// Int64Value 解包可选 int64 返回值。
func Int64Value(resp *game_control_pb.Int64Value, err error) (value int64, ok bool, retErr error) {
	if err != nil || resp == nil {
		return 0, false, err
	}
	return resp.Value, resp.Ok, nil
}

// Uint32Value 解包可选 uint32 返回值。
func Uint32Value(resp *game_control_pb.Uint32Value, err error) (value uint32, ok bool, retErr error) {
	if err != nil || resp == nil {
		return 0, false, err
	}
	return resp.Value, resp.Ok, nil
}

// Uint64Value 解包可选 uint64 返回值。
func Uint64Value(resp *game_control_pb.Uint64Value, err error) (value uint64, ok bool, retErr error) {
	if err != nil || resp == nil {
		return 0, false, err
	}
	return resp.Value, resp.Ok, nil
}

// BoolValue 解包可选 bool 返回值。
func BoolValue(resp *game_control_pb.BoolValue, err error) (value bool, ok bool, retErr error) {
	if err != nil || resp == nil {
		return false, false, err
	}
	return resp.Value, resp.Ok, nil
}

// Float32Value 解包可选 float32 返回值。
func Float32Value(resp *game_control_pb.Float32Value, err error) (value float32, ok bool, retErr error) {
	if err != nil || resp == nil {
		return 0, false, err
	}
	return resp.Value, resp.Ok, nil
}

// Vec3Value 解包可选 Vec3 返回值。
func Vec3Value(resp *resources_control_pb.Vec3Value, err error) (value *mgl32_pb.Vec3, ok bool, retErr error) {
	if err != nil || resp == nil || resp.Value == nil {
		return nil, false, err
	}
	return resp.Value, resp.Ok, nil
}
