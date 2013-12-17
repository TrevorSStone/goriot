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
	smallRateChan   rateChan
	longRateChan    rateChan
)

type rateChan struct {
	RateQueue   chan bool
	TriggerChan chan bool
}

func SetAPIKey(key string) {
	apikey = key
}

func SetSmallRateLimit(numrequests int, pertime time.Duration) {
	smallRateChan = rateChan{
		RateQueue:   make(chan bool, numrequests),
		TriggerChan: make(chan bool),
	}
	go rateLimitHandler(smallRateChan, pertime)
}
func SetLongRateLimit(numrequests int, pertime time.Duration) {
	longRateChan = rateChan{
		RateQueue:   make(chan bool, numrequests),
		TriggerChan: make(chan bool),
	}
	go rateLimitHandler(longRateChan, pertime)
}

func rateLimitHandler(RateChan rateChan, pertime time.Duration) {
	returnChan := make(chan bool)
	go timeTriggerWatcher(RateChan.TriggerChan, returnChan)
	for {
		<-returnChan
		<-time.After(pertime)
		go timeTriggerWatcher(RateChan.TriggerChan, returnChan)
		length := len(RateChan.RateQueue)
		for i := 0; i < length; i++ {
			<-RateChan.RateQueue
		}
	}
}

func timeTriggerWatcher(timeTrigger chan bool, returnChan chan bool) {
	timeTrigger <- true
	returnChan <- true
}

func IsKeySet() bool {
	return !(apikey == "")
}

func RequestAndUnmarshal(requestURL string, v interface{}) (err error) {
	checkRateLimiter(smallRateChan)
	checkRateLimiter(longRateChan)
	resp, err := http.Get(requestURL)
	if err != nil {
		return
	}
	checkTimeTrigger(smallRateChan)
	checkTimeTrigger(longRateChan)
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

func checkRateLimiter(RateChan rateChan) {
	if RateChan.RateQueue != nil && RateChan.TriggerChan != nil {
		RateChan.RateQueue <- true
	}
}

func checkTimeTrigger(RateChan rateChan) {
	if RateChan.RateQueue != nil && RateChan.TriggerChan != nil {
		select {
		case <-RateChan.TriggerChan:
		default:
		}
	}
}
