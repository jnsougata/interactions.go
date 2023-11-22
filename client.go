package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var globalHandlerMap = make(map[string]func(interaction *Interaction))

type AppState struct {
	Port          int
	Path          string
	DiscordToken  string
	PublicKey     string
	ApplicationId string
	ReleaseMode   string
	commands      []ApplicationCommand
}

type Client struct {
	AppState
	Engine *gin.Engine
	Http   *HttpClient
}

func (c *Client) Run() error {
	c.Engine.POST(c.Path, func(ctx *gin.Context) { handler(ctx, c) })
	var PORT = ":8080"
	if c.Port != 0 {
		PORT = fmt.Sprintf(":%d", c.Port)
	}
	return c.Engine.Run(PORT)
}

func (c *Client) Sync() (*http.Response, error) {
	return c.Http.sync(c.commands)
}

func (c *Client) AddCommands(commands ...ApplicationCommand) {
	for _, command := range commands {
		globalHandlerMap[fmt.Sprintf("%s:%d", command.Name, command.Type)] = command.Handler
	}
	c.commands = append(c.commands, commands...)
}

func (c *Client) PreloadComponents(comps ...Component) {
	for _, comp := range comps {
		if comp.CustomId == "" {
			continue
		}
		globalHandlerMap[fmt.Sprintf("%s:%d", comp.CustomId, comp.Type)] = comp.Handler
	}
}

func (c *Client) PreloadModal(m Modal) {
	globalHandlerMap[m.CustomId] = m.Handler
}

func App(state *AppState) *Client {
	gin.SetMode(state.ReleaseMode)
	return &Client{*state, gin.Default(), &HttpClient{state}}
}
