package resources_control

import (
	"context"
	"slices"

	resources_control_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/resources_control"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// ConstantPacket 记载远程登录序列常量。
type ConstantPacket struct {
	client resources_control_pb.ConstantPacketServiceClient
}

// AllCreativeContent 返回租赁服在登录序列发送的创造物品数据。
func (c *ConstantPacket) AllCreativeContent(ctx context.Context) (items []*protocol_pb.CreativeItem, err error) {
	resp, err := c.client.ListCreativeItems(ctx, &resources_control_pb.ListCreativeItemsRequest{})
	if err != nil {
		return nil, err
	}
	items = make([]*protocol_pb.CreativeItem, 0, len(resp.Items))
	for _, item := range resp.Items {
		if item != nil {
			items = append(items, item)
		}
	}
	return items, nil
}

// CreativeItemByCNI 返回创造物品网络 ID 为 creativeNetworkID 的创造物品。
func (c *ConstantPacket) CreativeItemByCNI(ctx context.Context, creativeNetworkID uint32) (item *protocol_pb.CreativeItem, existed bool, err error) {
	resp, err := c.client.GetCreativeItemByCNI(ctx, &resources_control_pb.GetCreativeItemByCNIRequest{CreativeNetworkID: creativeNetworkID})
	if err != nil {
		return nil, false, err
	}
	if !resp.Found || resp.Item == nil {
		return nil, false, nil
	}
	return resp.Item, true, nil
}

// CreativeItemByNI 返回物品数字网络 ID 为 networkID 的多个创造物品。
func (c *ConstantPacket) CreativeItemByNI(ctx context.Context, networkID int32) (items []*protocol_pb.CreativeItem, err error) {
	resp, err := c.client.ListCreativeItemsByNetworkID(ctx, &resources_control_pb.ListCreativeItemsByNetworkIDRequest{NetworkID: networkID})
	if err != nil {
		return nil, err
	}
	items = make([]*protocol_pb.CreativeItem, 0, len(resp.Items))
	for _, item := range resp.Items {
		if item != nil {
			items = append(items, item)
		}
	}
	return items, nil
}

// CreativeItemByName 返回名称为 name 的多个创造物品。
func (c *ConstantPacket) CreativeItemByName(ctx context.Context, name string) (items []*protocol_pb.CreativeItem, err error) {
	resp, err := c.client.ListCreativeItemsByName(ctx, &resources_control_pb.ListCreativeItemsByNameRequest{Name: name})
	if err != nil {
		return nil, err
	}
	items = make([]*protocol_pb.CreativeItem, 0, len(resp.Items))
	for _, item := range resp.Items {
		if item != nil {
			items = append(items, item)
		}
	}
	return items, nil
}

// AllAvailableItems 返回租赁服在登录序列发送的所有可用物品。
func (c *ConstantPacket) AllAvailableItems(ctx context.Context) (items []*protocol_pb.ItemEntry, err error) {
	resp, err := c.client.ListAvailableItems(ctx, &resources_control_pb.ListAvailableItemsRequest{})
	if err != nil {
		return nil, err
	}
	items = make([]*protocol_pb.ItemEntry, 0, len(resp.Items))
	for _, item := range resp.Items {
		if item != nil {
			items = append(items, item)
		}
	}
	return items, nil
}

// ItemByNetworkID 返回网络 ID 为 networkID 的物品。
func (c *ConstantPacket) ItemByNetworkID(ctx context.Context, networkID int32) (item *protocol_pb.ItemEntry, existed bool, err error) {
	resp, err := c.client.GetItemByNetworkID(ctx, &resources_control_pb.GetItemByNetworkIDRequest{NetworkID: networkID})
	if err != nil {
		return nil, false, err
	}
	if !resp.Found || resp.Item == nil {
		return nil, false, nil
	}
	return resp.Item, true, nil
}

// ItemByName 返回物品名称为 name 的物品。
func (c *ConstantPacket) ItemByName(ctx context.Context, name string) (item *protocol_pb.ItemEntry, existed bool, err error) {
	resp, err := c.client.GetItemByName(ctx, &resources_control_pb.GetItemByNameRequest{Name: name})
	if err != nil {
		return nil, false, err
	}
	if !resp.Found || resp.Item == nil {
		return nil, false, nil
	}
	return resp.Item, true, nil
}

// ItemNameByNetworkID 返回网络 ID 为 networkID 的物品名称。
func (c *ConstantPacket) ItemNameByNetworkID(ctx context.Context, networkID int32) (name string, existed bool, err error) {
	resp, err := c.client.GetItemNameByNetworkID(ctx, &resources_control_pb.GetItemNameByNetworkIDRequest{NetworkID: networkID})
	if err != nil {
		return "", false, err
	}
	return resp.Name, resp.Found, nil
}

// AllCommandItems 返回可以通过指令获得的全部物品。
func (c *ConstantPacket) AllCommandItems(ctx context.Context) (items []string, err error) {
	resp, err := c.client.ListCommandItems(ctx, &resources_control_pb.ListCommandItemsRequest{})
	if err != nil {
		return nil, err
	}
	return slices.Clone(resp.Names), nil
}

// ItemCanGetByCommand 检查物品名为 name 的物品是否可以通过命令获取。
func (c *ConstantPacket) ItemCanGetByCommand(ctx context.Context, name string) (canGet bool, err error) {
	resp, err := c.client.ItemCanGetByCommand(ctx, &resources_control_pb.ItemCanGetByCommandRequest{Name: name})
	if err != nil {
		return false, err
	}
	return resp.CanGet, nil
}

// TrimRecipeNetworkID 返回锻造台纹饰操作对应合成配方的网络 ID。
func (c *ConstantPacket) TrimRecipeNetworkID(ctx context.Context) (networkID uint32, err error) {
	resp, err := c.client.GetTrimRecipeNetworkID(ctx, &resources_control_pb.GetTrimRecipeNetworkIDRequest{})
	if err != nil {
		return 0, err
	}
	return resp.NetworkID, nil
}
