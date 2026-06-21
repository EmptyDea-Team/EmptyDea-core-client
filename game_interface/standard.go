package game_interface

import (
	"github.com/EmptyDea-Team/EmptyDea-core-api/define"
	"github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_operation"
	"github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_transaction"
)

var (
	_ define.GameInterface[
		*Commands,
		*StructureBackup,
		*Querytarget,
		*Movement,
		*SetBlock,
		*Replaceitem,
		*BotClick,
		*ItemStackOperation,
		*item_stack_transaction.ItemStackTransaction,
		item_stack_operation.ItemStackOperation,
		*ContainerOpenAndClose,
		*ItemCopy,
		*ItemTransition,
		*PlayerKit,
		Player,
		*AbilityBuilder,
	] = (*GameInterface)(nil)

	_ define.Commands                           = (*Commands)(nil)
	_ define.StructureBackup                    = (*StructureBackup)(nil)
	_ define.Querytarget                        = (*Querytarget)(nil)
	_ define.Movement                           = (*Movement)(nil)
	_ define.SetBlock                           = (*SetBlock)(nil)
	_ define.Replaceitem                        = (*Replaceitem)(nil)
	_ define.BotClick                           = (*BotClick)(nil)
	_ define.ContainerOpenAndClose              = (*ContainerOpenAndClose)(nil)
	_ define.ItemCopy                           = (*ItemCopy)(nil)
	_ define.ItemTransition                     = (*ItemTransition)(nil)
	_ define.PlayerKit[Player, *AbilityBuilder] = (*PlayerKit)(nil)
	_ define.Player[*AbilityBuilder]            = Player{}
	_ define.AbilityBuilder[*AbilityBuilder]    = (*AbilityBuilder)(nil)
)
