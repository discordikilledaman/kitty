package kitty

import (
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/naoina/toml"

	"sync"

	"github.com/bwmarrin/discordgo"
)

// Config maps how a general kitty config should be.
type Config struct {
	Required struct {
		Token  string
		Prefix string
	}
	Logging struct {
		File string
	}
}

// GetConfig is used for acquiring configuration from a file called "config.toml"
// If it does not exist, this will panic.
func GetConfig() Config {
	file, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Panicln("no config found, exiting....")
	}
	configData := Config{}
	if err := toml.Unmarshal(file, &configData); err != nil {
		log.Panicln("failed parsing toml, exiting....")
	}
	return configData
}

// GetISOTimestamp returns an ISO6301 based timestamp from the current time.
func GetISOTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.070000")
}

// Logger is a wrapper around `log`'s one but adds support for multiple outputs.
type Logger struct {
	*log.Logger
	sync.Mutex
	outputs []io.Writer
}

// AddOutput adds an output for the logger.
// This is a one-time job so be careful on what writers you give this logger.
func (l *Logger) AddOutput(writer io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.outputs = append(l.outputs, writer)
}

// Setup setups the logger
func (l *Logger) Setup() {
	l.Lock()
	defer l.Unlock()
	multi := io.MultiWriter(l.outputs...)
	l.Logger = log.New(multi, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// Field is a shortcut to boilerplate.
func Field(name, value string) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{Name: name, Value: value, Inline: true}
}
