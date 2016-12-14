package commands

import (
	"image"

	"image/color"

	"image/png"

	"bytes"

	"fmt"

	"math/rand"

	"time"

	"github.com/acdenisSK/kitty"
)

// RandoImg f
type RandoImg struct{}

// Checks f
func (RandoImg) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (RandoImg) Process(context kitty.Context) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	img := image.NewNRGBA(image.Rect(0, 0, 500, 500))
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			r, g, b := uint8(x*r.Intn(255)), uint8(r.Intn(255)*y), uint8(r.Intn(255)*x/r.Int())
			img.SetNRGBA(x, y, color.NRGBA{r, g, b, uint8(rand.Intn(255))})
		}
	}
	buff := &bytes.Buffer{}
	if err := png.Encode(buff, img); err != nil {
		context.Say(fmt.Sprint("Error encoding image:", err))
		return
	}
	context.Session.ChannelFileSend(context.Channel.ID, "rando.png", buff)
}
