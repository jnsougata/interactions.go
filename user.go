package main

type User struct {
	Id               string `json:"id"`
	Username         string `json:"username"`
	Discriminator    string `json:"discriminator"`
	GlobalName       string `json:"global_name"`
	Avatar           string `json:"avatar"`
	Bot              bool   `json:"bot"`
	System           bool   `json:"system"`
	MfaEnabled       bool   `json:"mfa_enabled"`
	Banner           string `json:"banner"`
	AccentColor      int    `json:"accent_color"`
	Locale           string `json:"locale"`
	Verified         bool   `json:"verified"`
	Email            string `json:"email"`
	Flags            int    `json:"flags"`
	PremiumType      int    `json:"premium_type"`
	PublicFlags      int    `json:"public_flags"`
	AvatarDecoration string `json:"avatar_decoration"`
}
