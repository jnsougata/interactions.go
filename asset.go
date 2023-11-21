package main

import (
	"fmt"
	"strings"
)

var BaseAssetURL = "https://cdn.discordapp.com"

type Asset struct {
	Hash     string
	Fragment string
}

func (a *Asset) Dynamic() bool {
	return strings.HasPrefix(a.Hash, "a_")
}

func (a *Asset) Default() bool {
	return len(a.Hash) == 1
}

func (a *Asset) URL() string {
	if !a.Dynamic() {
		return fmt.Sprintf("%s/%s/%s.png?size=1024", BaseAssetURL, a.Fragment, a.Hash)
	}
	return fmt.Sprintf("%s/%s/%s.gif?size=1024", BaseAssetURL, a.Fragment, a.Hash)
}

func (a *Asset) CustomURL(size int, format string) string {
	if size <= 0 {
		size = 1024
	}
	return fmt.Sprintf("%s/%s/%s.%s?size=%d", BaseAssetURL, a.Fragment, a.Hash, format, size)
}
