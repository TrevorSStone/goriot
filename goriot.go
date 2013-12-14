/*
Package goriot is a library that provides a means of connecting and recieving data from Riot's League of Legend API
*/
package goriot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrAPIKeyNotSet = errors.New("goriot: API key has not been set. If you need a key visit https://developer.riotgames.com/")
	apikey          string
	BaseURL         = "https://prod.api.pvp.net/api/"
	NA              = "na"
	EUW             = "euw"
	EUNE            = "eune"
	SEASON3         = "SEASON3"
	SEASON4         = "SEASON4"
)

func SetAPIKey(key string) {
	apikey = key

}

func IsKeySet() bool {
	return !(apikey == "")
}

func RequestAndUnmarshal(url string, v interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Error: HTTP Status %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return
	}
	return
}
