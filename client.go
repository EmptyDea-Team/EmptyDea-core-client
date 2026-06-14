package client

import (
	gi_client "github.com/EmptyDea-Team/EmptyDea-core-client/game_interface"
	rc_client "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"
	"google.golang.org/grpc"
)

// Client 是 EmptyDeaCore gRPC API 客户端入口。
type Client struct {
	conn  grpc.ClientConnInterface
	frame *Frame
}

// New 基于 gRPC 连接创建客户端入口。
func New(conn grpc.ClientConnInterface) *Client {
	return &Client{conn: conn, frame: NewFrame(conn)}
}

// Frame 返回框架层服务客户端。
func (c *Client) Frame() *Frame {
	return c.frame
}

// Resources 返回资源层客户端集合。
func (c *Client) Resources() *rc_client.Resources {
	return rc_client.New(c.conn)
}

// GameInterface 返回游戏交互层客户端集合。
func (c *Client) GameInterface() *gi_client.GameInterface {
	return gi_client.New(c.conn)
}
