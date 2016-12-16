package commands

import (
	"fmt"
	"strings"

	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
	"github.com/robertkrimen/otto"
	// only importing it in a library-type-level for otto to autoload underscore
	_ "github.com/robertkrimen/otto/underscore"
)

var vm = otto.New()

// Eval f
type Eval struct{}

// Checks f
func (Eval) Checks() kitty.Checks {
	return kitty.Checks{OwnerOnly: true}
}

// Process f
func (Eval) Process(context kitty.Context) {
	code := strings.Join(context.Args, " ")
	vm.Set("context", context)
	value, err := vm.Run(code)
	if err != nil {
		context.SayEmbed(&discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{
				kitty.Field("Input", fmt.Sprintf("```js\n%s```", code)),
				kitty.Field("Output", fmt.Sprintf("```%s```", value)),
			},
			Color: 0x454ff,
		})
		return
	}
	context.SayEmbed(&discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{
			kitty.Field("Input", fmt.Sprintf("```js\n%s```", code)),
			kitty.Field("Output", fmt.Sprintf("```js\n%s```", value)),
		},
		Color: 0x454ff,
	})
}
