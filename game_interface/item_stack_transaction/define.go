package item_stack_transaction

import (
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	"github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_operation"
)

// ItemStackTransaction 是远程单个物品操作事务，
// 它希望使用者尽可能多的将物品堆栈请求内联在一个数据包中，
// 这样可以有效的节省操作的时间消耗。
type ItemStackTransaction struct {
	client     game_interface_pb.ItemStackTransactionServiceClient
	operations []item_stack_operation.ItemStackOperation
}

// NewItemStackTransaction 基于 client 创建并返回一个新的 ItemStackTransaction。
func NewItemStackTransaction(client game_interface_pb.ItemStackTransactionServiceClient) *ItemStackTransaction {
	return &ItemStackTransaction{client: client}
}
