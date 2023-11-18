package main

type PartialMember struct {
	Nick         string   `json:"nick"`
	Avatar       string   `json:"avatar"`
	Roles        []string `json:"roles"`
	JoinedAt     string   `json:"joined_at"`
	PremiumSince string   `json:"premium_since"`

	Flags       int    `json:"flags"`
	Pending     bool   `json:"pending"`
	Permissions string `json:"permissions"`
}

type Memebr struct {
	PartialMember
	User User `json:"user"`
	Deaf bool `json:"deaf"`
	Mute bool `json:"mute"`
}
