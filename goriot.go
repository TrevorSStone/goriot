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
	"runtime"
	"time"
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
	smallRateChan   chan bool
	longRateChan    chan bool
)

func SetAPIKey(key string) {
	apikey = key
}

func SetSmallRateLimit(numrequests int, pertime time.Duration) {
	smallRateChan = make(chan bool, numrequests-1)
	go rateLimitHandler(smallRateChan, pertime+time.Second)
}
func SetLongRateLimit(numrequests int, pertime time.Duration) {
	longRateChan = make(chan bool, numrequests-1)
	go rateLimitHandler(longRateChan, pertime+time.Second)
}

func rateLimitHandler(rateChan chan bool, pertime time.Duration) {
	for {
		<-rateChan
		<-time.After(pertime)
		length := len(rateChan)
		for i := 0; i < length; i++ {
			<-rateChan
		}
	}
}

func IsKeySet() bool {
	return !(apikey == "")
}

func RequestAndUnmarshal(requestURL string, v interface{}) (err error) {
	if smallRateChan != nil {
		smallRateChan <- true
		runtime.Gosched()
	}
	if longRateChan != nil {
		longRateChan <- true
		runtime.Gosched()
	}

	resp, err := http.Get(requestURL)
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
