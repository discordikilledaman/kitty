package commands

import (
	"fmt"
	"strings"

	"github.com/acdenisSK/kitty"
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
		context.Say(fmt.Sprintf("```elixir\nInput: %s\nError Output: %s```", code, err))
		return
	}
	context.Say(fmt.Sprintf("```elixir\nInput: %s\nOutput: %s```", code, value.String()))
}
