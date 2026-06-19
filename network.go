package client

import (
	"context"
	"net"
)

// Network 表示 EmptyDeaCore gRPC 客户端支持的网络层实现。
type Network interface {
	// DialContext 尝试连接到指定地址。
	DialContext(ctx context.Context, address string) (net.Conn, error)
}

var networks = map[string]func() Network{}

// RegisterNetwork 注册一个网络实现。
func RegisterNetwork(id string, n func() Network) {
	networks[id] = n
}

func networkByID(id string) (Network, bool) {
	n, ok := networks[id]
	if ok {
		return n(), true
	}
	return nil, false
}

// TCP 是基于标准库 net.Dialer 的网络实现。
type TCP struct {
	network string
}

// DialContext 建立 TCP/Unix 连接。
func (t TCP) DialContext(ctx context.Context, address string) (net.Conn, error) {
	var d net.Dialer
	return d.DialContext(ctx, t.network, address)
}

func init() {
	RegisterNetwork("tcp", func() Network { return TCP{network: "tcp"} })
	RegisterNetwork("tcp4", func() Network { return TCP{network: "tcp4"} })
	RegisterNetwork("tcp6", func() Network { return TCP{network: "tcp6"} })
	RegisterNetwork("unix", func() Network { return TCP{network: "unix"} })
	RegisterNetwork("unixpacket", func() Network { return TCP{network: "unixpacket"} })
}
