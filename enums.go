package main

type InteractionType int
type InteractionCallbackType int
type ApplicationCommandType int
type ApplicationCommandOptionType int

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

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = 1

	ApplicationCommandOptionTypeSubCommandGroup ApplicationCommandOptionType = 2

	ApplicationCommandOptionTypeString ApplicationCommandOptionType = 3

	ApplicationCommandOptionTypeInteger ApplicationCommandOptionType = 4

	ApplicationCommandOptionTypeBoolean ApplicationCommandOptionType = 5

	ApplicationCommandOptionTypeUser ApplicationCommandOptionType = 6

	ApplicationCommandOptionTypeChannel ApplicationCommandOptionType = 7

	ApplicationCommandOptionTypeRole ApplicationCommandOptionType = 8

	ApplicationCommandOptionTypeMentionable ApplicationCommandOptionType = 9

	ApplicationCommandOptionTypeNumber ApplicationCommandOptionType = 10

	ApplicationCommandOptionTypeAttachment ApplicationCommandOptionType = 11
)