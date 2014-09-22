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
	"strconv"
	"strings"
	"time"
)

var (
	apikey string
	//BaseURL is the base of the url used by Riot's API service
	BaseURL = "api.pvp.net/api"
	//BR represents the string for the Brazilian League of Legends Servers,
	//only used as a helper to prevent typos
	BR = "br"
	//EUNE represents the string for the North Eastern European League of Legends Servers,
	//only used as a helper to prevent typos
	EUNE = "eune"
	//EUW represents the string for the West European League of Legends Servers,
	//only used as a helper to prevent typos
	EUW = "euw"
	//KR represents the string for the Korean League of Legends Servers,
	//only used as a helper to prevent typos
	KR = "kr"
	//LAN represents the string for the Latin America North League of Legends Servers,
	//only used as a helper to prevent typos
	LAN = "lan"
	//LAS represents the string for the Latin America South League of Legends Servers,
	//only used as a helper to prevent typos
	LAS = "las"
	//NA represents the string for the North American League of Legends Servers,
	//only used as a helper to prevent typos
	NA = "na"
	//OCE represents the string for the Oceanic League of Legends Servers,
	//only used as a helper to prevent typos
	OCE = "oce"
	//RU represents the string for the Russian League of Legends Servers,
	//only used as a helper to prevent typos
	RU = "ru"
	//TR represents the string for the Turkish League of Legends Servers,
	//only used as a helper to prevent typos
	TR = "tr"
	//SEASON3 is the string of "SEASON3". Used to help prevent typos
	SEASON3 = "SEASON3"
	//SEASON4 is the string of "SEASON4". Used to help prevent typos
	SEASON4 = "SEASON4"
	//Ranked Solo 5s
	RANKED_SOLO_5x5 = "RANKED_SOLO_5x5"
	//Ranked Team 3s
	RANKED_TEAM_3x3 = "RANKED_TEAM_3x3"
	//Ranked Team 5s
	RANKED_TEAM_5x5 = "RANKED_TEAM_5x5"
	//ErrAPIKeyNotSet is the error returned when no global API key has been set
	ErrAPIKeyNotSet = errors.New("goriot: API key has not been set. If you need a key visit https://developer.riotgames.com/")
	smallRateChan   rateChan
	longRateChan    rateChan
)

type rateChan struct {
	RateQueue   chan bool
	TriggerChan chan bool
}

//RiotError contains the http status code of the erro
type RiotError struct {
	StatusCode int
}

// types for URL query parameter formatting
type strURLParameter []string
type intURLParameter []int64

// Creates a comma separated string
func (q strURLParameter) String() string {
	return strings.Join(q, ",")
}

// Convers int64 slice to string slice and returns
// comma separated string
func (q intURLParameter) String() string {
	s := make(strURLParameter, len(q))
	for k, v := range q {
		s[k] = strconv.FormatInt(v, 10)
	}

	return s.String()
}

//SetAPIKey sets the global key for the Riot Games API. If you do not have one you can get one for free at
//https://developer.riotgames.com/
func SetAPIKey(key string) {
	apikey = key
}

//SetSmallRateLimit allows a custom rate limit to be set. For at the time of this writing the default
//for a development API key is 10 requests every 10 seconds
func SetSmallRateLimit(numrequests int, pertime time.Duration) {
	smallRateChan = rateChan{
		RateQueue:   make(chan bool, numrequests),
		TriggerChan: make(chan bool),
	}
	go rateLimitHandler(smallRateChan, pertime)
}

//SetLongRateLimit allows a custom rate limit to be set. For at the time of this writing the default
//for a development API key is 500 requests every 10 minutes
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

//IsKeySet returns a bool of if the global API key has been set
func IsKeySet() bool {
	return !(apikey == "")
}

func requestAndUnmarshal(requestURL string, v interface{}) (err error) {
	checkRateLimiter(smallRateChan)
	checkRateLimiter(longRateChan)
	resp, err := http.Get(requestURL)
	if err != nil {
		return
	}
	checkTimeTrigger(smallRateChan)
	checkTimeTrigger(longRateChan)
	if resp.StatusCode != http.StatusOK {
		return RiotError{StatusCode: resp.StatusCode}
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

//Error prints the error message for a RiotError
func (err RiotError) Error() string {
	return fmt.Sprintf("Error: HTTP Status %d", err.StatusCode)
}
