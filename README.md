# interactions.go
Discord HTTP Interaction Handler for Serverless Applications

## Work In Progress

### Quick Start
```go
package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func main() {
	app := App(
		&AppState{
			Port:          8080,
			Path:          "...",
			DiscordToken:  "...",
			PublicKey:     "...",
			ApplicationId: "...",
			ReleaseMode:   gin.ReleaseMode,
		})
	app.AddCommands(echo)
	if err := app.Run(); err != nil {
		log.Fatal(fmt.Errorf("failed to run app: %w", err))
	}
}

var echo = ApplicationCommand{
	Name:        "echo",
	Type:        ApplicationCommandTypeChatInput,
	Description: "Echoes a message",
	Options: []Option{
		{
			Name:        "message",
			Description: "The message to echo",
			Type:        ApplicationCommandOptionTypeString,
			Required:    true,
		},
	},
	Handler: func(interaction *Interaction) {
		var options struct {
			Message string `json:"message"`
		}
		interaction.Bind(&options)
		embed := Embed{
			Description: fmt.Sprintf("```\n%s\n```", options.Message),
		}
		interaction.Response(MessageOptions{Embeds: []Embed{embed}})
	},
}
```