# interactions.go
Discord HTTP Interaction Handler for Serverless Applications

## WIP

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
	app.AddCommands(randint)
	if err := app.Run(); err != nil {
		log.Fatal(fmt.Errorf("failed to run app: %w", err))
	}
}

var randint = ApplicationCommand{
	Name:        "randint",
	Type:        ApplicationCommandTypeChatInput,
	Description: "Get a random integer",
	Options: []Option{
		{
			Name:        "min",
			Description: "The minimum value",
			Type:        ApplicationCommandOptionTypeInteger,
			Required:    true,
		},
		{
			Name:        "max",
			Description: "The maximum value",
			Type:        ApplicationCommandOptionTypeInteger,
			Required:    true,
		},
	},
	Handler: func(interaction *Interaction) {
		var options struct {
			Max int `json:"max"`
			Min int `json:"min"`
		}
		interaction.Bind(&options)
		if options.Max < options.Min {
			interaction.SendMessage(SendingPayload{
				Content: "Max must be greater than min",
			})
			return
		}
		embed := Embed{
			Description: fmt.Sprintf("```\n%d\n```", options.Min+rand.Intn(options.Max-options.Min+1)),
		}
		interaction.SendMessage(SendingPayload{
			Embeds: []Embed{embed},
		})
	},
}

```