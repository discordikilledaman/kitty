package commands

import (
	"fmt"

	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
)

// Sinfo f
type Sinfo struct{}

// Checks f
func (Sinfo) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (Sinfo) Process(context kitty.Context) {
	guild := context.Guild
	owner, err := context.State.Member(guild.ID, guild.OwnerID)
	if err != nil {
		return
	}
	var roles, channels string
	for i, r := range guild.Roles {
		if r.Name == "@everyone" {
			continue
		} else if i == len(guild.Roles)-1 {
			roles += r.Name
			break
		}
		roles += r.Name + ", "
	}
	for i, c := range guild.Channels {
		if i == len(guild.Channels)-1 {
			channels += c.Name
			break
		}
		channels += c.Name + ", "
	}
	var bots int
	for _, u := range guild.Members {
		if u.User.Bot {
			bots++
		}
	}
	context.SayEmbed(&discordgo.MessageEmbed{
		Title: "Information about **" + guild.Name + "**",
		Fields: []*discordgo.MessageEmbedField{
			kitty.Field("Owner", owner.User.Username),
			kitty.Field("Member Count", fmt.Sprintf("%d (bots %d)", guild.MemberCount, bots)),
			kitty.Field("Roles", fmt.Sprintf("%s (%d)", roles, len(guild.Roles))),
			kitty.Field("Channels", fmt.Sprintf("%s (%d)", channels, len(guild.Channels))),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: discordgo.EndpointGuildIcon(guild.ID, guild.Icon),
		},
		Color: 45475,
	})
}
