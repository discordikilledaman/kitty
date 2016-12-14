package kitty

import "github.com/bwmarrin/discordgo"

// Context f
type Context struct {
	Session *discordgo.Session
	State   *discordgo.State
	Message *discordgo.Message
	Author  *discordgo.User
	Channel *discordgo.Channel
	Guild   *discordgo.Guild
	Args    []string
}

// Say ayy lmao
func (context Context) Say(text string) (*discordgo.Message, error) {
	return context.Session.ChannelMessageSend(context.Channel.ID, text)
}

// Edit ayy lmao
func (context Context) Edit(msgID, text string) (*discordgo.Message, error) {
	return context.Session.ChannelMessageEdit(context.Channel.ID, msgID, text)
}

// SayEmbed ayy lmao
func (context Context) SayEmbed(msg *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return context.Session.ChannelMessageSendEmbed(context.Channel.ID, msg)
}
