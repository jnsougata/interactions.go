package main

import (
	"fmt"
	"strconv"
)

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

func (u *User) Mention() string {
	return "<@" + u.Id + ">"
}

func (u *User) String() string {
	if u.Discriminator == "0" {
		if u.GlobalName != "" {
			return u.GlobalName
		} else {
			return u.Username
		}
	} else {
		return u.Username + "#" + u.Discriminator
	}
}

func (u *User) AvatarAsset() *Asset {
	if u.Avatar != "" {
		return &Asset{
			Hash:     u.Avatar,
			Fragment: fmt.Sprintf("avatars/%s", u.Id),
		}
	}
	asset := &Asset{
		Fragment: "embed/avatars",
	}
	if u.Discriminator == "0" {
		id, _ := strconv.Atoi(u.Id)
		asset.Hash = strconv.Itoa((id >> 22) % 6)
	} else {
		id, _ := strconv.Atoi(u.Discriminator)
		asset.Hash = strconv.Itoa(id % 5)
	}
	return asset
}
