package main

import (
	"os"

	"github.com/acdenisSK/kitty"
	"github.com/acdenisSK/kitty/commands"
	"github.com/bwmarrin/discordgo"
)

func main() {
	conf := kitty.GetConfig()
	logger := &kitty.Logger{}
	logger.AddOutput(os.Stdout)
	if conf.Logging.File != "" {
		file, err := os.OpenFile(conf.Logging.File, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			logger.Panicln("error opening file:", err)
		}
		defer file.Close()
		logger.AddOutput(file)
	}
	logger.Setup()
	client, err := discordgo.New("Bot " + conf.Required.Token)
	if err != nil {
		logger.Panicln("failed creating an instance of discordgo:", err)
	}
	listOfCommands := map[string]kitty.Command{
		"ping":     &commands.Ping{},
		"stats":    &commands.Stats{},
		"eval":     &commands.Eval{},
		"sinfo":    &commands.Sinfo{},
		"randoimg": &commands.RandoImg{},
		"clean":    &commands.Clean{},
		"sort":     &commands.Sort{},
	}
	listOfCommands["help"] = &commands.Help{Commands: listOfCommands}
	kitty := kitty.New(logger, listOfCommands, conf, client)
	kitty.Setup()
	<-make(chan struct{})
}
