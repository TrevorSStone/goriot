package goriot

import (
	"strconv"
)

type League struct {
	Entries   []LeagueItem
	Name      string
	Queue     string
	Tier      string
	Timestamp int64
}

type LeagueItem struct {
	IsFreshBlood bool
	IsHotStreak  bool
	IsInactive   bool
	IsVeteran    bool
	LastPlayed   int64
	LeagueName   string
	LeaguePoints int
	Losses       int
	MiniSeries   MiniSeries
}

type MiniSeries struct {
	Losses               int
	Progress             string
	Target               int
	TimeLeftToPlayMillis int64
	Wins                 int
}

func LeagueBySummoner(region string, summonerID int64) (league League, err error) {
	tempMap := make(map[string]League)
	if !IsKeySet() {
		return league, ErrAPIKeyNotSet
	}
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	url := BaseURL + region + "/v2.1/league/by-summoner/" + summonerIDstr
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &tempMap)
	if err != nil {
		return
	}
	return tempMap[summonerIDstr], err
}
