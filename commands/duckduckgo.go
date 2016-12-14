package commands

import (
	"fmt"
	"strings"

	"github.com/acdenisSK/kitty"
	"github.com/ajanicij/goduckgo/goduckgo"
)

// Duckduckgo uses duckduckgo's api to search things.
type Duckduckgo struct{}

// IsOwnerOnly f
func (Duckduckgo) IsOwnerOnly() bool {
	return false
}

// Help f
func (Duckduckgo) Help() [2]string {
	return [2]string{"Searches the web with duckduckgo", "<query>"}
}

// Process f
func (Duckduckgo) Process(context kitty.Context) {
	query := strings.Join(context.Args, " ")
	msg, err := goduckgo.Query(query)
	if err != nil {
		context.Error(err)
		return
	}
	thumbnail := "https://images.duckduckgo.com/iu/?u=http%3A%2F%2Fcore2.staticworld.net%2Fimages%2Farticle%2F2014%2F05%2Fduckduckgo-logo-100266737-large.png&f=1"
	var onlyRedirect bool
	if strings.HasPrefix(query, "!") {
		onlyRedirect = true
		thumbnail = ""
	}
	embed := kitty.NewEmbed("Results for your query")
	var description string
	if msg.Redirect != "" && onlyRedirect {
		description = msg.Redirect
	}
	if msg.Results != nil && len(msg.Results) != 0 && !onlyRedirect {
		for index, result := range msg.Results {
			if len(msg.Results) > 4 && index > 4 {
				break
			}
			if !result.Icon.IsEmpty() {
				thumbnail = result.Icon.URL
			}
			embed.Fieldf("Result", "[%s](%s)", result.Text, result.FirstURL)
		}
	}
	if msg.AbstractText != "" && !onlyRedirect {
		embed.Field("Summary", msg.AbstractText)
	} else if msg.AbstractURL != "" && msg.AbstractText != "" && !onlyRedirect {
		embed.Fieldf("Summary", "[Link](%s)\n\n%s", msg.AbstractURL, msg.AbstractText)
	}
	if msg.RelatedTopics != nil && len(msg.RelatedTopics) != 0 && !onlyRedirect {
		var topics []string
		for index, topic := range msg.RelatedTopics {
			if len(msg.RelatedTopics) > 1 && index > 1 {
				break
			}
			topics = append(topics, fmt.Sprintf("[%s](%s)", topic.Text, topic.FirstURL))
		}
		embed.Field("Related topics", strings.Join(topics, "\n\n"))
	}
	if msg.Type != "" && !onlyRedirect {
		embed.Field("Category", toCategory(msg.Type))
	}
	if msg.Heading != "" && !onlyRedirect {
		embed.Field("Main topic", msg.Heading)
	}
	if msg.Answer != "" && !onlyRedirect {
		embed.Field("Answer", msg.Answer)
	}
	if msg.Definition != "" && !onlyRedirect {
		embed.Field("Definition", msg.Definition)
	}
	if msg.Image != "" && !onlyRedirect {
		thumbnail = msg.Image
	}
	embed.Thumbnail(thumbnail)
	embed.Color = 0xec1f26
	if onlyRedirect {
		embed.Title = "Redirection for your query"
		embed.Description = description
	}
	context.SayEmbed(embed)
}

func toCategory(msgtype string) string {
	switch msgtype {
	case "A":
		return "article"
	case "C":
		return "category"
	case "D":
		return "disambiguation"
	case "E":
		return "exclusive"
	case "N":
		return "name"
	default:
		return "nothing"
	}
}
