package commands

import (
	"strings"

	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
)

// Help f
type Help struct {
	Commands map[string]kitty.Command
}

// Checks f
func (Help) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (h Help) Process(context kitty.Context) {
	var commands []string
	for command := range h.Commands {
		commands = append(commands, command)
	}
	context.SayEmbed(&discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{
			kitty.Field("Currently available commands", strings.Join(commands, ", ")),
		},
		Color: 0x6545ff,
	})

}
