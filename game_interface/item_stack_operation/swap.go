package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Swap 指示物品交换操作。
type Swap struct {
	Source      resources_control.SlotLocation
	Destination resources_control.SlotLocation
}

func (Swap) ID() uint8 {
	return IDItemStackOperationSwap
}

func (Swap) CanInline() bool {
	return true
}

// SwapBetweenInventory 指示背包内两个槽位之间的物品交换操作。
type SwapBetweenInventory struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
}

func (SwapBetweenInventory) ID() uint8 {
	return IDItemStackOperationSwap
}

func (SwapBetweenInventory) CanInline() bool {
	return true
}

// SwapInventoryBetweenContainer 指示背包和已打开容器之间的物品交换操作。
type SwapInventoryBetweenContainer struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
}

func (SwapInventoryBetweenContainer) ID() uint8 {
	return IDItemStackOperationSwap
}

func (SwapInventoryBetweenContainer) CanInline() bool {
	return true
}
