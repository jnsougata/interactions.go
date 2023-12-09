package main

type Choice struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type Option struct {
	Type         ApplicationCommandOptionType `json:"type"`
	Name         string                       `json:"name"`
	Value        any                          `json:"value,omitempty"`
	Focus        bool                         `json:"focused,omitempty"`
	Description  string                       `json:"description,omitempty"`
	Required     bool                         `json:"required,omitempty"`
	Choices      []Choice                     `json:"choices,omitempty"`
	MaxValue     float64                      `json:"max_value,omitempty"`
	MinValue     float64                      `json:"min_value,omitempty"`
	MaxLength    int                          `json:"max_length,omitempty"`
	MinLength    int                          `json:"min_length,omitempty"`
	Autocomplete bool                         `json:"autocomplete,omitempty"`
	ChannelTypes []ChannelType                `json:"channel_types,omitempty"`
}

type OptionConfigString struct {
	Name         string
	Description  string
	Required     bool
	Choices      []Choice
	MaxLength    int
	MinLength    int
	AutoComplete bool
}

type OptionConfigNumber struct {
	Name         string
	Description  string
	Required     bool
	Choices      []Choice
	MaxValue     float64
	MinValue     float64
	AutoComplete bool
}

type OptionConfigChannel struct {
	Name         string
	Description  string
	Required     bool
	ChannelTypes []ChannelType
}

type OptionConfigBasic struct {
	Name        string
	Description string
	Required    bool
}

func Options(opts ...Option) []Option {
	return opts
}

func StringOption(conf OptionConfigString) Option {
	return Option{
		Type:         ApplicationCommandOptionTypeString,
		Name:         conf.Name,
		Description:  conf.Description,
		Required:     conf.Required,
		Choices:      conf.Choices,
		MaxLength:    conf.MaxLength,
		MinLength:    conf.MinLength,
		Autocomplete: conf.AutoComplete,
	}
}

func IntegerOption(conf OptionConfigNumber) Option {
	return Option{
		Type:         ApplicationCommandOptionTypeInteger,
		Name:         conf.Name,
		Description:  conf.Description,
		Required:     conf.Required,
		Choices:      conf.Choices,
		MaxValue:     conf.MaxValue,
		MinValue:     conf.MinValue,
		Autocomplete: conf.AutoComplete,
	}
}

func BooleanOption(conf OptionConfigBasic) Option {
	return Option{
		Type:        ApplicationCommandOptionTypeBoolean,
		Name:        conf.Name,
		Description: conf.Description,
		Required:    conf.Required,
	}
}

func UserOption(conf OptionConfigBasic) Option {
	return Option{
		Type:        ApplicationCommandOptionTypeUser,
		Name:        conf.Name,
		Description: conf.Description,
		Required:    conf.Required,
	}
}

func ChannelOption(conf OptionConfigChannel) Option {
	return Option{
		Type:         ApplicationCommandOptionTypeChannel,
		Name:         conf.Name,
		Description:  conf.Description,
		Required:     conf.Required,
		ChannelTypes: conf.ChannelTypes,
	}
}

func RoleOption(conf OptionConfigBasic) Option {
	return Option{
		Type:        ApplicationCommandOptionTypeRole,
		Name:        conf.Name,
		Description: conf.Description,
		Required:    conf.Required,
	}
}

func MentionableOption(conf OptionConfigBasic) Option {
	return Option{
		Type:        ApplicationCommandOptionTypeMentionable,
		Name:        conf.Name,
		Description: conf.Description,
		Required:    conf.Required,
	}
}

func NumberOption(conf OptionConfigNumber) Option {
	return Option{
		Type:         ApplicationCommandOptionTypeNumber,
		Name:         conf.Name,
		Description:  conf.Description,
		Required:     conf.Required,
		Choices:      conf.Choices,
		MaxValue:     conf.MaxValue,
		MinValue:     conf.MinValue,
		Autocomplete: conf.AutoComplete,
	}
}

func AttachmentOption(conf OptionConfigBasic) Option {
	return Option{
		Type:        ApplicationCommandOptionTypeAttachment,
		Name:        conf.Name,
		Description: conf.Description,
		Required:    conf.Required,
	}
}

type ApplicationCommand struct {
	Name        string                               `json:"name"`
	Type        ApplicationCommandType               `json:"type"`
	Description string                               `json:"description,omitempty"`
	Options     []Option                             `json:"options,omitempty"`
	Handler     func(interaction *Interaction) error `json:"-"`
}

type UserCommandConfig struct {
	Name    string
	Handler func(interaction *Interaction) error
}

type MessageCommandConfig struct {
	Name    string
	Handler func(interaction *Interaction) error
}

type SlashCommandConfig struct {
	Name        string
	Description string
	Options     []Option
	Handler     func(interaction *Interaction) error
}

func UserCommand(conf UserCommandConfig) ApplicationCommand {
	return ApplicationCommand{
		Name:    conf.Name,
		Type:    ApplicationCommandTypeUser,
		Handler: conf.Handler,
	}
}

func MessageCommand(conf MessageCommandConfig) ApplicationCommand {
	return ApplicationCommand{
		Name:    conf.Name,
		Type:    ApplicationCommandTypeMessage,
		Handler: conf.Handler,
	}
}

func SlashCommand(conf SlashCommandConfig) ApplicationCommand {
	return ApplicationCommand{
		Name:        conf.Name,
		Type:        ApplicationCommandTypeChatInput,
		Description: conf.Description,
		Options:     conf.Options,
		Handler:     conf.Handler,
	}
}
