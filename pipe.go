package client

import (
	"context"
	"net"
)

// Pipe 是跨进程本机 IPC 网络实现。
type Pipe struct{}

// DialContext 连接到指定地址对应的本机 IPC 监听器。
func (Pipe) DialContext(ctx context.Context, address string) (net.Conn, error) {
	return PipeDialContext(ctx, address)
}

// PipeDialer 返回可传给 grpc.WithContextDialer 的本机 IPC 拨号器。
func PipeDialer(address string) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, _ string) (net.Conn, error) {
		return PipeDialContext(ctx, address)
	}
}

// PipeDialContext 连接到指定地址对应的本机 IPC 监听器。
func PipeDialContext(ctx context.Context, address string) (net.Conn, error) {
	return dialPipeContext(ctx, address)
}

// PipeAddress 返回指定地址规范化后的完整本机 IPC 地址。
func PipeAddress(address string) (string, error) {
	return pipeAddress(address)
}

func init() {
	RegisterNetwork("pipe", func() Network { return Pipe{} })
}
