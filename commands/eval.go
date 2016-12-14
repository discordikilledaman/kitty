package commands

import (
	"strings"

	"github.com/acdenisSK/kitty"
	"github.com/robertkrimen/otto"
	// only importing it in a library-type-level for otto to autoload underscore, also god damn you linters
	_ "github.com/robertkrimen/otto/underscore"
)

var vm = otto.New()

// Eval f
type Eval struct{}

// IsOwnerOnly f
func (Eval) IsOwnerOnly() bool {
	return true
}

// Help f
func (Eval) Help() [2]string {
	return [2]string{"Executes javascript code", "<code>"}
}

// Process f
func (Eval) Process(context kitty.Context) {
	code := strings.Join(context.Args, " ")
	vm.Set("context", context)
	value, err := vm.Run(code)
	if err != nil {
		embed := kitty.NewEmbed("")
		embed.Fieldf("Input", "```js\n%s```", code)
		embed.Fieldf("Output", "```%s```", value)
		context.SayEmbed(embed)
		return
	}
	embed := kitty.NewEmbed("")
	embed.Fieldf("Input", "```js\n%s```", code)
	embed.Fieldf("Output", "```js\n%s```", value)
	context.SayEmbed(embed)
}
