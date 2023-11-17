package main

type SendingPayload struct {
	TTS             bool   `json:"tts,omitempty"`
	Content         string `json:"content,omitempty"`
	Embeds          any    `json:"embeds,omitempty"`
	AllowedMentions any    `json:"allowed_mentions,omitempty"`
	Flags           int    `json:"flags,omitempty"`
	Components      any    `json:"components,omitempty"`
	Attchments      []File `json:"attachments,omitempty"`
}
