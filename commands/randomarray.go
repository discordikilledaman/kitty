package commands

import (
	"fmt"
	"strconv"

	"math/rand"
	"time"

	"github.com/acdenisSK/kitty"
)

// RandomArray f
type RandomArray struct{}

// Checks f
func (RandomArray) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (RandomArray) Process(context kitty.Context) {
	rand.Seed(time.Now().UnixNano())
	randomints := []int{}
	var a string
	if len(context.Args) == 0 {
		a = "50"
	} else {
		a = context.Args[0]
	}
	howmuch, err := strconv.Atoi(a)
	if err != nil {
		context.Say(fmt.Sprint("An error occured when converting your number: ", err))
		return
	}
	for i := 0; i < howmuch; i++ {
		randomints = append(randomints, rand.Int())
	}
	context.Say(fmt.Sprint("Result: ", randomints))
}
