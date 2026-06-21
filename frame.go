package client

import (
	"context"

	frame_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame"
	api_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Frame 是框架层远程客户端。
type Frame struct {
	client api_pb.FrameServiceClient
}

type ConnectionState = frame_api.ConnectionState

type ClosedEvent = frame_api.ClosedEvent

// FrameConfig 是启动 Minecraft 连接所需的配置。
type FrameConfig = frame_api.FrameConfig

// NewFrame 创建框架层客户端。
func NewFrame(conn grpc.ClientConnInterface) *Frame {
	return &Frame{client: api_pb.NewFrameServiceClient(conn)}
}

var _ frame_api.Frame = (*Frame)(nil)

// StartConnection 使用配置启动连接。
func (f *Frame) StartConnection(ctx context.Context, config FrameConfig) (string, error) {
	resp, err := f.client.StartConnection(ctx, &api_pb.StartConnectionRequest{
		AuthServer:     config.AuthServer,
		UserToken:      config.UserToken,
		ServerCode:     config.ServerCode,
		ServerPassword: config.ServerPassword,
	})
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// StopConnection 主动关闭当前连接。
func (f *Frame) StopConnection(ctx context.Context) error {
	_, err := f.client.StopConnection(ctx, &api_pb.StopConnectionRequest{})
	return err
}

// GetConnectionState 查询当前连接是否仍然可用，以及关闭原因。
func (f *Frame) GetConnectionState(ctx context.Context) (ConnectionState, error) {
	resp, err := f.client.GetConnectionState(ctx, &api_pb.GetConnectionStateRequest{})
	if err != nil {
		return ConnectionState{}, err
	}
	return ConnectionState{Connected: resp.Connected, CloseReason: resp.CloseReason}, nil
}

// Ping 检查 API 服务是否可响应。
func (f *Frame) Ping(ctx context.Context) (bool, error) {
	resp, err := f.client.Ping(ctx, &emptypb.Empty{})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

// WatchClosed 监听连接关闭事件。
func (f *Frame) WatchClosed(ctx context.Context, callback func(ClosedEvent, error)) error {
	stream, err := f.client.WatchClosed(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	go func() {
		for {
			event, err := stream.Recv()
			if err != nil {
				callback(ClosedEvent{}, err)
				return
			}
			callback(ClosedEvent{Reason: event.Reason}, nil)
		}
	}()
	return nil
}
