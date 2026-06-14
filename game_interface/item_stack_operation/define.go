package item_stack_operation

const (
	IDItemStackOperationMove uint8 = iota
	IDItemStackOperationSwap
	IDItemStackOperationDrop
	IDItemStackOperationCreativeItem
	IDItemStackOperationHighLevelRenaming
	IDItemStackOperationHighLevelLooming
	IDItemStackOperationHighLevelCrafting
	IDItemStackOperationHighLevelTrimming
)

// ItemStackOperation 指示所有可提交给远端服务端执行的物品操作。
type ItemStackOperation interface {
	// CanInline 指示该物品操作是否可以内联到单个物品堆栈操作请求中。
	CanInline() bool
	// ID 指示该物品操作的编号，它是自定义的。
	ID() uint8
}
