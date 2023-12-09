package main

type TextInputOptions struct {
	Type        ComponentType  `json:"type"`
	CustomId    string         `json:"custom_id"`
	Style       TextInputStyle `json:"style"`
	Label       string         `json:"label"`
	MinLength   int            `json:"min_length,omitempty"`
	MaxLength   int            `json:"max_length,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Value       string         `json:"value,omitempty"`
	Placeholder string         `json:"placeholder,omitempty"`
}

func TextInput(o TextInputOptions) Component {
	o.Type = ComponentTypeTextInput
	return Component{
		Type:        o.Type,
		CustomId:    o.CustomId,
		Style:       ButtonStyle(o.Style),
		Label:       o.Label,
		MinLength:   o.MinLength,
		MaxLength:   o.MaxLength,
		Required:    o.Required,
		Value:       o.Value,
		Placeholder: o.Placeholder,
	}
}

type ModalConfig struct {
	Title    string                               `json:"title"`
	CustomId string                               `json:"custom_id"`
	Fields   []ActionRow                          `json:"components"`
	Handler  func(interaction *Interaction) error `json:"-"`
}

func Modal(config ModalConfig) Component {
	return Component{
		Type:     ComponentTypeActionRow,
		Title:    config.Title,
		CustomId: config.CustomId,
		Fields:   config.Fields,
		Handler:  config.Handler,
	}
}

func ModalFields(fields ...TextInputOptions) []ActionRow {
	var rows []ActionRow
	for _, f := range fields {
		rows = append(rows, Row(TextInput(f)))
	}
	return rows
}
