package main

import (
	"flag"

	"github.com/acdenisSK/gol"
	"github.com/acdenisSK/kitty"
	"github.com/acdenisSK/kitty/commands"
	"github.com/bwmarrin/discordgo"
)

var configfile string
var loglevel int
var requestOfflineUsers bool

func init() {
	flag.StringVar(&configfile, "config", "config.toml", "<file name>.toml")
	flag.BoolVar(&requestOfflineUsers, "requestoffusers", false, "")
	flag.IntVar(&loglevel, "loglevel", 3, "<1-to-5>")
	flag.Parse()
	gol.SetLogger(func(filter *gol.Filter) gol.Logger {
		*filter = gol.FilterInfo
		return kitty.Logger{Level: gol.LevelFromInt(loglevel)}
	})
	if err := kitty.ReadConfigFromFile(configfile); err != nil {
		panic(err)
	}
	if err := commands.LoadSound(); err != nil {
		gol.Warnf("Couldn't load sound file because of the error: `%s` (note that this isn't fully required for the bot to work)", err)
	}
}

func main() {
	session, err := discordgo.New("Bot " + kitty.DefaultConfig.Token)
	if err != nil {
		kitty.Panicf("Error creating a discordgo session: %s", err)
	}
	listOfCommands := map[string]kitty.Command{
		"ping":          commands.Ping{},
		"stats":         commands.Stats{},
		"eval":          commands.Eval{},
		"sinfo":         commands.Sinfo{},
		"clean":         commands.Clean{},
		"sort":          commands.Sort{},
		"randomnumbers": commands.RandomNumbers{},
		"play":          commands.Play{},
		"cat":           commands.Cat{},
		"ddg":           commands.Duckduckgo{},
		"tag":           commands.Tag{},
		"calc":          commands.Calc{},
	}
	listOfCommands["help"] = commands.Help{Commands: listOfCommands}
	session.AddHandlerOnce(kitty.Ready(requestOfflineUsers)) // no need to listen for the READY event multiple times
	session.AddHandler(kitty.GuildMemberChunk)
	session.AddHandler(kitty.MessageCreate(listOfCommands))
	if err := session.Open(); err != nil {
		kitty.Panicf("Error opening websocket: %s", err)
	}
	<-make(chan struct{})
}
