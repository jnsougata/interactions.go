package main

type Component struct {
	Type  ComponentType `json:"type"`
	Label string        `json:"label,omitempty"`
	Style ButtonStyle   `json:"style,omitempty"`
	Emoji struct {
		Name     string `json:"name,omitempty"`
		ID       string `json:"id,omitempty"`
		Animated bool   `json:"animated,omitempty"`
	} `json:"emoji,omitempty"`
	CustomId string                         `json:"custom_id,omitempty"`
	URL      string                         `json:"url,omitempty"`
	Disabled bool                           `json:"disabled,omitempty"`
	Handler  func(interaction *Interaction) `json:"-"`
}

type ButtonOptions struct {
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

type ActionRow struct {
	Type       ComponentType `json:"type"`
	Components []Component   `json:"components"`
}

func Button(options ButtonOptions) Component {
	return Component{
		Type:     ComponentTypeButton,
		Label:    options.Label,
		Style:    options.Style,
		Emoji:    options.Emoji,
		CustomId: options.CustomId,
		URL:      options.URL,
		Disabled: options.Disabled,
		Handler:  options.Handler,
	}
}

func (a *ActionRow) AddButtons(comps ...Component) {
	if len(a.Components) >= 5 {
		return
	}
	a.Components = append(a.Components, comps...)
}
