package item_stack_transaction

import (
	"fmt"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	"github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_operation"
	"github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"
	"google.golang.org/protobuf/types/known/structpb"
)

func slotLocationToProto(src resources_control.SlotLocation) *resources_control_pb.SlotLocation {
	return &resources_control_pb.SlotLocation{
		WindowName: &resources_control_pb.WindowName{
			WindowID:           int32(src.WindowID),
			DynamicContainerID: uint32(src.DynamicContainerID),
		},
		SlotID: uint32(src.SlotID),
	}
}

func inventorySlotLocation(slot resources_control.SlotID) resources_control.SlotLocation {
	return resources_control.SlotLocation{
		WindowName: resources_control.WindowNameInventory,
		SlotID:     slot,
	}
}

func moveBySlot(id string, source resources_control.SlotID, destination resources_control.SlotID, count uint8) *game_interface_pb.MoveBySlotRequest {
	return &game_interface_pb.MoveBySlotRequest{
		TransactionID: id,
		SrcSlot:       uint32(source),
		DstSlot:       uint32(destination),
		Count:         uint32(count),
	}
}

func expectedNewItemToProto(src item_stack_operation.ExpectedNewItem) *game_interface_pb.ExpectedNewItem {
	return &game_interface_pb.ExpectedNewItem{
		ItemType: &game_interface_pb.ItemNewType{
			UseNetworkID: src.ItemType.UseNetworkID,
			NetworkID:    src.ItemType.NetworkID,
			UseMetadata:  src.ItemType.UseMetadata,
			Metadata:     src.ItemType.Metadata,
		},
		BlockRuntimeID: &game_interface_pb.ItemNewBlockRuntimeID{
			UseBlockRuntimeID: src.BlockRuntimeID.UseBlockRuntimeID,
			BlockRuntimeID:    src.BlockRuntimeID.BlockRuntimeID,
		},
		NBT: &game_interface_pb.ItemNewNBTData{
			UseNBTData:       src.NBT.UseNBTData,
			UseOriginDamage:  src.NBT.UseOriginDamage,
			NBTData:          structMapToProto(src.NBT.NBTData),
			ChangeRepairCost: src.NBT.ChangeRepairCost,
			RepairCostDelta:  src.NBT.RepairCostDelta,
			ChangeDamage:     src.NBT.ChangeDamage,
			DamageDelta:      src.NBT.DamageDelta,
		},
		Component: &game_interface_pb.ItemNewComponent{
			UseCanPlaceOn: src.Component.UseCanPlaceOn,
			CanPlaceOn:    src.Component.CanPlaceOn,
			UseCanDestroy: src.Component.UseCanDestroy,
			CanDestroy:    src.Component.CanDestroy,
		},
	}
}

func structMapToProto(src map[string]any) map[string]*structpb.Value {
	result := make(map[string]*structpb.Value, len(src))
	for key, value := range src {
		converted, err := structpb.NewValue(value)
		if err != nil {
			continue
		}
		result[key] = converted
	}
	return result
}

func operationsToProto(src []item_stack_operation.ItemStackOperation) ([]*game_interface_pb.ItemStackOperation, error) {
	result := make([]*game_interface_pb.ItemStackOperation, 0, len(src))
	for _, op := range src {
		converted, err := operationToProto(op)
		if err != nil {
			return nil, err
		}
		result = append(result, converted)
	}
	return result, nil
}

func operationToProto(src item_stack_operation.ItemStackOperation) (*game_interface_pb.ItemStackOperation, error) {
	switch op := src.(type) {
	case item_stack_operation.Move:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_MoveItem{MoveItem: &game_interface_pb.MoveItemOperation{
			Source:      slotLocationToProto(op.Source),
			Destination: slotLocationToProto(op.Destination),
			Count:       uint32(op.Count),
		}}}, nil
	case item_stack_operation.MoveBetweenInventory:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_MoveBetweenInventory{MoveBetweenInventory: &game_interface_pb.MoveBetweenInventoryOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
			Count:   uint32(op.Count),
		}}}, nil
	case item_stack_operation.MoveBetweenContainer:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_MoveBetweenContainer{MoveBetweenContainer: &game_interface_pb.MoveBetweenContainerOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
			Count:   uint32(op.Count),
		}}}, nil
	case item_stack_operation.MoveToContainer:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_MoveToContainer{MoveToContainer: &game_interface_pb.MoveToContainerOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
			Count:   uint32(op.Count),
		}}}, nil
	case item_stack_operation.MoveToInventory:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_MoveToInventory{MoveToInventory: &game_interface_pb.MoveToInventoryOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
			Count:   uint32(op.Count),
		}}}, nil
	case item_stack_operation.MoveToCraftingTable:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_MoveToCraftingTable{MoveToCraftingTable: &game_interface_pb.MoveToCraftingTableOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
			Count:   uint32(op.Count),
		}}}, nil
	case item_stack_operation.MoveFromCraftingTable:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_MoveFromCraftingTable{MoveFromCraftingTable: &game_interface_pb.MoveFromCraftingTableOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
			Count:   uint32(op.Count),
		}}}, nil
	case item_stack_operation.Swap:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_SwapItem{SwapItem: &game_interface_pb.SwapItemOperation{
			Source:      slotLocationToProto(op.Source),
			Destination: slotLocationToProto(op.Destination),
		}}}, nil
	case item_stack_operation.SwapBetweenInventory:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_SwapBetweenInventory{SwapBetweenInventory: &game_interface_pb.SwapBetweenInventoryOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
		}}}, nil
	case item_stack_operation.SwapInventoryBetweenContainer:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_SwapInventoryBetweenContainer{SwapInventoryBetweenContainer: &game_interface_pb.SwapInventoryBetweenContainerOperation{
			SrcSlot: uint32(op.Source),
			DstSlot: uint32(op.Destination),
		}}}, nil
	case item_stack_operation.Drop:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_DropItem{DropItem: &game_interface_pb.DropItemOperation{
			Slot:  slotLocationToProto(op.Path),
			Count: uint32(op.Count),
		}}}, nil
	case item_stack_operation.DropInventoryItem:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_DropInventoryItem{DropInventoryItem: &game_interface_pb.DropInventoryItemOperation{
			Slot:  uint32(op.Slot),
			Count: uint32(op.Count),
		}}}, nil
	case item_stack_operation.DropContainerItem:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_DropContainerItem{DropContainerItem: &game_interface_pb.DropContainerItemOperation{
			Slot:  uint32(op.Slot),
			Count: uint32(op.Count),
		}}}, nil
	case item_stack_operation.CreativeItem:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_GetCreativeItem{GetCreativeItem: &game_interface_pb.GetCreativeItemOperation{
			CreativeItemNetworkID: op.CreativeItemNetworkID,
			Destination:           slotLocationToProto(op.Path),
			Count:                 uint32(op.Count),
		}}}, nil
	case item_stack_operation.CreativeItemToInventory:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_GetCreativeItemToInventory{GetCreativeItemToInventory: &game_interface_pb.GetCreativeItemToInventoryOperation{
			CreativeItemNetworkID: op.CreativeItemNetworkID,
			Slot:                  uint32(op.Slot),
			Count:                 uint32(op.Count),
		}}}, nil
	case item_stack_operation.Renaming:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_RenameItem{RenameItem: &game_interface_pb.RenameItemOperation{
			Slot:    slotLocationToProto(op.Path),
			NewName: op.NewName,
		}}}, nil
	case item_stack_operation.RenameInventoryItem:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_RenameInventoryItem{RenameInventoryItem: &game_interface_pb.RenameInventoryItemOperation{
			Slot:    uint32(op.Slot),
			NewName: op.NewName,
		}}}, nil
	case item_stack_operation.Looming:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_Looming{Looming: &game_interface_pb.LoomingOperation{
			PatternName: op.PatternName,
			Pattern:     slotLocationToProto(op.PatternPath),
			Banner:      slotLocationToProto(op.BannerPath),
			Dye:         slotLocationToProto(op.DyePath),
			ResultItem:  expectedNewItemToProto(op.ResultItem),
		}}}, nil
	case item_stack_operation.LoomingFromInventory:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_LoomingFromInventory{LoomingFromInventory: &game_interface_pb.LoomingFromInventoryOperation{
			PatternName: op.PatternName,
			PatternSlot: uint32(op.PatternSlot),
			BannerSlot:  uint32(op.BannerSlot),
			DyeSlot:     uint32(op.DyeSlot),
			ResultItem:  expectedNewItemToProto(op.ResultItem),
		}}}, nil
	case item_stack_operation.Crafting:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_Crafting{Crafting: &game_interface_pb.CraftingOperation{
			RecipeNetworkID: op.RecipeNetworkID,
			ResultSlotID:    uint32(op.ResultSlotID),
			ResultCount:     uint32(op.ResultCount),
			ResultItem:      expectedNewItemToProto(op.ResultItem),
		}}}, nil
	case item_stack_operation.Trimming:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_Trimming{Trimming: &game_interface_pb.TrimmingOperation{
			Input:      slotLocationToProto(op.TrimItem),
			Material:   slotLocationToProto(op.Material),
			Template:   slotLocationToProto(op.Template),
			ResultItem: expectedNewItemToProto(op.ResultItem),
		}}}, nil
	case item_stack_operation.TrimmingFromInventory:
		return &game_interface_pb.ItemStackOperation{Operation: &game_interface_pb.ItemStackOperation_TrimmingFromInventory{TrimmingFromInventory: &game_interface_pb.TrimmingFromInventoryOperation{
			InputSlot:    uint32(op.TrimItemSlot),
			MaterialSlot: uint32(op.MaterialSlot),
			TemplateSlot: uint32(op.TemplateSlot),
			ResultItem:   expectedNewItemToProto(op.ResultItem),
		}}}, nil
	default:
		return nil, fmt.Errorf("unknown item stack transaction operation %T", src)
	}
}
