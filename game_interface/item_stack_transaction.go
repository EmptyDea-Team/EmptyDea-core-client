package game_interface

import (
	item_stack_transaction_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_transaction"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	item_stack_transaction_client "github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_transaction"
)

// ItemStackOperation 是物品操作请求的包装实现。
type ItemStackOperation struct {
	client game_interface_pb.ItemStackTransactionServiceClient
}

// OpenTransaction 打开一个新的物品堆栈操作事务。
//
// 多个事务可以被同时打开，但各个事务的操作内容不
// 应该发生重叠，否则操作的结果是未定义的。
//
// 另外，同一个事务应当只能被同一个 go 惯例所使用，
// 这意味着同时并发使用同一个事务不保证线程安全性。
func (i *ItemStackOperation) OpenTransaction() item_stack_transaction_api.ItemStackTransaction {
	return item_stack_transaction_client.NewItemStackTransaction(i.client)
}
