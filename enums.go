package main

type InteractionType int
type InteractionCallbackType int
type ApplicationCommandType int
type ApplicationCommandOptionType int
type ComponentType int
type SelectType int
type ButtonStyle int
type ChannelType int
type DefaultValueType string

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

	InteractionCallbackTypeDeferredUpdateMessage InteractionCallbackType = 6

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

const (
	ComponentTypeActionRow ComponentType = 1

	ComponentTypeButton ComponentType = 2

	ComponentTypeSelectMenu ComponentType = 3

	ComponentTypeTextInput ComponentType = 4

	ComponentTypeUserSelect ComponentType = 5

	ComponentTypeRoleSelect ComponentType = 6

	ComponentTypeMentionableSelect ComponentType = 7

	ComponentTypeChannelSelect ComponentType = 8
)

const (
	SelectTypeText SelectType = 3

	SelectTypeUser SelectType = 5

	SelectTypeRole SelectType = 6

	SelectTypeMentionable SelectType = 7

	SelectTypeChannel SelectType = 8
)

const (
	ButtonStylePrimary ButtonStyle = 1

	ButtonStyleSecondary ButtonStyle = 2

	ButtonStyleSuccess ButtonStyle = 3

	ButtonStyleDanger ButtonStyle = 4

	ButtonStyleLink ButtonStyle = 5
)

const (
	ChannelTypeGuildText ChannelType = 0

	ChannelTypeDm ChannelType = 1

	ChannelTypeGuildVoice ChannelType = 2

	ChannelTypeGroupDm ChannelType = 3

	ChannelTypeGuildCategory ChannelType = 4

	ChannelTypeGuildAnnouncement ChannelType = 5

	ChannelTypeGuildAnnouncementThread ChannelType = 10

	ChannelTypeGuildPublicThread ChannelType = 11

	ChannelTypeGuildPrivateThread ChannelType = 12

	ChannelTypeGuildStageVoice ChannelType = 13

	ChannelTypeGuildDirectory ChannelType = 14

	ChannelTypeGuildForum ChannelType = 15

	ChannelTypeGuildMedia ChannelType = 16
)

const (
	DefaultValueTypeUser DefaultValueType = "user"

	DefaultValueTypeRole DefaultValueType = "role"

	DefaultValueTypeChannel DefaultValueType = "channel"
)
