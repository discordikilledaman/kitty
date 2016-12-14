package commands

import (
	"fmt"
	"time"

	"github.com/acdenisSK/kitty"
)

// Ping is used for measuring time between the location of the bot and discord's servers.
// Well it's not really *between* discord's servers but you get the idea.
type Ping struct{}

// Checks f
func (Ping) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (Ping) Process(context kitty.Context) {
	start := time.Now()
	message, err := context.Say("pong!")
	if err != nil {
		return
	}
	context.Edit(message.ID, fmt.Sprintf("%s | %f ms", message.Content, time.Since(start).Seconds()*1e3))
}
