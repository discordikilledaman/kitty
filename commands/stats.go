package commands

import (
	"fmt"
	"runtime"
	"time"

	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
	humanize "github.com/dustin/go-humanize"
)

var uptime = time.Now()

func getFormattedTime(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}

// Stats f
type Stats struct{}

// Checks f
func (Stats) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (Stats) Process(context kitty.Context) {
	var users, channels int
	for _, g := range context.State.Guilds {
		users += len(g.Members)
		channels += len(g.Channels)
	}
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	context.SayEmbed(&discordgo.MessageEmbed{
		Title: "Stats about " + context.State.User.Username,
		Fields: []*discordgo.MessageEmbedField{
			kitty.Field("Go", runtime.Version()[2:]),
			kitty.Field("Lib", fmt.Sprintf("discordgo (%s)", discordgo.VERSION)),
			kitty.Field("Uptime", getFormattedTime(time.Since(uptime))),
			kitty.Field("Memory Usage (total / heap / gc'd)", fmt.Sprintf("%s / %s / %s", humanize.Bytes(memStats.Sys), humanize.Bytes(memStats.Alloc), humanize.Bytes(memStats.TotalAlloc))),
			kitty.Field("Running Tasks", fmt.Sprint(runtime.NumGoroutine())),
			kitty.Field("Users | Channels | Guilds", fmt.Sprintf("%d | %d | %d", users, channels, len(context.State.Guilds))),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: discordgo.EndpointUserAvatar(context.State.User.ID, context.State.User.Avatar),
		},
		Color:     0x79c879,
		Timestamp: kitty.GetISOTimestamp(),
	})
}
