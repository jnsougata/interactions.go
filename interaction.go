package main

import "github.com/gin-gonic/gin"

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

func (i *Interaction) Send(v any) {
	i.App.http.SendInteractionCallback(i.Id, i.Token, v)
}