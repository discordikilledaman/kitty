package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/acdenisSK/kitty"
	"github.com/bwmarrin/discordgo"
)

// probably bad that i'm abusing toml like this.
var db = func() database {
	file, err := ioutil.ReadFile("tags.toml")
	if err != nil {
		kitty.Panicf("Error opening database file: %s", err)
	}
	var a database
	if err := toml.Unmarshal(file, &a); err != nil {
		kitty.Panicf("Error unmarshalling the file: %s", err)
	}
	return a
}()

// Tag f
type Tag struct{}

// IsOwnerOnly f
func (Tag) IsOwnerOnly() bool {
	return false
}

// Help f
func (Tag) Help() [2]string {
	return [2]string{}
}

// Process f
func (Tag) Process(context kitty.Context) {
	switch context.Args[0] {
	case "get":
		context.Args = context.Args[1:]
		name := strings.Join(context.Args, " ")
		hasFlag := strings.Contains(name, "-noE")
		if hasFlag {
			name = strings.Trim(strings.Replace(name, "-noE", "", 1), " ")
		}
		tag := db.get(context.Author.ID, name, guildID(context.Guild))
		if tag == "" {
			// cfes == "couldn't find error, and suggestion"
			cfes := `Sorry couldn't find your tag in my database. Perhaps try creating it first?
If you're sure it does exist, please check that your name matches exactly (tag names are case sensitive and are per server).`
			context.Say(cfes)
			return
		}
		if hasFlag {
			context.Say(tag)
			return
		}
		embed := kitty.NewEmbed(name)
		embed.Description = tag
		embed.Footer("If you don't want to see this in an embed, add `-noE` at the end of your command call.")
		context.SayEmbed(embed)
	case "add", "addglobal":
		var id string
		if context.Args[0] == "addglobal" {
			id = "global"
		} else {
			id = guildID(context.Guild)
		}
		context.Args = context.Args[1:]
		err := db.add(context.Author.ID, context.Args[0], strings.Join(context.Args[1:], " "), id)
		if err != nil {
			context.Error(err)
			return
		}
		context.Say("Successfully added your tag!")
	case "edit":
		context.Args = context.Args[1:]
		db.edit(context.Author.ID, context.Args[0], strings.Join(context.Args[1:], " "), guildID(context.Guild))
		context.Say("Successfully edited your tag!")
	case "delete":
		context.Args = context.Args[1:]
		db.remove(context.Author.ID, strings.Join(context.Args, " "), guildID(context.Guild))
		context.Say("Successfully removed your tag!")
	}
}

type tag struct {
	Author    string    `toml:"author"`
	Name      string    `toml:"name"`
	Content   string    `toml:"content"`
	GuildID   string    `toml:"guildid"`
	CreatedAt time.Time `toml:"created_at"`
}

type database struct {
	Tags []tag `toml:"tag"`
}

func (d database) save() error {
	file, err := os.Create("tags.toml")
	if err != nil {
		return err
	}
	defer file.Close()
	if err := toml.NewEncoder(file).Encode(d); err != nil {
		return err
	}
	return nil
}

func (d database) add(authorID, name, content, guildID string) error {
	for _, tag := range d.Tags {
		if tag.Author == authorID && tag.Name == name && (tag.GuildID == "global" || tag.GuildID == guildID) {
			return fmt.Errorf("the tag `%s` already exists", name)
		}
	}
	d.Tags = append(d.Tags, tag{authorID, name, content, guildID, time.Now()})
	if err := d.save(); err != nil {
		kitty.Panicf("Error saving db: %s", err)
	}
	return nil
}

// Optimise these three

func (d database) get(authorID, name, guildID string) string {
	for _, tag := range d.Tags {
		if tag.Author == authorID && tag.Name == name && (tag.GuildID == "global" || tag.GuildID == guildID) {
			return tag.Content
		}
	}
	return ""
}

func (d database) remove(authorID, name, guildID string) {
	for index, tag := range d.Tags {
		if tag.Author == authorID && tag.Name == name && (tag.GuildID == "global" || tag.GuildID == guildID) {
			d.Tags = append(d.Tags[:index], d.Tags[index+1:]...)
		}
	}
	if err := d.save(); err != nil {
		kitty.Panicf("Error saving db: %s", err)
	}
}

func (d database) edit(authorID, name, content, guildID string) {
	for _, tag := range d.Tags {
		if tag.Author == authorID && tag.Name == name && (tag.GuildID == "global" || tag.GuildID == guildID) {
			tag.Content = content
		}
	}
	if err := d.save(); err != nil {
		kitty.Panicf("Error saving db: %s", err)
	}
}

func guildID(guild *discordgo.Guild) string {
	if guild == nil {
		return "global" // assume this is a dm and the user wants the tag globally.
	}
	return guild.ID
}
