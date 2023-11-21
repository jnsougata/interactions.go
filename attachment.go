package main

import "strconv"

type PartialAttachment struct {
	Id          string `json:"id"`
	Filename    string `json:"filename"`
	Description string `json:"description,omitempty"`
	Content     []byte `json:"-"`
}

type Attachment struct {
	PartialAttachment
	ContentType  string `json:"content_type,omitempty"`
	Size         int64  `json:"size,omitempty"`
	URL          string `json:"url,omitempty"`
	ProxyURL     string `json:"proxy_url,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
	Ephemeral    bool   `json:"ephemeral,omitempty"`
	DurationSecs int    `json:"duration_secs,omitempty"`
	Waveform     string `json:"waveform,omitempty"`
	Flags        int    `json:"flags,omitempty"`
}

func Attachments(atts ...PartialAttachment) []PartialAttachment {
	for i := range atts {
		atts[i].Id = strconv.Itoa(i)
	}
	return atts
}
