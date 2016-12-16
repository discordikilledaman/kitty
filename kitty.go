package kitty

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Kitty the main struct of the bot
type Kitty struct {
	Logger   *Logger
	Commands map[string]Command
	Config   Config
	Session  *discordgo.Session
}

// New creates a new kitty instance
func New(logger *Logger, commands map[string]Command, config Config, session *discordgo.Session) *Kitty {
	return &Kitty{logger, commands, config, session}
}

// Setup ... setups
func (k *Kitty) Setup() {
	k.Session.AddHandler(func(_ *discordgo.Session, event interface{}) {
		switch data := event.(type) {
		case *discordgo.Ready:
			k.Logger.Println("Running!")
		case *discordgo.MessageCreate:
			if data.Author.ID == k.Session.State.User.ID || data.Author.Bot || !strings.HasPrefix(data.Content, k.Config.Required.Prefix) {
				return
			}
			args := strings.Split(data.Content[len(k.Config.Required.Prefix):], " ")
			command, args := args[0], args[1:]
			cmd, ok := k.Commands[command]
			if !ok {
				k.Logger.Printf("Author %s tried to issue command %s but it doesn't exist\n", data.Author.Username, command)
				return
			}
			channel, err := k.Session.State.Channel(data.ChannelID)
			if err != nil {
				k.Logger.Println(err)
				return
			}
			guild, err := k.Session.State.Guild(channel.GuildID)
			if err != nil {
				k.Logger.Println(err)
				return
			}
			if cmd.Checks().OwnerOnly && data.Author.ID != k.Config.Required.OwnerID {
				return
			}
			k.Logger.Printf("Author %s issued command %s", data.Author.Username, command)
			go cmd.Process(Context{
				Session: k.Session,
				State:   k.Session.State,
				Message: data.Message,
				Channel: channel,
				Author:  data.Author,
				Guild:   guild,
				Args:    args,
			})
		case *discordgo.GuildMemberAdd:
			if data.GuildID != "198101180180594688" {
				return
			}
			k.Session.ChannelMessageSend(data.GuildID, fmt.Sprintf("**%s** is a heck", data.User.Username))
		case *discordgo.GuildBanRemove:
			if data.GuildID != "198101180180594688" {
				return
			}
			k.Session.ChannelMessageSend(data.GuildID, fmt.Sprintf("no, **%s** should stay hecked!!", data.User.Username))
		case *discordgo.GuildMemberRemove:
			if data.GuildID != "198101180180594688" {
				return
			}
			k.Session.ChannelMessageSend(data.GuildID, fmt.Sprintf("**%s** got hecked", data.User.Username))
		}
	})
	if err := k.Session.Open(); err != nil {
		k.Logger.Panicln("error opening websocket:", err)
	}
}
