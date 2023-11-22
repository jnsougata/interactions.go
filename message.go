package main

import "net/http"

type MessageInteraction struct {
	Id     string `json:"id"`
	Type   int    `json:"type"`
	Name   string `json:"name"`
	User   User   `json:"user"`
	Memebr Memebr `json:"member"`
}

type Message struct {
	Id              string `json:"id"`
	ChannelId       string `json:"channel_id"`
	Author          User   `json:"author"`
	Content         string `json:"content"`
	Timestamp       string `json:"timestamp"`
	EditedTimestamp string `json:"edited_timestamp"`
	TTS             bool   `json:"tts"`
	MentionEveryone bool   `json:"mention_everyone"`
	Mentions        []User `json:"mentions"`
	MentionRoles    []Role `json:"mention_roles"`
	// MentionChannels []Channel `json:"mention_channels"`
	// Attachments     []Attachment `json:"attachments"`
	Embeds []Embed `json:"embeds"`
	// Reactions       []Reaction `json:"reactions"`
	Nonce         string `json:"nonce"`
	Pinned        bool   `json:"pinned"`
	WebhookId     string `json:"webhook_id"`
	Type          int    `json:"type"`
	Activity      string `json:"activity"`
	Application   any    `json:"application"`
	ApplicationId string `json:"application_id"`
	// MessageReference  any         `json:"message_reference"`
	Flags             int                `json:"flags"`
	ReferencedMessage *Message           `json:"referenced_message"`
	Interaction       MessageInteraction `json:"interaction"`
	// Thread            any         `json:"thread"`
	Components []Component `json:"components"`
	// StickerItems      []Sticker   `json:"sticker_items"`
	// Stickers          []Sticker   `json:"stickers"`
	Position int `json:"position"`
	// RoleSubscriptionData any `json:"role_subscription_data"`
	Resolved InteractionDataResolved `json:"resolved"`
}

func (m *Message) Delete(c *Client) (*http.Response, error) {
	return c.Http.DeleteMessage(m.Id, m.ChannelId)
}
