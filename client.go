package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port               int
	Path               string
	DiscordToken       string
	PublicKey          string
	ApplicationId      string
	ReleaseMode        string
	OnInteractionError func(i *Interaction, err error)
}

type Client struct {
	Config
	Engine   *gin.Engine
	Http     *HttpClient
	commands []ApplicationCommand
	handlers map[string]func(interaction *Interaction) error
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
		c.handlers[fmt.Sprintf("%s:%d", command.Name, command.Type)] = command.Handler
	}
	c.commands = append(c.commands, commands...)
}

func (c *Client) Preload(comps ...Component) {
	for _, comp := range comps {
		if comp.CustomId == "" {
			continue
		}
		c.handlers[comp.CustomId] = comp.Handler
	}
}

func App(config *Config) *Client {
	gin.SetMode(config.ReleaseMode)
	return &Client{
		*config,
		gin.Default(),
		&HttpClient{config},
		[]ApplicationCommand{},
		map[string]func(interaction *Interaction) error{},
	}
}
