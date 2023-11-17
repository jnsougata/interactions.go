package main

type ApplicationCommand struct {
	Name        string                         `json:"name"`
	Type        int                            `json:"type"`
	Description string                         `json:"description,omitempty"`
	Options     []any                          `json:"options,omitempty"`
	Handler     func(interaction *Interaction) `json:"-"`
}
