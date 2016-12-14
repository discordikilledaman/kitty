package commands

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/acdenisSK/kitty"
)

var buffer = make([][]byte, 0)
var isplaying bool

// Play f
type Play struct{}

// IsOwnerOnly f
func (Play) IsOwnerOnly() bool {
	return true
}

// Help f
func (Play) Help() [2]string {
	return [2]string{"Plays the FitnessGramPacerTest in your connected voice channel", ""}
}

// Process f
func (Play) Process(context kitty.Context) {
	if isplaying {
		return
	}
	if context.Guild == nil {
		context.Error(errors.New("this command doesn't work in dms"))
		return
	}
	channelID := kitty.VoiceChannelID(context)
	if channelID == "" {
		context.Error(fmt.Errorf("couldn't find the voice channel's id in %s", context.Guild.Name))
		return
	}
	vc, err := context.JoinVoiceChannel(channelID)
	if err != nil {
		context.Error(err)
		return
	}
	defer vc.Disconnect()
	time.Sleep(250 * time.Millisecond)
	vc.Speaking(true)

	isplaying = true
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	vc.Speaking(false)
	time.Sleep(250 * time.Millisecond)
}

// LoadSound loads the FitnessGramPacerTest into memory.
// Yes this is a copypaste of the example from dgo because this is for funzies anyway.
func LoadSound() error {
	file, err := os.Open("FitnessGramPacerTest.dca")
	if err != nil {
		return err
	}
	var opuslen int16
	for {
		err = binary.Read(file, binary.LittleEndian, &opuslen)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}
		if err != nil {
			return err
		}
		inBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &inBuf)
		if err != nil {
			return err
		}
		buffer = append(buffer, inBuf)
	}
}
