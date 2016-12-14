package commands

import (
	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
)

// Clean f
type Clean struct{}

// IsOwnerOnly f
func (Clean) IsOwnerOnly() bool {
	return false
}

// Help f
func (Clean) Help() [2]string {
	return [2]string{"Cleans the bot's messages from the current channel", ""}
}

// Process f
func (Clean) Process(context kitty.Context) {
	perms, err := context.ChannelPermissions()
	if err != nil {
		context.Error(err)
		return
	}
	if perms&discordgo.PermissionManageMessages == 0 {
		return
	}
	messages, err := context.GetMessages(100)
	if err != nil {
		return
	}
	for index, message := range messages {
		if message.Author.ID != context.State.User.ID {
			continue
		}
		messages = append(messages[:index], messages[index+1:]...) // delete it
	}
	context.BulkDelete(messages)
}
