package commands

import (
	"fmt"
	"time"

	"github.com/acdenisSK/kitty"
)

// Ping is used for measuring time between the location of the bot and discord's servers.
// Well it's not really *between* discord's servers but you get the idea.
type Ping struct{}

// IsOwnerOnly f
func (Ping) IsOwnerOnly() bool {
	return false
}

// Help f
func (Ping) Help() [2]string {
	return [2]string{"Measures how fast the bot can send to discord's api", ""}
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
