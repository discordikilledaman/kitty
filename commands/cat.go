package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/acdenisSK/kitty"
)

// Cat f
type Cat struct{}

// IsOwnerOnly f
func (Cat) IsOwnerOnly() bool {
	return false
}

// Help f
func (Cat) Help() [2]string {
	return [2]string{"Makes a request to the random.cat api", ""}
}

// Process f
func (Cat) Process(context kitty.Context) {
	res, err := http.Get("http://random.cat/meow")
	if err != nil {
		context.Error(err)
		return
	}
	defer res.Body.Close()

	catJSON := struct {
		File string `json:"file"`
	}{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		context.Error(err)
		return
	}
	if err := json.Unmarshal(body, &catJSON); err != nil {
		context.Error(err)
		return
	}
	context.Say(catJSON.File)
}
