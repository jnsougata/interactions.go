package main

type RoleTag struct {
	BotId                 string      `json:"bot_id"`
	IntegrationId         string      `json:"integration_id"`
	PremiumSubscriber     interface{} `json:"premium_subscriber"`
	SubscriptionListingId string      `json:"subscription_listing_id"`
	AvailableForPurchase  interface{} `json:"available_for_purchase"`
	GuildConnections      interface{} `json:"guild_connections"`
}

type Role struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Color        int       `json:"color"`
	Hoist        bool      `json:"hoist"`
	Icon         string    `json:"icon"`
	UnicodeEmoji string    `json:"unicode_emoji"`
	Position     int       `json:"position"`
	Permissions  string    `json:"permissions"`
	Managed      bool      `json:"managed"`
	Mentionable  bool      `json:"mentionable"`
	Tags         []RoleTag `json:"tags"`
	Flags        int       `json:"flags"`
}
