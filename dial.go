package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Dial 连接 EmptyDeaCore gRPC 服务端并返回客户端入口。
func Dial(ctx context.Context, target string, opts ...grpc.DialOption) (*Client, *grpc.ClientConn, error) {
	if len(opts) == 0 {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, nil, err
	}
	return New(conn), conn, nil
}
