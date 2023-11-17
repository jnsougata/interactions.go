package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
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
	_ = json.Unmarshal(body, &interaction)

	switch interaction.Type {
	case InteractionTypePing:
		c.JSON(200, gin.H{"type": 1})
	default:
		c.JSON(200, gin.H{"type": 1})
	}

}
