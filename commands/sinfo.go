package commands

import (
	"errors"
	"fmt"
	"sort"

	"strings"

	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
)

// Sinfo f
type Sinfo struct{}

// IsOwnerOnly f
func (Sinfo) IsOwnerOnly() bool {
	return false
}

// Help f
func (Sinfo) Help() [2]string {
	return [2]string{"Shows information about the server", ""}
}

// Process f
func (Sinfo) Process(context kitty.Context) {
	if context.Guild == nil {
		context.Error(errors.New("this command doesn't work in dms"))
		return
	}
	guild := context.Guild
	sort.Sort(discordgo.Roles(guild.Roles))
	sort.Slice(guild.Channels, func(i, j int) bool { return guild.Channels[i].Position < guild.Channels[j].Position })
	owner, err := context.State.Member(guild.ID, guild.OwnerID)
	if err != nil {
		return
	}
	embed := kitty.NewEmbed("Information about **" + guild.Name + "**")
	embed.Field("Owner", owner.User.Username)
	embed.Field("Member Count", "All / Normal / Bots\n"+memberCount(guild))
	embed.Field("Server Region", guild.Region)
	embed.Fieldf("Roles", "%s (%d)", roleNames(guild), len(guild.Roles))
	embed.Fieldf("Channels", "%s (%d)", channelNames(guild.Channels), len(guild.Channels))
	embed.Thumbnail(discordgo.EndpointGuildIcon(guild.ID, guild.Icon))
	context.SayEmbed(embed)
}

func channelNames(channels []*discordgo.Channel) string {
	var channelnames []string
	for index, channel := range channels {
		if index >= 15 {
			break
		}
		channelnames = append(channelnames, channel.Name)
	}
	channelss := strings.Join(channelnames, ", ")
	if len(channelnames) >= 15 {
		channelss += ", ..."
	}
	return channelss
}

func roleNames(guild *discordgo.Guild) string {
	var rolenames []string
	for index, role := range guild.Roles {
		if role.ID == guild.ID {
			continue
		} else if index >= 15 {
			break
		}
		rolenames = append(rolenames, role.Name)
	}
	roless := strings.Join(rolenames, ", ")
	if len(rolenames) >= 15 {
		roless += ", ..."
	}
	return roless
}

func memberCount(guild *discordgo.Guild) string {
	var bots int
	for _, u := range guild.Members {
		if u.User.Bot {
			bots++
		}
	}
	return fmt.Sprintf("%d / %d / %d", guild.MemberCount, guild.MemberCount-bots, bots)
}
