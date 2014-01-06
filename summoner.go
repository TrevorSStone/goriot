package goriot

import (
	"fmt"
	"strconv"
)

type masteryPages struct {
	Pages      []MasteryPage `json:"pages"`
	SummonerID int64         `json:"summonerId"`
}

//MasteryPage represents a League of Legends mastery page
type MasteryPage struct {
	Current bool     `json:"current"`
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Talents []Talent `json:"talents"`
}

//Talent is a amstery inside of a Mastery Page
type Talent struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

type runePages struct {
	Pages      []RunePage `json:"pages"`
	SummonerID int        `json:"summonerId"`
}

//RunePage is a League of Legends rune page
type RunePage struct {
	Current bool       `json:"current"`
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Slots   []RuneSlot `json:"slots"`
}

//RuneSlot is a slot for a Rune to go on a RunePage
type RuneSlot struct {
	Rune       Rune `json:"rune"`
	RuneSlotID int  `json:"runeSlotId"`
}

//Rune is a League of Legends Rune
type Rune struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Tier        int    `json:"tier"`
}

//Summoner is a player of League of Legends
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	SummonerLevel int    `json:"summonerLevel"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	RevisionDate  int64  `json:"revisionDate"`
}

//MasteriesBySummoner retrieves the mastery pages of the supplied summonerID from Riot Games API.
//It returns an array of MasteryPage and any errors that occured from the server
//The global API key must be set before use
func MasteriesBySummoner(region string, summonerID int64) (masteries []MasteryPage, err error) {
	if !IsKeySet() {
		return masteries, ErrAPIKeyNotSet
	}
	var pages masteryPages
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v1.2/summoner/%d/masteries?%v", BaseURL, region, summonerID, args)
	err = requestAndUnmarshal(url, &pages)
	if err != nil {
		return
	}
	return pages.Pages, err
}

//RunesBySummoner retrieves the rune pages of the supplied summonerID from Riot Games API.
//It returns an array of RunePage and any errors that occured from the server
//The global API key must be set before use
func RunesBySummoner(region string, summonerID int64) (runes []RunePage, err error) {
	if !IsKeySet() {
		return runes, ErrAPIKeyNotSet
	}
	var pages runePages
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v1.2/summoner/%d/runes?%v", BaseURL, region, summonerID, args)

	err = requestAndUnmarshal(url, &pages)
	if err != nil {
		return
	}
	return pages.Pages, err
}

//SummonerByName retrieves the summoner information of the provided summoner name from Riot Games API.
//It returns a Summoner and any errors that occured from the server
//The global API key must be set before use
func SummonerByName(region string, name string) (summoner Summoner, err error) {
	if !IsKeySet() {
		return summoner, ErrAPIKeyNotSet
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v1.2/summoner/by-name/%v?%v", BaseURL, region, name, args)
	err = requestAndUnmarshal(url, &summoner)
	if err != nil {
		return
	}
	return
}

//SummonerByID retrieves the summoner information of the provided summoner ID from Riot Games API.
//It returns a Summoner and any errors that occured from the server
//The global API key must be set before use
func SummonerByID(region string, summonerID int64) (summoner Summoner, err error) {
	if !IsKeySet() {
		return summoner, ErrAPIKeyNotSet
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v1.2/summoner/%d?%v", BaseURL, region, summonerID, args)

	err = requestAndUnmarshal(url, &summoner)
	if err != nil {
		return
	}
	return
}

//SummonerNamesByID retrieves multiple summoner's information of the provided summoner IDs from Riot Games API.
//It returns an array of Summoner and any errors that occured from the server
//The global API key must be set before use
func SummonerNamesByID(region string, summonerID ...int64) (summoner []Summoner, err error) {
	if !IsKeySet() {
		return summoner, ErrAPIKeyNotSet
	}
	var summoners map[string][]Summoner
	var summonerIDstr string
	for k, v := range summonerID {
		summonerIDstr += strconv.FormatInt(v, 10)
		if k != len(summonerID)-1 {
			summonerIDstr += ","
		}
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v1.2/summoner/%v/name?%v", BaseURL, region, summonerIDstr, args)

	err = requestAndUnmarshal(url, &summoners)
	if err != nil {
		return
	}
	return summoners["summoners"], err
}
