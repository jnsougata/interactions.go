package main

import (
	"io"
	"net/http"
)

var DISCORD_API_VERSION = "10"
var DISCORD_API_BASE_URL = "https://discord.com/api/v" + DISCORD_API_VERSION

type HttpClient struct {
	state *AppState
}

func (h *HttpClient) Request(
	method, path string,
	auth bool,
	body io.Reader,
	kwargs map[string]interface{},
) (*http.Response, error) {
	url := DISCORD_API_BASE_URL + path
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if auth {
		req.Header.Set("Authorization", "Bot "+h.state.DiscordToken)
	}
	reason, ok := kwargs["reason"]
	if ok {
		req.Header.Set("X-Audit-Log-Reason", reason.(string))
	}
	return http.DefaultClient.Do(req)
}

func NewHttpClient(state *AppState) *HttpClient {
	return &HttpClient{state}
}
