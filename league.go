package goriot

import (
	"fmt"
	"strconv"
)

//League represents a League of Legends league
type League struct {
	Entries []LeagueItem `json:"entries"`
	Name    string       `json:"name"`
	Queue   string       `json:"queue"`
	Tier    string       `json:"tier"`
}

//LeagueItem is an entry in a League. It represents a player or team
type LeagueItem struct {
	IsFreshBlood     bool       `json:"isFreshBlood"`
	IsHotStreak      bool       `json:"isHotStreak"`
	IsInactive       bool       `json:"isInactive"`
	IsVeteran        bool       `json:"isVeteran"`
	LastPlayed       int64      `json:"lastPlayed"`
	LeagueName       string     `json:"leagueName"`
	LeaguePoints     int        `json:"leaguePoints"`
	MiniSeries       MiniSeries `json:"miniSeries"`
	PlayerOrTeamID   string     `json:"playerOrTeamId"`
	PlayerOrTeamName string     `json:"playerOrTeamName"`
	QueueType        string     `json:"queueType"`
	Rank             string     `json:"rank"`
	Tier             string     `json:"tier"`
	Wins             int        `json:"wins"`
}

//MiniSeries shows if a player is in their Series and how far they are within it
type MiniSeries struct {
	Losses               int    `json:"losses"`
	Progress             string `json:"progress"`
	Target               int    `json:"target"`
	TimeLeftToPlayMillis int64  `json:"timeLeftToPlayMillis"`
	Wins                 int    `json:"wins"`
}

//LeagueBySummoner retrieves the league of the supplied summonerID from Riot Games API.
//It returns a League and any errors that occured from the server
//The global API key must be set before use
func LeagueBySummoner(region string, summonerID int64) (league League, err error) {
	tempMap := make(map[string]League)
	if !IsKeySet() {
		return league, ErrAPIKeyNotSet
	}
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v2.2/league/by-summoner/%d?%v", BaseURL, region, summonerID, args)
	err = requestAndUnmarshal(url, &tempMap)
	if err != nil {
		return
	}
	return tempMap[summonerIDstr], err
}
