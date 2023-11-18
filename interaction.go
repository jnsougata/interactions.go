package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type InteractionDataResolved struct {
	Users       map[string]User          `json:"users"`
	Members     map[string]PartialMember `json:"members"`
	Roles       map[string]Role          `json:"roles"`
	Channels    map[string]interface{}   `json:"channels"`
	Messages    map[string]interface{}   `json:"messages"`
	Attachments map[string]interface{}   `json:"attachments"`
}

type InteractionData struct {
	Id            string                  `json:"id"`
	Name          string                  `json:"name"`
	Type          int                     `json:"type"`
	Options       any                     `json:"options"`
	GuildId       string                  `json:"guild_id"`
	TargetId      string                  `json:"target_id"`
	CustomId      string                  `json:"custom_id"`
	ComponentType int                     `json:"component_type"`
	Values        []string                `json:"values"`
	Resolved      InteractionDataResolved `json:"resolved"`
}

type Interaction struct {
	App            *app
	Context        *gin.Context
	Id             string          `json:"id"`
	ApplicationId  string          `json:"application_id"`
	Type           InteractionType `json:"type"`
	Data           InteractionData `json:"data"`
	GuildId        string          `json:"guild_id"`
	Channel        any             `json:"channel"`
	ChannelId      string          `json:"channel_id"`
	Member         any             `json:"member"`
	User           any             `json:"user"`
	Token          string          `json:"token"`
	Version        int             `json:"version"`
	Message        Message         `json:"message"`
	AppPermissions string          `json:"app_permissions"`
	Locale         string          `json:"locale"`
	GuildLocale    string          `json:"guild_locale"`
	Entitlements   any             `json:"entitlements"`
}

func (i *Interaction) Bind(v any) {
	options := map[string]any{}
	for _, option := range i.Data.Options.([]interface{}) {
		var o Option
		b, _ := json.Marshal(option)
		_ = json.Unmarshal(b, &o)
		switch o.Type {
		case ApplicationCommandOptionTypeSubCommand:
		case ApplicationCommandOptionTypeSubCommandGroup:
		case ApplicationCommandOptionTypeString:
			options[o.Name] = o.Value.(string)
		case ApplicationCommandOptionTypeInteger:
			options[o.Name] = int(o.Value.(float64))
		case ApplicationCommandOptionTypeBoolean:
			options[o.Name] = o.Value.(bool)
		case ApplicationCommandOptionTypeUser:
			options[o.Name] = i.Data.Resolved.Users[o.Value.(string)]
		case ApplicationCommandOptionTypeChannel:
			options[o.Name] = i.Data.Resolved.Channels[o.Value.(string)]
		case ApplicationCommandOptionTypeRole:
			options[o.Name] = i.Data.Resolved.Roles[o.Value.(string)]
		case ApplicationCommandOptionTypeMentionable:
			user, ok := i.Data.Resolved.Users[o.Value.(string)]
			if ok {
				options[o.Name] = user
			} else {
				options[o.Name] = i.Data.Resolved.Roles[o.Value.(string)]
			}
		case ApplicationCommandOptionTypeNumber:
			options[o.Name] = o.Value.(float64)
		case ApplicationCommandOptionTypeAttachment:
			options[o.Name] = i.Data.Resolved.Attachments[o.Value.(string)]
		}
	}
	ob, _ := json.Marshal(options)
	_ = json.Unmarshal(ob, v)
}

func (i *Interaction) Response(message MessageOptions) {
	i.App.http.SendInteractionCallback(i, InteractionCallbackTypeChannelMessageWithSource, message)
}

func (i *Interaction) FollowUp(message MessageOptions) {
	i.App.http.SendInteractionFollowup(i, message)
}

func (i *Interaction) Defer(ephemral ...bool) {
	var payload MessageOptions
	var kind InteractionCallbackType
	if i.Type == InteractionTypeApplicationCommand {
		kind = InteractionCallbackTypeDeferredChannelMessageWithSource
		if len(ephemral) > 0 && ephemral[0] {
			payload.Flags = 64
		}
	} else {
		kind = InteractionCallbackTypeDeferredUpdateMessage
	}

	i.App.http.SendInteractionCallback(i, kind, payload)
}

func (i *Interaction) GetOriginalResponse() Message {
	resp, _ := i.App.http.GetOriginalInteractionResponse(i)
	var message Message
	_ = json.NewDecoder(resp.Body).Decode(&message)
	return message
}

func (i *Interaction) EditOriginalResponse(message MessageOptions) {
	i.App.http.EditOriginalInteractionResponse(i, message)
}

func (i *Interaction) DeleteOriginalResponse() {
	i.App.http.DeleteOriginalInteractionResponse(i)
}

func (i *Interaction) UpdateComponentMessage(message MessageOptions) {
	i.App.http.SendInteractionCallback(i, InteractionCallbackTypeUpdateMessage, message)
}
