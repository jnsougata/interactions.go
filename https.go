package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

var DISCORD_API_VERSION = "10"
var DISCORD_API_BASE_URL = "https://discord.com/api/v" + DISCORD_API_VERSION

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func MultipartForm(data any, files []PartialAttachment) ([]byte, string) {
	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
	payload, _ := json.Marshal(data)
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", `form-data; name="payload_json"`)
	header.Set("Content-Type", `application/json`)
	field, _ := writer.CreatePart(header)
	_, _ = field.Write(payload)
	for _, f := range files {
		fw, _ := writer.CreateFormFile(fmt.Sprintf(`files[%s]`, f.Id), quoteEscaper.Replace(f.Filename))
		_, _ = fw.Write(f.Content)
	}
	_ = writer.Close()
	return buff.Bytes(), writer.Boundary()
}

type RequestOptions struct {
	Method    string
	Path      string
	Authorize bool
	Body      io.Reader
	Boundary  string
	Kwargs    map[string]interface{}
}

type HttpClient struct {
	config *Config
}

func (c *HttpClient) Request(o RequestOptions) (*http.Response, error) {
	url := DISCORD_API_BASE_URL + o.Path
	req, err := http.NewRequest(o.Method, url, o.Body)
	if err != nil {
		return nil, err
	}
	if o.Boundary != "" {
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+o.Boundary)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
	if o.Authorize {
		req.Header.Set("Authorization", "Bot "+c.config.DiscordToken)
	}
	if o.Kwargs != nil {
		if reason, ok := o.Kwargs["reason"]; ok {
			req.Header.Set("X-Audit-Log-Reason", reason.(string))
		}
	}
	return http.DefaultClient.Do(req)
}

func (c *HttpClient) sync(commands []ApplicationCommand) (*http.Response, error) {
	return c.Request(
		RequestOptions{
			Method:    http.MethodPut,
			Path:      fmt.Sprintf("/applications/%s/commands", c.config.ApplicationId),
			Authorize: true,
			Body:      ReaderFromMap(commands),
		})
}

func (c *HttpClient) DeleteMessage(messageId, channelId string) (*http.Response, error) {
	return c.Request(RequestOptions{
		Method:    http.MethodDelete,
		Path:      fmt.Sprintf("/channels/%s/messages/%s", channelId, messageId),
		Authorize: true,
	})
}

func (c *HttpClient) SendInteractionCallback(
	interaction *Interaction,
	kind InteractionCallbackType,
	payload MessageOptions,
) (*http.Response, error) {
	f := map[string]interface{}{"type": int(kind), "data": payload}
	data, boundary := MultipartForm(f, payload.Attchments)
	return c.Request(RequestOptions{
		Method:    http.MethodPost,
		Path:      fmt.Sprintf("/interactions/%s/%s/callback", interaction.Id, interaction.Token),
		Authorize: false,
		Body:      bytes.NewReader(data),
		Boundary:  boundary,
	})
}

func (c *HttpClient) SendInteractionCallbackModal(
	interaction *Interaction,
	kind InteractionCallbackType,
	modal Component,
) (*http.Response, error) {
	payload := map[string]interface{}{"type": int(kind), "data": modal}
	return c.Request(RequestOptions{
		Method:    http.MethodPost,
		Path:      fmt.Sprintf("/interactions/%s/%s/callback", interaction.Id, interaction.Token),
		Authorize: false,
		Body:      ReaderFromMap(payload),
	})
}

func (c *HttpClient) SendInteractionFollowup(interaction *Interaction, payload MessageOptions) (*http.Response, error) {
	data, bounday := MultipartForm(payload, payload.Attchments)
	return c.Request(RequestOptions{
		Method:    http.MethodPost,
		Path:      fmt.Sprintf("/webhooks/%s/%s", c.config.ApplicationId, interaction.Token),
		Authorize: false,
		Body:      bytes.NewReader(data),
		Boundary:  bounday,
	})
}

func (c *HttpClient) GetOriginalInteractionResponse(interaction *Interaction) (*http.Response, error) {
	return c.Request(RequestOptions{
		Method:    http.MethodGet,
		Path:      fmt.Sprintf("/webhooks/%s/%s/messages/@original", c.config.ApplicationId, interaction.Token),
		Authorize: false,
	})
}

func (c *HttpClient) EditOriginalInteractionResponse(interaction *Interaction, payload MessageOptions) (*http.Response, error) {
	data, bounday := MultipartForm(payload, payload.Attchments)
	return c.Request(RequestOptions{
		Method:    http.MethodPatch,
		Path:      fmt.Sprintf("/webhooks/%s/%s/messages/@original", c.config.ApplicationId, interaction.Token),
		Authorize: false,
		Body:      bytes.NewReader(data),
		Boundary:  bounday,
	})
}

func (c *HttpClient) DeleteOriginalInteractionResponse(interaction *Interaction) (*http.Response, error) {
	return c.Request(RequestOptions{
		Method:    http.MethodDelete,
		Path:      fmt.Sprintf("/webhooks/%s/%s/messages/@original", c.config.ApplicationId, interaction.Token),
		Authorize: false,
	})
}

func (c *HttpClient) CreateMessage(channelId string, payload MessageOptions) (*http.Response, error) {
	data, bounday := MultipartForm(payload, payload.Attchments)
	return c.Request(RequestOptions{
		Method:    http.MethodPost,
		Path:      fmt.Sprintf("/channels/%s/messages", channelId),
		Authorize: true,
		Body:      bytes.NewReader(data),
		Boundary:  bounday,
	})
}

func (c *HttpClient) CreateDM(userId string, msg MessageOptions) (*http.Response, error) {
	payload := map[string]interface{}{"recipient_id": userId}
	resp, _ := c.Request(RequestOptions{
		Method:    http.MethodPost,
		Path:      "/users/@me/channels",
		Authorize: true,
		Body:      ReaderFromMap(payload),
	})
	var channel struct {
		Id string `json:"id"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&channel)
	return c.CreateMessage(channel.Id, msg)
}
