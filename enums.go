package main

type InteractionType int
type InteractionCallbackType int
type ApplicationCommandType int

const (
	InteractionTypePing InteractionType = 1

	InteractionTypeApplicationCommand InteractionType = 2

	InteractionTypeMessageComponent InteractionType = 3

	InteractionTypeApplicationCommandAutocomplete InteractionType = 4

	InteractionTypeModalSubmit InteractionType = 5
)

const (
	InteractionCallbackTypePong InteractionCallbackType = 1

	InteractionCallbackTypeChannelMessageWithSource InteractionCallbackType = 4

	InteractionCallbackTypeDeferredChannelMessageWithSource InteractionCallbackType = 5

	InteractionCallbackTypeDeferredMessageUpdate InteractionCallbackType = 6

	InteractionCallbackTypeUpdateMessage InteractionCallbackType = 7

	InteractionCallbackTypeAutoCompleteResult InteractionCallbackType = 8

	InteractionCallbackTypeModal InteractionCallbackType = 9

	InteractionCallbackTypePremiumRequired InteractionCallbackType = 10
)

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1

	ApplicationCommandTypeUser ApplicationCommandType = 2

	ApplicationCommandTypeMessage ApplicationCommandType = 3
)
