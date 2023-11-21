package main

type MessageOptions struct {
	TTS             bool                `json:"tts,omitempty"`
	Content         string              `json:"content,omitempty"`
	Embeds          []Embed             `json:"embeds,omitempty"`
	AllowedMentions any                 `json:"allowed_mentions,omitempty"`
	Flags           int                 `json:"flags,omitempty"`
	Components      []ActionRow         `json:"components,omitempty"`
	Attchments      []PartialAttachment `json:"attachments,omitempty"`
}
