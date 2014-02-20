package goriot

import (
	"fmt"
)

//League represents a League of Legends league
type League struct {
	Entries       []LeagueItem `json:"entries"`
	Name          string       `json:"name"`
	ParticipantId string       `json:"participantId"`
	Queue         string       `json:"queue"`
	Tier          string       `json:"tier"`
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
func LeagueBySummoner(region string, summonerID int64) (league []League, err error) {
	if !IsKeySet() {
		return league, ErrAPIKeyNotSet
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v2.3/league/by-summoner/%d?%v", BaseURL, region, summonerID, args)
	err = requestAndUnmarshal(url, &league)
	if err != nil {
		return
	}
	return league, err
}

//LeagueEntryBySummoner retrieves the league entry of the supplied summonerID.
//The difference from LeagueBySummoner is that LeagueEntryBySummoner only returns data for the given summonerID vs everyone in that summoner's league
func LeagueEntryBySummoner(region string, summonerID int64) (entry []LeagueItem, err error) {
	if !IsKeySet() {
		return entry, ErrAPIKeyNotSet
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"%v/lol/%v/v2.3/league/by-summoner/%d/entry?%v",
		BaseURL,
		region,
		summonerID,
		args)
	err = requestAndUnmarshal(url, &entry)
	if err != nil {
		return
	}
	return entry, err
}

//LeagueByChallenger retrieves all the league entries for the Challenger group
//It returns a League and any errors that occured from the server
//The global API key must be set before use
func LeagueByChallenger(region string, queueType string) (league League, err error) {
	if !IsKeySet() {
		return league, ErrAPIKeyNotSet
	}
	args := fmt.Sprintf(
		"type=%v&api_key=%v",
		queueType,
		apikey)
	url := fmt.Sprintf(
		"%v/lol/%v/v2.3/league/challenger?%v",
		BaseURL,
		region,
		args)
	err = requestAndUnmarshal(url, &league)
	if err != nil {
		return
	}
	return league, err
}
