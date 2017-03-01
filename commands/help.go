package commands

import (
	"sort"
	"strings"

	"github.com/acdenisSK/kitty"
)

// Help f
type Help struct {
	Commands map[string]kitty.Command
}

// IsOwnerOnly f
func (Help) IsOwnerOnly() bool {
	return false
}

// Help the irony
func (Help) Help() [2]string {
	return [2]string{`Provides information about what current commands are available and what the command is for and such.
Usages use <name> to denote "required" and [name] for "optional".`, "[command]"}
}

// Process f
func (h Help) Process(context kitty.Context) {
	if len(context.Args) == 0 {
		var commands []string
		for command := range h.Commands {
			commands = append(commands, command)
		}
		sort.Strings(commands)
		embed := kitty.NewEmbed("")
		embed.Field("Currently available commands", strings.Join(commands, ", "))
		context.SayEmbed(embed)
		return
	}
	name := strings.Join(context.Args, " ")
	if command, ok := h.Commands[name]; ok {
		help := command.Help()
		description, usage := help[0], help[1]
		embed := kitty.NewEmbed("")
		embed.Field("Description", description)
		embed.Fieldf("Usage", "%s%s %s", kitty.DefaultConfig.Prefix, name, usage)
		context.SayEmbed(embed)
	}
}
