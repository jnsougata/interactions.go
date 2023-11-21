package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context, state *app) {
	signatureHex := c.GetHeader("X-Signature-Ed25519")
	timestamp := c.GetHeader("X-Signature-Timestamp")
	if signatureHex == "" || timestamp == "" {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	message := []byte(timestamp)
	message = append(message, body...)
	signatureBytes, _ := hex.DecodeString(signatureHex)
	publicKeyBytes, _ := hex.DecodeString(state.PublicKey)
	ok := ed25519.Verify(publicKeyBytes, message, signatureBytes)
	if !ok {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	var interaction Interaction
	interaction.App = state
	interaction.Context = c
	_ = json.Unmarshal(body, &interaction)
	switch interaction.Type {
	case InteractionTypePing:
		c.JSON(200, gin.H{"type": 1})
	case InteractionTypeApplicationCommand:
		key := fmt.Sprintf("%s:%d", interaction.Data.Name, interaction.Data.Type)
		globalHandlerMap[key](&interaction)
	case InteractionTypeMessageComponent:
		key := fmt.Sprintf("%s:%d", interaction.Data.CustomId, interaction.Data.ComponentType)
		globalHandlerMap[key](&interaction)
	case InteractionTypeModalSubmit:
		globalHandlerMap[interaction.Data.CustomId](&interaction)
	default:
		c.JSON(200, gin.H{"type": 1})
	}
}
