package main

type ActionRow struct {
	Type       ComponentType `json:"type"`
	Components []Component   `json:"components"`
}

type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       struct {
		Name     string `json:"name,omitempty"`
		Id       string `json:"id,omitempty"`
		Animated bool   `json:"animated,omitempty"`
	} `json:"emoji,omitempty"`
	Default bool `json:"default,omitempty"`
}

type DefaultValue struct {
	Id   string           `json:"id"`
	Type DefaultValueType `json:"type"`
}

type Component struct {
	CustomId string        `json:"custom_id,omitempty"`
	Type     ComponentType `json:"type"`
	Label    string        `json:"label,omitempty"`
	Style    ButtonStyle   `json:"style,omitempty"`
	Emoji    struct {
		Name     string `json:"name,omitempty"`
		ID       string `json:"id,omitempty"`
		Animated bool   `json:"animated,omitempty"`
	} `json:"emoji,omitempty"`
	URL           string                         `json:"url,omitempty"`
	Options       []SelectOption                 `json:"options,omitempty"`
	ChannelTypes  []ChannelType                  `json:"channel_types,omitempty"`
	Placeholder   string                         `json:"placeholder,omitempty"`
	DefaultValues []DefaultValue                 `json:"default_values,omitempty"`
	MinValues     int                            `json:"min_values,omitempty"`
	MaxValues     int                            `json:"max_values,omitempty"`
	Disabled      bool                           `json:"disabled,omitempty"`
	Handler       func(interaction *Interaction) `json:"-"`
}

type ButtonConfig struct {
	Label string
	Style ButtonStyle
	Emoji struct {
		Name     string `json:"name,omitempty"`
		ID       string `json:"id,omitempty"`
		Animated bool   `json:"animated,omitempty"`
	}
	CustomId string
	URL      string
	Disabled bool
	Handler  func(interaction *Interaction)
}

func Button(config ButtonConfig) Component {
	return Component{
		Type:     ComponentTypeButton,
		Label:    config.Label,
		Style:    config.Style,
		Emoji:    config.Emoji,
		CustomId: config.CustomId,
		URL:      config.URL,
		Disabled: config.Disabled,
		Handler:  config.Handler,
	}
}

type SelectConfig struct {
	CustomId      string
	Type          SelectType
	Options       []SelectOption
	DefaultValues []DefaultValue
	ChannelTypes  []ChannelType
	Placeholder   string
	MinValues     int
	MaxValues     int
	Disabled      bool
	Handler       func(interaction *Interaction)
}

func Select(config SelectConfig) Component {
	return Component{
		Type:          ComponentType(config.Type),
		CustomId:      config.CustomId,
		Options:       config.Options,
		Placeholder:   config.Placeholder,
		DefaultValues: config.DefaultValues,
		MinValues:     config.MinValues,
		MaxValues:     config.MaxValues,
		Disabled:      config.Disabled,
		Handler:       config.Handler,
		ChannelTypes:  config.ChannelTypes,
	}
}
