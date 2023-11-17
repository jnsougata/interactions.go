package main

import (
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

type app struct{ AppState }

func (state *app) Run() error {
	gin.SetMode(state.ReleaseMode)
	app := gin.Default()
	app.POST(state.Path, func(c *gin.Context) { handler(c, state) })
	return app.Run()
}

func App(state *AppState) *app {
	return &app{*state}
}
