package resources_control

import (
	resources_control_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	uqholder_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control/uqholder"
)

var (
	_ resources_control_api.Resources[
		*Inventories,
		*Inventory,
		*ContainerManager,
		*ConstantPacket,
		*PacketListener,
		*UQHolder,
		*Bot,
		*Players,
		*Player,
		*PlayerAbilities,
		*World,
		*GameRule,
		*Entities,
		*Entity,
		*MobEffect,
	] = (*Resources)(nil)

	_ resources_control_api.Inventories[*Inventory] = (*Inventories)(nil)
	_ resources_control_api.Inventory               = (*Inventory)(nil)
	_ resources_control_api.ContainerManager        = (*ContainerManager)(nil)
	_ resources_control_api.ConstantPacket          = (*ConstantPacket)(nil)
	_ resources_control_api.PacketListener          = (*PacketListener)(nil)

	_ uqholder_api.UQHolder[
		*Bot,
		*Players,
		*Player,
		*PlayerAbilities,
		*World,
		*GameRule,
		*Entities,
		*Entity,
		*MobEffect,
	] = (*UQHolder)(nil)

	_ uqholder_api.Bot                                = (*Bot)(nil)
	_ uqholder_api.Players[*Player, *PlayerAbilities] = (*Players)(nil)
	_ uqholder_api.Player[*PlayerAbilities]           = (*Player)(nil)
	_ uqholder_api.PlayerAbilities                    = (*PlayerAbilities)(nil)
	_ uqholder_api.World[*GameRule]                   = (*World)(nil)
	_ uqholder_api.GameRule                           = (*GameRule)(nil)
	_ uqholder_api.Entities[*Entity, *MobEffect]      = (*Entities)(nil)
	_ uqholder_api.Entity[*MobEffect]                 = (*Entity)(nil)
	_ uqholder_api.MobEffect                          = (*MobEffect)(nil)
)
