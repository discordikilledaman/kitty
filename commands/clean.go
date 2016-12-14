package commands

import (
	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
)

// Clean f
type Clean struct{}

// Checks f
func (Clean) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (Clean) Process(context kitty.Context) {
	perms, err := context.Session.UserChannelPermissions(context.State.User.ID, context.Channel.ID)
	if err != nil {
		context.Say(err.Error())
		return
	}
	if perms&discordgo.PermissionManageMessages != 0 {
		messages, err := context.Session.ChannelMessages(context.Channel.ID, 100, "", "")
		if err != nil {
			return
		}
		var _messages []string
		for _, message := range messages {
			if message.Author.ID != context.State.User.ID {
				continue
			}
			_messages = append(_messages, message.ID)
		}
		context.Session.ChannelMessagesBulkDelete(context.Channel.ID, _messages)
	}
}
