package client

import (
	frame_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame"
	game_interface_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface"
	resources_control_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	gi_client "github.com/EmptyDea-Team/EmptyDea-core-client/game_interface"
	rc_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"
	"google.golang.org/grpc"
)

// Client 是 EmptyDeaCore gRPC API 客户端入口。
type Client struct {
	conn     grpc.ClientConnInterface
	grpcConn *grpc.ClientConn
	frame    *Frame
}

// New 基于 gRPC 连接创建客户端入口。
func New(conn grpc.ClientConnInterface) *Client {
	client := &Client{conn: conn, frame: NewFrame(conn)}
	if grpcConn, ok := conn.(*grpc.ClientConn); ok {
		client.grpcConn = grpcConn
	}
	return client
}

// Conn 返回底层 gRPC 连接。
func (c *Client) Conn() *grpc.ClientConn {
	return c.grpcConn
}

// Frame 返回框架层服务客户端。
func (c *Client) Frame() frame_api.Frame {
	return c.frame
}

// Resources 返回资源层客户端集合。
func (c *Client) Resources() resources_control_api.Resources {
	return rc_client.New(c.conn)
}

// GameInterface 返回游戏交互层客户端集合。
func (c *Client) GameInterface() game_interface_api.GameInterface {
	return gi_client.New(c.conn)
}

// Close 关闭底层 gRPC 连接。
func (c *Client) Close() error {
	if c.grpcConn == nil {
		return nil
	}
	return c.grpcConn.Close()
}

var _ frame_api.Client = (*Client)(nil)
