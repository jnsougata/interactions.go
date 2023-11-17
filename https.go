package main

import (
	"fmt"
	"io"
	"net/http"
)

var DISCORD_API_VERSION = "10"
var DISCORD_API_BASE_URL = "https://discord.com/api/v" + DISCORD_API_VERSION

type HttpClient struct {
	state *AppState
}

func (c *HttpClient) Request(
	method, path string,
	authorize bool,
	body io.Reader,
	kwargs map[string]interface{},
) (*http.Response, error) {
	url := DISCORD_API_BASE_URL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if authorize {
		req.Header.Set("Authorization", "Bot "+c.state.DiscordToken)
	}
	if kwargs != nil {
		if reason, ok := kwargs["reason"]; ok {
			req.Header.Set("X-Audit-Log-Reason", reason.(string))
		}
	}
	return http.DefaultClient.Do(req)
}

func (c *HttpClient) Sync(commands []ApplicationCommand) (*http.Response, error) {
	return c.Request(
		http.MethodPut, 
		fmt.Sprintf("/applications/%s/commands", c.state.ApplicationId), 
		true, ReaderFromAny(commands), 	nil)
}

func NewHttpClient(state *AppState) *HttpClient {
	return &HttpClient{state}
}
