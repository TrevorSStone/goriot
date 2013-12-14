package goriot

import (
	"strconv"
)

type masteryPages struct {
	Pages      []MasteryPage
	SummonerID int64
}

type MasteryPage struct {
	Current bool
	Name    string
	Talents []Talent
}

type Talent struct {
	ID   int
	Name string
	Rank int
}

type runePages struct {
	Pages      []RunePage
	SummonerID int
}

type RunePage struct {
	Current bool
	ID      int64
	Name    string
	Slots   []RuneSlot
}

type RuneSlot struct {
	Rune       Rune
	RuneSlotID int
}

type Rune struct {
	Description string
	ID          int
	Name        string
	Tier        int
}

type Summoner struct {
	ProfileIconId int
	SummonerLevel int
	ID            int64
	Name          string
	RevisionDate  int64
}

func GetMasteriesBySummoner(region string, summonerID int64) (masteries []MasteryPage, err error) {
	if !IsKeySet() {
		return masteries, ErrAPIKeyNotSet
	}
	var pages masteryPages
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	url := BaseURL + "lol/" + region + "/v1.1/summoner/" + summonerIDstr + "/masteries"
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &pages)
	if err != nil {
		return
	}
	return pages.Pages, err
}

func GetRunesBySummoner(region string, summonerID int64) (runes []RunePage, err error) {
	if !IsKeySet() {
		return runes, ErrAPIKeyNotSet
	}
	var pages runePages
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	url := BaseURL + "lol/" + region + "/v1.1/summoner/" + summonerIDstr + "/runes"
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &pages)
	if err != nil {
		return
	}
	return pages.Pages, err
}

func GetSummonerByName(region string, name string) (summoner Summoner, err error) {
	if !IsKeySet() {
		return summoner, ErrAPIKeyNotSet
	}
	url := BaseURL + "lol/" + region + "/v1.1/summoner/by-name/" + name
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &summoner)
	if err != nil {
		return
	}
	return
}

func GetSummonerByID(region string, summonerID int64) (summoner Summoner, err error) {
	if !IsKeySet() {
		return summoner, ErrAPIKeyNotSet
	}
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	url := BaseURL + "lol/" + region + "/v1.1/summoner/" + summonerIDstr
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &summoner)
	if err != nil {
		return
	}
	return
}

func GetSummonerNamesByID(region string, summonerID ...int64) (summoner []Summoner, err error) {
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
	url := BaseURL + "lol/" + region + "/v1.1/summoner/" + summonerIDstr + "/name"
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &summoners)
	if err != nil {
		return
	}
	return summoners["summoners"], err
}
