package main

type EmbedAuthor struct {
	Name     string `json:"name"`
	URL      string `json:"url,omitempty"`
	IconURL  string `json:"icon_url,omitempty"`
	ProxyURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type EmbedThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type Embed struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	Color       int            `json:"color,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
	Author      EmbedAuthor    `json:"author,omitempty"`
	Thumbnail   EmbedThumbnail `json:"thumbnail,omitempty"`
	Image       EmbedImage     `json:"image,omitempty"`
	Footer      EmbedFooter    `json:"footer,omitempty"`
	Fields      []EmbedField   `json:"fields,omitempty"`
}

func (e *Embed) AppendField(name, value string, inline bool) {
	e.Fields = append(e.Fields, EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
}
