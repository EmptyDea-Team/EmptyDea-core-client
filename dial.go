package client

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Dialer 允许指定连接 EmptyDeaCore gRPC 服务端时使用的设置。
type Dialer struct {
	// Options 是传给 grpc.DialContext 的额外拨号选项。
	Options []grpc.DialOption
}

// DialContext 通过指定网络连接 EmptyDeaCore gRPC 服务端并返回客户端入口。
func (d Dialer) DialContext(ctx context.Context, network string, address string) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	n, ok := networkByID(network)
	if !ok {
		return nil, &net.OpError{Op: "dial", Net: "emptydea", Err: fmt.Errorf("dial: no network under id %v", network)}
	}
	opts := append([]grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return n.DialContext(ctx, address)
		}),
	}, d.Options...)
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	return New(conn), nil
}

// DialContext 通过指定网络连接 EmptyDeaCore gRPC 服务端。
func DialContext(ctx context.Context, network, address string) (*Client, error) {
	return Dialer{}.DialContext(ctx, network, address)
}

// DialTimeout 通过指定网络连接 EmptyDeaCore gRPC 服务端，并使用 timeout 限制拨号耗时。
func DialTimeout(network, address string, timeout time.Duration) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return DialContext(ctx, network, address)
}

// Dial 通过指定网络连接 EmptyDeaCore gRPC 服务端。
func Dial(network, address string) (*Client, error) {
	return DialContext(context.Background(), network, address)
}
