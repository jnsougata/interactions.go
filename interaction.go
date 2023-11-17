package main

type Interaction struct {
	Id             string `json:"id"`
	ApplicationId  string `json:"application_id"`
	Type           int    `json:"type"`
	Data           any    `json:"data"`
	GuildId        string `json:"guild_id"`
	Channel        any    `json:"channel"`
	ChannelId      string `json:"channel_id"`
	Member         any    `json:"member"`
	User           any    `json:"user"`
	Token          string `json:"token"`
	Version        int    `json:"version"`
	Message        any    `json:"message"`
	AppPermissions string `json:"app_permissions"`
	Locale         string `json:"locale"`
	GuildLocale    string `json:"guild_locale"`
	Entitlements   any    `json:"entitlements"`
}
