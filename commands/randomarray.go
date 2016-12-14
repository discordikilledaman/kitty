package commands

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/acdenisSK/kitty"
)

// RandomNumbers f
type RandomNumbers struct{}

// IsOwnerOnly f
func (RandomNumbers) IsOwnerOnly() bool {
	return false
}

// Help f
func (RandomNumbers) Help() [2]string {
	return [2]string{"Randomly generates an array of numbers", "[num]"}
}

// Process f
func (RandomNumbers) Process(context kitty.Context) {
	rand.Seed(time.Now().UnixNano())
	randomints := []int{}
	a := "50"
	if len(context.Args) != 0 {
		a = strings.Join(context.Args, " ")
	}
	howmuch, err := strconv.Atoi(a)
	if err != nil {
		context.Error(err)
		return
	}
	for i := 0; i < howmuch; i++ {
		randomints = append(randomints, rand.Int())
	}
	context.Say("Result: ", randomints)
}
