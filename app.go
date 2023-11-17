package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AppState struct {
	Port          int
	Path          string
	DiscordToken  string
	PublicKey     string
	ApplicationId string
	ReleaseMode   string
}

type app struct {
	AppState
	Engine *gin.Engine
	http   *HttpClient
}

func (state *app) Run() error {
	state.Engine.POST(state.Path, func(c *gin.Context) { handler(c, state) })
	var PORT = ":8080"
	if state.Port != 0 {
		PORT = fmt.Sprintf(":%d", state.Port)
	}
	return state.Engine.Run(PORT)
}

func App(state *AppState) *app {
	gin.SetMode(state.ReleaseMode)
	return &app{*state, gin.Default(), NewHttpClient(state)}
}
