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
	ChannelTypes []int                        `json:"channel_types,omitempty"`
}

type ApplicationCommand struct {
	Name        string                                      `json:"name"`
	Type        ApplicationCommandType                      `json:"type"`
	Description string                                      `json:"description,omitempty"`
	Options     []Option                                    `json:"options,omitempty"`
	Handler     func(interaction *Interaction) `json:"-"`
}
