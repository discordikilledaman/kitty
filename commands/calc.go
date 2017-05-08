package commands

import (
	"strings"

	"github.com/acdenisSK/kitty"
	sy "github.com/mgenware/go-shunting-yard"
)

// Calc f
type Calc struct{}

// IsOwnerOnly f
func (Calc) IsOwnerOnly() bool { return false }

// Help f
func (Calc) Help() [2]string {
	return [2]string{"Calculates the passed expression.", "<expr>"}
}

// Process f
func (Calc) Process(context kitty.Context) {
	tokens, err := sy.Scan(strings.Join(context.Args, " "))
	if err != nil {
		context.Error(err)
		return
	}
	pftokens, err := sy.Parse(tokens)
	if err != nil {
		context.Error(err)
		return
	}
	res, err := sy.Evaluate(pftokens)
	if err != nil {
		context.Error(err)
		return
	}
	e := kitty.NewEmbed("")
	e.Fieldf("The result for your passed expression:", "%d", res)
	context.SayEmbed(e)
}
