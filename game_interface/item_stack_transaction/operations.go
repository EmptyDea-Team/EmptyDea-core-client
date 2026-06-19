package item_stack_transaction

import (
	"github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_operation"
	"github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"
)

// MoveItem 将 source 处的物品移动到 destination 处，
// 且只移动 count 个物品。
func (i *ItemStackTransaction) MoveItem(source resources_control.SlotLocation, destination resources_control.SlotLocation, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.Move{
		Source:      source,
		Destination: destination,
		Count:       count,
	})
	return i
}

// MoveBetweenInventory 将背包中 source 处的物品移动到 destination 处，
// 且只移动 count 个物品。
func (i *ItemStackTransaction) MoveBetweenInventory(source resources_control.SlotID, destination resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveBetweenInventory{
		Source:      source,
		Destination: destination,
		Count:       count,
	})
	return i
}

func dynamicContainerWindowName(dynamicContainerID resources_control.DynamicContainerID) resources_control.WindowName {
	return resources_control.NewDynamicContainerWindowName(dynamicContainerID)
}

// MoveBetweenDynamicContainer 将动态容器中 source 处的物品
// 移动到同一动态容器的 destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveBetweenDynamicContainer(
	dynamicContainerID resources_control.DynamicContainerID,
	source resources_control.SlotID,
	destination resources_control.SlotID,
	count uint8,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveBetweenDynamicContainer{
		DynamicContainerID: dynamicContainerID,
		Source:             source,
		Destination:        destination,
		Count:              count,
	})
	return i
}

// MoveToDynamicContainer 将背包中 source 处的物品移动到动态容器
// 的 destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveToDynamicContainer(
	source resources_control.SlotID,
	dynamicContainerID resources_control.DynamicContainerID,
	destination resources_control.SlotID,
	count uint8,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveToDynamicContainer{
		Source:             source,
		DynamicContainerID: dynamicContainerID,
		Destination:        destination,
		Count:              count,
	})
	return i
}

// MoveFromDynamicContainer 将动态容器中 source 处的物品移动到背包
// 的 destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveFromDynamicContainer(
	dynamicContainerID resources_control.DynamicContainerID,
	source resources_control.SlotID,
	destination resources_control.SlotID,
	count uint8,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveFromDynamicContainer{
		DynamicContainerID: dynamicContainerID,
		Source:             source,
		Destination:        destination,
		Count:              count,
	})
	return i
}

// MoveBetweenContainer 将已打开容器中 source 处的物品
// 移动到已打开容器的 destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveBetweenContainer(source resources_control.SlotID, destination resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveBetweenContainer{
		Source:      source,
		Destination: destination,
		Count:       count,
	})
	return i
}

// MoveToContainer 将背包中 source 处的物品移动到已打开容器的
// destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveToContainer(source resources_control.SlotID, destination resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveToContainer{
		Source:      source,
		Destination: destination,
		Count:       count,
	})
	return i
}

// MoveToInventory 将已打开容器中 source 处的物品移动到
// 背包的 destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveToInventory(source resources_control.SlotID, destination resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveToInventory{
		Source:      source,
		Destination: destination,
		Count:       count,
	})
	return i
}

// MoveToCraftingTable 将背包中 source 处的物品移动
// 到合成栏的 destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveToCraftingTable(source resources_control.SlotID, destination resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveToCraftingTable{
		Source:      source,
		Destination: destination,
		Count:       count,
	})
	return i
}

// MoveFromCraftingTable 将已放入合成栏 source 处的物品
// 移动到背包的 destination 处，且只移动 count 个物品。
func (i *ItemStackTransaction) MoveFromCraftingTable(source resources_control.SlotID, destination resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.MoveFromCraftingTable{
		Source:      source,
		Destination: destination,
		Count:       count,
	})
	return i
}

// SwapItem 交换 source 处和 destination 处的物品。
func (i *ItemStackTransaction) SwapItem(source resources_control.SlotLocation, destination resources_control.SlotLocation) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.Swap{
		Source:      source,
		Destination: destination,
	})
	return i
}

// SwapBetweenInventory 交换背包中 source
// 处和背包中 destination 处的物品。
func (i *ItemStackTransaction) SwapBetweenInventory(source resources_control.SlotID, destination resources_control.SlotID) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.SwapBetweenInventory{
		Source:      source,
		Destination: destination,
	})
	return i
}

// SwapBetweenDynamicContainer 交换同一动态容器中 source 和 destination 处的物品。
func (i *ItemStackTransaction) SwapBetweenDynamicContainer(
	dynamicContainerID resources_control.DynamicContainerID,
	source resources_control.SlotID,
	destination resources_control.SlotID,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.SwapBetweenDynamicContainer{
		DynamicContainerID: dynamicContainerID,
		Source:             source,
		Destination:        destination,
	})
	return i
}

// SwapInventoryBetweenDynamicContainer 交换背包中 source 处和动态容器
// destination 处的物品。
func (i *ItemStackTransaction) SwapInventoryBetweenDynamicContainer(
	source resources_control.SlotID,
	dynamicContainerID resources_control.DynamicContainerID,
	destination resources_control.SlotID,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.SwapInventoryBetweenDynamicContainer{
		Source:             source,
		DynamicContainerID: dynamicContainerID,
		Destination:        destination,
	})
	return i
}

// SwapInventoryBetweenContainer 交换背包中 source
// 处和已打开容器 destination 处的物品。
func (i *ItemStackTransaction) SwapInventoryBetweenContainer(source resources_control.SlotID, destination resources_control.SlotID) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.SwapInventoryBetweenContainer{
		Source:      source,
		Destination: destination,
	})
	return i
}

// DropItem 将 slot 处的物品丢出，且只丢出 count 个。
func (i *ItemStackTransaction) DropItem(slot resources_control.SlotLocation, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.Drop{
		Path:  slot,
		Count: count,
	})
	return i
}

// DropInventoryItem 将背包中 slot 处的
// 物品丢出，且只丢出 count 个。
func (i *ItemStackTransaction) DropInventoryItem(slot resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.DropInventoryItem{
		Slot:  slot,
		Count: count,
	})
	return i
}

// DropDynamicContainerItem 将动态容器 slot 处的物品丢出，且只丢出 count 个。
func (i *ItemStackTransaction) DropDynamicContainerItem(
	dynamicContainerID resources_control.DynamicContainerID,
	slot resources_control.SlotID,
	count uint8,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.DropDynamicContainerItem{
		DynamicContainerID: dynamicContainerID,
		Slot:               slot,
		Count:              count,
	})
	return i
}

// DropContainerItem 将已打开容器 slot 处的物品丢出，且只丢出 count 个。
func (i *ItemStackTransaction) DropContainerItem(slot resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.DropContainerItem{
		Slot:  slot,
		Count: count,
	})
	return i
}

// GetCreativeItem 从创造物品栏获取 创造物品网络 ID 为
// creativeItemNetworkID 的物品到 slot 处，且只移动 count 个物品。
func (i *ItemStackTransaction) GetCreativeItem(creativeItemNetworkID uint32, slot resources_control.SlotLocation, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.CreativeItem{
		CreativeItemNetworkID: creativeItemNetworkID,
		Path:                  slot,
		Count:                 count,
	})
	return i
}

// GetCreativeItemToInventory 从创造物品栏获取创造物品网络
// ID 为 creativeItemNetworkID 的物品到背包中的 slot 处，且只移动 count 个物品。
func (i *ItemStackTransaction) GetCreativeItemToInventory(creativeItemNetworkID uint32, slot resources_control.SlotID, count uint8) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.CreativeItemToInventory{
		CreativeItemNetworkID: creativeItemNetworkID,
		Slot:                  slot,
		Count:                 count,
	})
	return i
}

// GetCreativeItemToDynamicContainer 从创造物品栏获取创造物品网络 ID 为
// creativeItemNetworkID 的物品到动态容器 slot 处，且只移动 count 个物品。
func (i *ItemStackTransaction) GetCreativeItemToDynamicContainer(
	creativeItemNetworkID uint32,
	dynamicContainerID resources_control.DynamicContainerID,
	slot resources_control.SlotID,
	count uint8,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.CreativeItemToDynamicContainer{
		CreativeItemNetworkID: creativeItemNetworkID,
		DynamicContainerID:    dynamicContainerID,
		Slot:                  slot,
		Count:                 count,
	})
	return i
}

// RenameItem 将 slot 处的物品全部重命名为 newName。
func (i *ItemStackTransaction) RenameItem(slot resources_control.SlotLocation, newName string) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.Renaming{
		Path:    slot,
		NewName: newName,
	})
	return i
}

// RenameInventoryItem 将背包中 slot 处的物品全部重命名为 newName。
func (i *ItemStackTransaction) RenameInventoryItem(slot resources_control.SlotID, newName string) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.RenameInventoryItem{
		Slot:    slot,
		NewName: newName,
	})
	return i
}

// Looming 将 patternSlot 处的旗帜放入织布机中，
// 并通过使用 dyeSlot 处的染料合成新旗帜。
func (i *ItemStackTransaction) Looming(
	patternName string,
	patternSlot resources_control.SlotLocation,
	bannerSlot resources_control.SlotLocation,
	dyeSlot resources_control.SlotLocation,
	resultItem item_stack_operation.ExpectedNewItem,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.Looming{
		PatternName: patternName,
		PatternPath: patternSlot,
		BannerPath:  bannerSlot,
		DyePath:     dyeSlot,
		ResultItem:  resultItem,
	})
	return i
}

// LoomingFromInventory 将背包中 patternSlot 处的旗帜放入织布机中，
// 并通过使用背包中 dyeSlot 处的染料合成新旗帜。
func (i *ItemStackTransaction) LoomingFromInventory(
	patternName string,
	patternSlot resources_control.SlotID,
	bannerSlot resources_control.SlotID,
	dyeSlot resources_control.SlotID,
	resultItem item_stack_operation.ExpectedNewItem,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.LoomingFromInventory{
		PatternName: patternName,
		PatternSlot: patternSlot,
		BannerSlot:  bannerSlot,
		DyeSlot:     dyeSlot,
		ResultItem:  resultItem,
	})
	return i
}

// Crafting 用于合成一个物品。
func (i *ItemStackTransaction) Crafting(
	recipeNetworkID uint32,
	resultSlotID resources_control.SlotID,
	resultCount uint8,
	resultItem item_stack_operation.ExpectedNewItem,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.Crafting{
		RecipeNetworkID: recipeNetworkID,
		ResultSlotID:    resultSlotID,
		ResultCount:     resultCount,
		ResultItem:      resultItem,
	})
	return i
}

// Trimming 将物品放置在锻造台中，并进行对应的锻造台纹饰操作。
func (i *ItemStackTransaction) Trimming(
	trimItemPath resources_control.SlotLocation,
	materialPath resources_control.SlotLocation,
	templatePath resources_control.SlotLocation,
	resultItem item_stack_operation.ExpectedNewItem,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.Trimming{
		TrimItem:   trimItemPath,
		Material:   materialPath,
		Template:   templatePath,
		ResultItem: resultItem,
	})
	return i
}

// TrimmingFromInventory 将位于背包中的物品放置在锻造台中，
// 并进行对应的锻造台纹饰操作。
func (i *ItemStackTransaction) TrimmingFromInventory(
	trimItemSlot resources_control.SlotID,
	materialSlot resources_control.SlotID,
	templateSlot resources_control.SlotID,
	resultItem item_stack_operation.ExpectedNewItem,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.TrimmingFromInventory{
		TrimItemSlot: trimItemSlot,
		MaterialSlot: materialSlot,
		TemplateSlot: templateSlot,
		ResultItem:   resultItem,
	})
	return i
}

// BeaconPayment 将 paymentPath 处的 1 个物品作为信标支付物品，
// 并提交信标主/副效果的选择。
func (i *ItemStackTransaction) BeaconPayment(
	paymentPath resources_control.SlotLocation,
	primaryEffect int32,
	secondaryEffect int32,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.BeaconPayment{
		PaymentPath:     paymentPath,
		PrimaryEffect:   primaryEffect,
		SecondaryEffect: secondaryEffect,
	})
	return i
}

// BeaconPaymentFromInventory 将背包中 paymentSlot 处的 1 个物品
// 作为信标支付物品，并提交信标主/副效果的选择。
func (i *ItemStackTransaction) BeaconPaymentFromInventory(
	paymentSlot resources_control.SlotID,
	primaryEffect int32,
	secondaryEffect int32,
) *ItemStackTransaction {
	i.operations = append(i.operations, item_stack_operation.BeaconPaymentFromInventory{
		PaymentSlot:     paymentSlot,
		PrimaryEffect:   primaryEffect,
		SecondaryEffect: secondaryEffect,
	})
	return i
}
