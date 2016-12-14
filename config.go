package kitty

import (
	"fmt"
	"io/ioutil"

	"errors"

	"github.com/BurntSushi/toml"
)

// DefaultConfig is the default instance of the `Config`.
var DefaultConfig = &Config{
	Token:   "",
	Prefix:  "?",
	OwnerID: "",
}

// Config maps how a general kitty config should be.
type Config struct {
	Token   string
	Prefix  string
	OwnerID string `toml:"owner_id"`
}

// ReadConfigFromFile reads from disk the config file, unmarshal's it into the Go struct and sets it as the DefaultConfig.
func ReadConfigFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("no config found, exiting.... (%s)", err)
	}
	if err := toml.Unmarshal(file, DefaultConfig); err != nil {
		return fmt.Errorf("failed parsing toml, exiting.... (%s)", err)
	}
	if DefaultConfig.Token == "" || DefaultConfig.OwnerID == "" {
		return errors.New("config requires the token and owner_id fields to be set")
	}
	return nil
}
