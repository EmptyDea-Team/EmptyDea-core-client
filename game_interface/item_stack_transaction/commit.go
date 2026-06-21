package item_stack_transaction

import (
	"context"

	item_stack_transaction_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_transaction"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
	"github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_operation"
)

// Discord 丢弃曾经执行的更改。
// 从本质上说，它将清空底层操作序列。
func (i *ItemStackTransaction) Discord() item_stack_transaction_api.ItemStackTransaction {
	for index := range i.operations {
		i.operations[index] = nil
	}
	i.operations = nil
	return i
}

// Discard 是 Discord 的别名，用于避免误拼写时降低可读性。
func (i *ItemStackTransaction) Discard() item_stack_transaction_api.ItemStackTransaction {
	return i.Discord()
}

// Commit 将底层操作序列提交给远端服务端执行物品堆栈操作事务。
// 如果没有返回错误，Commit 在完成后将使用 Discord 清空底层操作序列。
func (i *ItemStackTransaction) Commit(ctx context.Context) (
	success bool,
	pk *packet_pb.ItemStackRequest,
	serverResponse []*protocol_pb.ItemStackResponse,
	err error,
) {
	if len(i.operations) == 0 {
		return true, nil, make([]*protocol_pb.ItemStackResponse, 0), nil
	}

	resp, err := i.CommitItemStackOperations(ctx, i.operations)
	if err != nil {
		return false, nil, nil, err
	}
	_ = i.Discord()
	return resp.Success, resp.Packet, resp.Responses, nil
}

// CommitItemStackOperations 一次性提交物品堆栈操作列表，避免每个操作都发生一次 RPC。
func (i *ItemStackTransaction) CommitItemStackOperations(ctx context.Context, operations []item_stack_operation.ItemStackOperation) (*game_interface_pb.CommitTransactionResponse, error) {
	converted, err := operationsToProto(operations)
	if err != nil {
		return nil, err
	}
	return i.commitItemStackOperations(ctx, converted)
}

func (i *ItemStackTransaction) commitItemStackOperations(ctx context.Context, operations []*game_interface_pb.ItemStackOperation) (*game_interface_pb.CommitTransactionResponse, error) {
	return i.client.CommitItemStackOperations(ctx, &game_interface_pb.CommitItemStackOperationsRequest{Operations: operations})
}
