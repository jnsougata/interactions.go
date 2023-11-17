package main

import (
	"crypto/ed25519"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context, state *app) {
	signature := c.GetHeader("X-Signature-Ed25519")
	timestamp := c.GetHeader("X-Signature-Timestamp")
	if signature == "" || timestamp == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	message := []byte(timestamp)
	message = append(message, body...)
	ok := ed25519.Verify([]byte(state.PublicKey), message, []byte(signature))
	if !ok {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var interaction Interaction
	_ = json.Unmarshal(body, &interaction)

	switch interaction.Type {
	case int(InteractionTypePing):
		c.JSON(200, gin.H{
			"type": 1,
		})
	default:
		c.JSON(200, gin.H{
			"type": 1,
		})
	}

}
