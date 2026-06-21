package resources_control

import "github.com/EmptyDea-Team/EmptyDea-core-api/define"

var (
	_ define.Resources[
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

	_ define.Inventories[*Inventory] = (*Inventories)(nil)
	_ define.Inventory               = (*Inventory)(nil)
	_ define.ContainerManager        = (*ContainerManager)(nil)
	_ define.ConstantPacket          = (*ConstantPacket)(nil)
	_ define.PacketListener          = (*PacketListener)(nil)

	_ define.UQHolder[
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

	_ define.Bot                                = (*Bot)(nil)
	_ define.Players[*Player, *PlayerAbilities] = (*Players)(nil)
	_ define.UQPlayer[*PlayerAbilities]         = (*Player)(nil)
	_ define.PlayerAbilities                    = (*PlayerAbilities)(nil)
	_ define.World[*GameRule]                   = (*World)(nil)
	_ define.GameRule                           = (*GameRule)(nil)
	_ define.Entities[*Entity, *MobEffect]      = (*Entities)(nil)
	_ define.Entity[*MobEffect]                 = (*Entity)(nil)
	_ define.MobEffect                          = (*MobEffect)(nil)
)
