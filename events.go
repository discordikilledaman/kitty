package kitty

import (
	"strings"

	"github.com/acdenisSK/gol"
	"github.com/bwmarrin/discordgo"
)

// Ready handles the READY event.
func Ready(chunkMembers bool) func(*discordgo.Session, *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		gol.Info("Running!")
		if chunkMembers {
			for _, g := range r.Guilds {
				if g.Large {
					s.RequestGuildMembers(g.ID, "", 0)
				}
			}
		}
	}
}

// GuildMemberChunk handles the GUILD_MEMBER_CHUNK event.
func GuildMemberChunk(s *discordgo.Session, c *discordgo.GuildMembersChunk) {
	for _, g := range s.State.Guilds {
		if g.ID == c.GuildID {
			g.Members = append(g.Members, c.Members...)
		}
	}
}

// MessageCreate handles the MESSAGE_CREATE event while also handling command parsing, execution, etc.
func MessageCreate(commands map[string]Command) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, msg *discordgo.MessageCreate) {
		defer func() {
			if r := recover(); r != nil {
				gol.Info("Recovered")
			}
		}()
		if checkMessage(msg.Message, s.State.User.ID) {
			return
		}
		args := strings.Fields(msg.Content[len(DefaultConfig.Prefix):])
		command, args := args[0], args[1:]
		cmd, ok := commands[strings.ToLower(command)]
		if !ok {
			return
		}
		if cmd.IsOwnerOnly() && msg.Author.ID != DefaultConfig.OwnerID {
			return
		}
		channel, err := s.State.Channel(msg.ChannelID)
		if err != nil {
			gol.Error(err)
			return
		}
		guild := guildFromState(channel, s.State)

		gol.Info(LogString(msg.Author, command, channel, guild))
		CommandCounter.Update(command)
		go cmd.Process(NewContext(s, msg.Message, channel, guild, args))
	}
}
