package main

type InteractionType int
type ApplicationCommandType int

const (
	InteractionTypePing InteractionType = 1

	InteractionTypeApplicationCommand InteractionType = 2

	InteractionTypeMessageComponent InteractionType = 3

	InteractionTypeApplicationCommandAutocomplete InteractionType = 4

	InteractionTypeModalSubmit InteractionType = 5
)

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1

	ApplicationCommandTypeUser ApplicationCommandType = 2

	ApplicationCommandTypeMessage ApplicationCommandType = 3
)
