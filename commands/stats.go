package commands

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"time"

	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
	humanize "github.com/dustin/go-humanize"
)

var uptime = time.Now()

// Stats f
type Stats struct{}

// IsOwnerOnly f
func (Stats) IsOwnerOnly() bool {
	return false
}

// Help f
func (Stats) Help() [2]string {
	return [2]string{"Shows stats about the bot", ""}
}

// Process f
func (Stats) Process(context kitty.Context) {
	var users, channels int
	for _, g := range context.State.Guilds {
		users += len(g.Members)
		channels += len(g.Channels)
	}
	counter := kitty.CommandCounter.Counter
	var mostused string
	var nums []int
	for _, num := range counter {
		nums = append(nums, num)
	}
	sort.Ints(nums)
	highest := nums[len(nums)-1]
	for command, num := range counter {
		if num == highest {
			mostused = fmt.Sprintf("%s (%d)", command, num)
		}
	}
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	embed := kitty.NewEmbed("Stats about " + context.State.User.Username)
	embed.Field("Go version", runtime.Version()[2:])
	embed.Fieldf("Lib", "discordgo (%s)", discordgo.VERSION)
	embed.Field("Uptime", formattedTime(time.Since(uptime)))
	embed.Field("Running Tasks", strconv.Itoa(runtime.NumGoroutine()))
	embed.Fieldf("Users | Channels | Guilds", "%d | %d | %d", users, channels, len(context.State.Guilds))
	embed.Fieldf("Memory Usage", "total / heap / gc'd\n%s / %s / %s", humanize.Bytes(memStats.Sys), humanize.Bytes(memStats.Alloc), humanize.Bytes(memStats.TotalAlloc))
	embed.Field("Most used command", mostused)
	embed.Field("Source code", "[github](https://github.com/acdenisSK/kitty)")
	embed.Thumbnail(discordgo.EndpointUserAvatar(context.State.User.ID, context.State.User.Avatar))
	context.SayEmbed(embed)
}

func formattedTime(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}
