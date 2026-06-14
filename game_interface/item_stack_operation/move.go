package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// Move 指示物品移动操作。
type Move struct {
	Source      resources_control.SlotLocation
	Destination resources_control.SlotLocation
	Count       uint8
}

func (Move) ID() uint8 {
	return IDItemStackOperationMove
}

func (Move) CanInline() bool {
	return true
}

// MoveBetweenInventory 指示背包内两个槽位之间的物品移动操作。
type MoveBetweenInventory struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
	Count       uint8
}

func (MoveBetweenInventory) ID() uint8 {
	return IDItemStackOperationMove
}

func (MoveBetweenInventory) CanInline() bool {
	return true
}

// MoveBetweenContainer 指示已打开容器内两个槽位之间的物品移动操作。
type MoveBetweenContainer struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
	Count       uint8
}

func (MoveBetweenContainer) ID() uint8 {
	return IDItemStackOperationMove
}

func (MoveBetweenContainer) CanInline() bool {
	return true
}

// MoveToContainer 指示从背包移动物品到已打开容器的操作。
type MoveToContainer struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
	Count       uint8
}

func (MoveToContainer) ID() uint8 {
	return IDItemStackOperationMove
}

func (MoveToContainer) CanInline() bool {
	return true
}

// MoveToInventory 指示从已打开容器移动物品到背包的操作。
type MoveToInventory struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
	Count       uint8
}

func (MoveToInventory) ID() uint8 {
	return IDItemStackOperationMove
}

func (MoveToInventory) CanInline() bool {
	return true
}

// MoveToCraftingTable 指示从背包移动物品到合成栏的操作。
type MoveToCraftingTable struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
	Count       uint8
}

func (MoveToCraftingTable) ID() uint8 {
	return IDItemStackOperationMove
}

func (MoveToCraftingTable) CanInline() bool {
	return true
}

// MoveFromCraftingTable 指示从合成栏移动物品到背包的操作。
type MoveFromCraftingTable struct {
	Source      resources_control.SlotID
	Destination resources_control.SlotID
	Count       uint8
}

func (MoveFromCraftingTable) ID() uint8 {
	return IDItemStackOperationMove
}

func (MoveFromCraftingTable) CanInline() bool {
	return true
}
