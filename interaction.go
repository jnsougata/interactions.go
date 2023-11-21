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
	Messages    map[string]Message       `json:"messages"`
	Attachments map[string]Attachment    `json:"attachments"`
}

type InteractionData struct {
	Id            string                  `json:"id"`
	Name          string                  `json:"name"`
	Type          ApplicationCommandType  `json:"type"`
	Options       any                     `json:"options"`
	GuildId       string                  `json:"guild_id"`
	TargetId      string                  `json:"target_id"`
	CustomId      string                  `json:"custom_id"`
	ComponentType ComponentType           `json:"component_type"`
	Components    []ActionRow             `json:"components"`
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
	Member         Memebr          `json:"member"`
	User           User            `json:"user"`
	Token          string          `json:"token"`
	Version        int             `json:"version"`
	Message        Message         `json:"message"`
	AppPermissions string          `json:"app_permissions"`
	Locale         string          `json:"locale"`
	GuildLocale    string          `json:"guild_locale"`
	Entitlements   any             `json:"entitlements"`
}

func (i *Interaction) Bind(v any) {
	switch i.Type {
	case InteractionTypeApplicationCommand:
		switch i.Data.Type {
		case ApplicationCommandTypeChatInput:
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
		case ApplicationCommandTypeUser:
			ob, _ := json.Marshal(i.Data.Resolved.Users[i.Data.TargetId])
			_ = json.Unmarshal(ob, v)
		case ApplicationCommandTypeMessage:
			ob, _ := json.Marshal(i.Data.Resolved.Messages[i.Data.TargetId])
			_ = json.Unmarshal(ob, v)
		}
	case InteractionTypeMessageComponent:
		switch i.Data.ComponentType {
		case ComponentTypeTextSelect:
			ob, _ := json.Marshal(i.Data.Values)
			_ = json.Unmarshal(ob, v)
		case ComponentTypeUserSelect:
			var users []User
			for _, id := range i.Data.Values {
				users = append(users, i.Data.Resolved.Users[id])
			}
			ob, _ := json.Marshal(users)
			_ = json.Unmarshal(ob, v)
		case ComponentTypeRoleSelect:
			var roles []Role
			for _, id := range i.Data.Values {
				roles = append(roles, i.Data.Resolved.Roles[id])
			}
			ob, _ := json.Marshal(roles)
			_ = json.Unmarshal(ob, v)
		default:
			// TODO: handle other component types
		}
	case InteractionTypeModalSubmit:
		fields := map[string]any{}
		for _, row := range i.Data.Components {
			for _, component := range row.Components {
				fields[component.CustomId] = component.Value
			}
		}
		ob, _ := json.Marshal(fields)
		_ = json.Unmarshal(ob, v)
	}
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

func (i *Interaction) EditComponentMessage(message MessageOptions) {
	i.App.http.SendInteractionCallback(i, InteractionCallbackTypeUpdateMessage, message)
}

func (i *Interaction) SendModal(modal Modal) {
	i.App.http.SendInteractionCallbackModal(i, InteractionCallbackTypeModal, modal)
}
