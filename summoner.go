package goriot

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type masteryBook struct {
	Pages      []MasteryPage `json:"pages"`
	SummonerID int64         `json:"summonerId"`
}

//MasteryPage represents a League of Legends mastery page
type MasteryPage struct {
	Current   bool      `json:"current"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Masteries []Mastery `json:"masteries"`
}

//Mastery located inside a page
type Mastery struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}

type runeBook struct {
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
	RuneId     int `json:"runeId"`
	RuneSlotID int `json:"runeSlotId"`
}

//Summoner is a player of League of Legends
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	SummonerLevel int    `json:"summonerLevel"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	RevisionDate  int64  `json:"revisionDate"`
}

//MasteriesBySummoner retrieves the mastery pages of the supplied summonerIDs from Riot Games API.
//It returns a map of array pf MasteryPage with the key being the summonerID and any errors that occured from the server
//The global API key must be set before use
func MasteriesBySummoner(region string, summonerID ...int64) (masteries map[int64][]MasteryPage, err error) {
	if !IsKeySet() {
		return masteries, ErrAPIKeyNotSet
	}
	masteries = make(map[int64][]MasteryPage)
	var pages map[string]masteryBook
	summonerIDstr := intURLParameter(summonerID).String()

	args := "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.4/summoner/%v/masteries?%v", region, BaseURL, region, summonerIDstr, args)
	err = requestAndUnmarshal(url, &pages)
	if err != nil {
		return
	}
	for k, v := range pages {
		n, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return masteries, err
		}
		masteries[n] = v.Pages
	}
	return masteries, err
}

//RunesBySummoner retrieves the rune pages of the supplied summonerIDs from Riot Games API.
//It returns a map of array of RunePage with the key being the summonerID and any errors that occured from the server
//The global API key must be set before use
func RunesBySummoner(region string, summonerID ...int64) (runes map[int64][]RunePage, err error) {
	if !IsKeySet() {
		return runes, ErrAPIKeyNotSet
	}
	runes = make(map[int64][]RunePage)
	var pages map[string]runeBook
	summonerIDstr := intURLParameter(summonerID).String()

	args := "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.4/summoner/%v/runes?%v", region, BaseURL, region, summonerIDstr, args)

	err = requestAndUnmarshal(url, &pages)
	if err != nil {
		return
	}
	for k, v := range pages {
		n, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return runes, err
		}
		runes[n] = v.Pages
	}
	return runes, err
}

//SummonerByName retrieves the summoner information of the provided summoner names from Riot Games API.
//It returns a Map of Summoner with the key being the summoner name and any errors that occured from the server
//The global API key must be set before use
//WARNING: The map's key is not necessarily the same string used in the request. It is
//recommended to use NormalizeSummonerName before calling this function
func SummonerByName(region string, name ...string) (summoners map[string]Summoner, err error) {
	if !IsKeySet() {
		return summoners, ErrAPIKeyNotSet
	}
	names := strURLParameter(name).String()
	summoners = make(map[string]Summoner)
	args := "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.4/summoner/by-name/%v?%v", region, BaseURL, region, names, args)
	err = requestAndUnmarshal(url, &summoners)
	if err != nil {
		return
	}
	return
}

//SummonerByID retrieves the summoner information of the provided summoner IDs from Riot Games API.
//It returns a map of Summoner with the key being summonerID and any errors that occured from the server
//The global API key must be set before use
func SummonerByID(region string, summonerID ...int64) (summoners map[int64]Summoner, err error) {
	if !IsKeySet() {
		return summoners, ErrAPIKeyNotSet
	}

	var summonersString map[string]Summoner
	summoners = make(map[int64]Summoner)
	summonerIDstr := intURLParameter(summonerID).String()

	args := "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.4/summoner/%v?%v", region, BaseURL, region, summonerIDstr, args)

	err = requestAndUnmarshal(url, &summonersString)
	if err != nil {
		return
	}
	for k, v := range summonersString {
		n, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return summoners, err
		}
		summoners[n] = v
	}
	return
}

//SummonerNamesByID retrieves multiple summoner's information of the provided summoner IDs from Riot Games API.
//It returns an map of string with the key being the summonerID and any errors that occured from the server
//The global API key must be set before use
func SummonerNamesByID(region string, summonerID ...int64) (summoners map[int64]string, err error) {
	if !IsKeySet() {
		return summoners, ErrAPIKeyNotSet
	}
	var summonerNames map[string]string
	summoners = make(map[int64]string)
	summonerIDstr := intURLParameter(summonerID).String()

	args := "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.4/summoner/%v/name?%v", region, BaseURL, region, summonerIDstr, args)

	err = requestAndUnmarshal(url, &summonerNames)
	if err != nil {
		return
	}
	for k, v := range summonerNames {
		n, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return summoners, err
		}
		summoners[n] = v
	}
	return summoners, err
}

func createSummonerIDString(summonerID []int64) (summonerIDstr string, err error) {
	if len(summonerID) > 40 {
		return summonerIDstr, errors.New("A Maximum of 40 SummonerIDs are allowed")
	}
	for k, v := range summonerID {
		summonerIDstr += strconv.FormatInt(v, 10)
		if k != len(summonerID)-1 {
			summonerIDstr += ","
		}
	}
	return
}

//NormalizeSummonerName takes an arbitrary number of strings and returns a string array containing the strings
//standardized to league of legends internal standard (lowecase and strings removed)
func NormalizeSummonerName(summonerNames ...string) []string {
	for i, v := range summonerNames {
		summonerName := strings.ToLower(v)
		summonerName = strings.Replace(summonerName, " ", "", -1)
		summonerNames[i] = summonerName
	}
	return summonerNames
}
