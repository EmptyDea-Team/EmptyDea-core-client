package game_interface

import (
	"context"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

// Querytarget 是远程 querytarget 实现。
type Querytarget struct {
	client game_interface_pb.QuerytargetServiceClient
}

// DoQuerytarget 对 target 执行 querytarget 查询并返回查询结果列表。
func (q *Querytarget) DoQuerytarget(ctx context.Context, target string) ([]*game_interface_pb.TargetQueryingInfo, error) {
	resp, err := q.client.DoQuerytarget(ctx, &game_interface_pb.DoQuerytargetRequest{Target: target})
	if err != nil {
		return nil, err
	}
	return resp.Results, nil
}
