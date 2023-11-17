package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type InteractionData struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	Resolved      any    `json:"resolved"`
	Options       any    `json:"options"`
	GuildId       string `json:"guild_id"`
	TargetId      string `json:"target_id"`
	CustomId      string `json:"custom_id"`
	ComponentType int    `json:"component_type"`
	Values        any    `json:"values"`
	Components    any    `json:"components"`
}

type Interaction struct {
	App            *app
	Context        *gin.Context
	Id             string          `json:"id"`
	ApplicationId  string          `json:"application_id"`
	Type           InteractionType `json:"type"`
	Data           InteractionData `json:"data"`
	GuildId        string          `json:"guild_id"`
	Channel        any             `json:"channel"`
	ChannelId      string          `json:"channel_id"`
	Member         any             `json:"member"`
	User           any             `json:"user"`
	Token          string          `json:"token"`
	Version        int             `json:"version"`
	Message        any             `json:"message"`
	AppPermissions string          `json:"app_permissions"`
	Locale         string          `json:"locale"`
	GuildLocale    string          `json:"guild_locale"`
	Entitlements   any             `json:"entitlements"`
}

func (i *Interaction) Bind(v any) {
	options := map[string]any{}
	for _, option := range i.Data.Options.([]interface{}) {
		om := option.(map[string]any)
		options[om["name"].(string)] = om["value"]
	}
	ob, _ := json.Marshal(options)
	_ = json.Unmarshal(ob, v)
}

func (i *Interaction) Respond(message MessageOptions) {
	i.App.http.SendInteractionCallback(i, InteractionCallbackTypeChannelMessageWithSource, message)
}

func (i *Interaction) FollowUp(message MessageOptions) {
	i.App.http.SendInteractionFollowup(i, message)
}

func (i *Interaction) Defer(ephemral ...bool) {
	var payload MessageOptions
	var kind InteractionCallbackType
	if i.Type == InteractionTypeApplicationCommand {
		kind = InteractionCallbackTypeDeferredChannelMessageWithSource
		if len(ephemral) > 0 && ephemral[0] {
			payload.Flags = 64
		}
	} else {
		kind = InteractionCallbackTypeDeferredMessageUpdate
	}

	i.App.http.SendInteractionCallback(i, kind, payload)
}
