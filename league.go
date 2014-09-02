package goriot

import (
	"errors"
	"fmt"
	"strconv"
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
	Division         string     `json:"division"`
	IsFreshBlood     bool       `json:"isFreshBlood"`
	IsHotStreak      bool       `json:"isHotStreak"`
	IsInactive       bool       `json:"isInactive"`
	IsVeteran        bool       `json:"isVeteran"`
	LeaguePoints     int        `json:"leaguePoints"`
	MiniSeries       MiniSeries `json:"miniSeries"`
	PlayerOrTeamID   string     `json:"playerOrTeamId"`
	PlayerOrTeamName string     `json:"playerOrTeamName"`
	Wins             int        `json:"wins"`
}

//MiniSeries shows if a player is in their Series and how far they are within it
type MiniSeries struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

//LeagueBySummoner retrieves the league of the supplied summonerID(s) from Riot Games API.
//It returns a League and any errors that occured from the server
//The global API key must be set before use
func LeagueBySummoner(region string, summonerID ...int64) (leagues map[int64][]League, err error) {
	if !IsKeySet() {
		return leagues, ErrAPIKeyNotSet
	}

	// this API method only allows 10 summoners
	if len(summonerID) > 10 {
		return leagues, errors.New("Too many summoners requested. Limit is 10")
	}

	leagues = make(map[int64][]League)
	preLeagues := make(map[string][]League)

	// Turn summoner array into string
	summonerIDstr, err := createSummonerIDString(summonerID)
	if err != nil {
		return leagues, err
	}

	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"https://%v.%v/lol/%v/v2.5/league/by-summoner/%v?%v",
		region,
		BaseURL,
		region,
		summonerIDstr,
		args)
	err = requestAndUnmarshal(url, &preLeagues)
	if err != nil {
		return
	}

	for k, v := range preLeagues {
		id, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil, err
		}

		leagues[id] = v
	}

	return leagues, err
}

//LeagueEntryBySummoner retrieves the league entry of the supplied summonerID.
//The difference from LeagueBySummoner is that LeagueEntryBySummoner only returns data for the given summonerID vs everyone in that summoner's league
func LeagueEntryBySummoner(region string, summonerID ...int64) (leagues map[int64][]League, err error) {
	if !IsKeySet() {
		return leagues, ErrAPIKeyNotSet
	}

	// this API method only allows 10 summoners
	if len(summonerID) > 10 {
		return leagues, errors.New("Too many summoners requested. Limit is 10")
	}

	leagues = make(map[int64][]League)
	preLeagues := make(map[string][]League)

	// Turn summoner array into string
	summonerIDstr, err := createSummonerIDString(summonerID)
	if err != nil {
		return leagues, err
	}

	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"https://%v.%v/lol/%v/v2.5/league/by-summoner/%v/entry?%v",
		region,
		BaseURL,
		region,
		summonerIDstr,
		args)
	err = requestAndUnmarshal(url, &preLeagues)
	if err != nil {
		return nil, err
	}

	for k, v := range preLeagues {
		id, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil, err
		}

		leagues[id] = v
	}

	return leagues, err
}

// LeagueByTeam retrieves all leagues for passed in Team(s)
// It returns a map of TeamID to slice of League and any errors
// The global API key must be set before use
func LeagueByTeam(region string, teamID ...string) (leagues map[string][]League, err error) {

	if !IsKeySet() {
		return leagues, ErrAPIKeyNotSet
	}

	if len(teamID) > 10 {
		return leagues, errors.New("Too many teams requested. Limit is 10")
	}

	leagues = make(map[string][]League)

	teamIDstr, err := createTeamIDString(teamID)
	if err != nil {
		return leagues, err
	}

	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"https://%v.%v/lol/%v/v2.5/league/by-team/%v?%v",
		region,
		BaseURL,
		region,
		teamIDstr,
		args)
	err = requestAndUnmarshal(url, &leagues)
	if err != nil {
		return leagues, err
	}

	return leagues, err
}

// LeagueEntryByTeam retrieves all entries for a Team
// It returns a League and any errors that occured from the server
// The global API key must be set before use
func LeagueEntryByTeam(region string, teamID ...string) (leagues map[string][]League, err error) {

	if !IsKeySet() {
		return leagues, ErrAPIKeyNotSet
	}

	if len(teamID) > 10 {
		return leagues, errors.New("Too many teams requested. Limit is 10")
	}

	leagues = make(map[string][]League)

	teamIDstr, err := createTeamIDString(teamID)
	if err != nil {
		return leagues, err
	}

	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"https://%v.%v/lol/%v/v2.5/league/by-team/%v/entry?%v",
		region,
		BaseURL,
		region,
		teamIDstr,
		args)
	err = requestAndUnmarshal(url, &leagues)
	if err != nil {
		return leagues, err
	}

	return leagues, err
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
		"https://%v.%v/lol/%v/v2.5/league/challenger?%v",
		region,
		BaseURL,
		region,
		args)
	err = requestAndUnmarshal(url, &league)
	if err != nil {
		return
	}
	return league, err
}
