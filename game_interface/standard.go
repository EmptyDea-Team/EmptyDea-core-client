package game_interface

import (
	game_interface_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface"
	item_stack_transaction_api "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_transaction"
	item_stack_transaction_client "github.com/EmptyDea-Team/EmptyDea-core-client/game_interface/item_stack_transaction"
)

var (
	_ game_interface_api.GameInterface = (*GameInterface)(nil)

	_ game_interface_api.Commands                     = (*Commands)(nil)
	_ game_interface_api.StructureBackup              = (*StructureBackup)(nil)
	_ game_interface_api.Querytarget                  = (*Querytarget)(nil)
	_ game_interface_api.Movement                     = (*Movement)(nil)
	_ game_interface_api.SetBlock                     = (*SetBlock)(nil)
	_ game_interface_api.Replaceitem                  = (*Replaceitem)(nil)
	_ game_interface_api.BotClick                     = (*BotClick)(nil)
	_ game_interface_api.ItemStackOperation           = (*ItemStackOperation)(nil)
	_ item_stack_transaction_api.ItemStackTransaction = (*item_stack_transaction_client.ItemStackTransaction)(nil)
	_ game_interface_api.ContainerOpenAndClose        = (*ContainerOpenAndClose)(nil)
	_ game_interface_api.ItemCopy                     = (*ItemCopy)(nil)
	_ game_interface_api.ItemTransition               = (*ItemTransition)(nil)
	_ game_interface_api.PlayerKit                    = (*PlayerKit)(nil)
	_ game_interface_api.Player                       = Player{}
	_ game_interface_api.AbilityBuilder               = (*AbilityBuilder)(nil)
)
