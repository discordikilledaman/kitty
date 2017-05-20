package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/acdenisSK/kitty"
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
	input := strings.Join(context.Args, " ")
	float := strings.Contains(input, "-f")
	if float {
		input = input[:strings.IndexAny(input, "-f")]
	}

	var cmd *exec.Cmd

	if float {
		cmd = exec.Command("shunt_yard", "-e", fmt.Sprintf("\"%s\"", input), "-f")
	} else {
		cmd = exec.Command("shunt_yard", "-e", fmt.Sprintf("\"%s\"", input))
	}
	res, err := cmd.Output()
	if err != nil {
		context.Error(err)
		return
	}
	e := kitty.NewEmbed("")
	e.Fieldf("The result for your passed expression:", "%s", string(res))
	context.SayEmbed(e)
}
