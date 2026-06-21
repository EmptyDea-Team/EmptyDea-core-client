package resources_control

import (
	resources_control_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	uqholder_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control/uqholder"
)

var (
	_ resources_control_api.Resources = (*Resources)(nil)

	_ resources_control_api.Inventories      = (*Inventories)(nil)
	_ resources_control_api.Inventory        = (*Inventory)(nil)
	_ resources_control_api.ContainerManager = (*ContainerManager)(nil)
	_ resources_control_api.ConstantPacket   = (*ConstantPacket)(nil)
	_ resources_control_api.PacketListener   = (*PacketListener)(nil)

	_ uqholder_api.UQHolder = (*UQHolder)(nil)

	_ uqholder_api.Bot             = (*Bot)(nil)
	_ uqholder_api.Players         = (*Players)(nil)
	_ uqholder_api.Player          = (*Player)(nil)
	_ uqholder_api.PlayerAbilities = (*PlayerAbilities)(nil)
	_ uqholder_api.World           = (*World)(nil)
	_ uqholder_api.GameRule        = (*GameRule)(nil)
	_ uqholder_api.Entities        = (*Entities)(nil)
	_ uqholder_api.Entity          = (*Entity)(nil)
	_ uqholder_api.MobEffect       = (*MobEffect)(nil)
)
