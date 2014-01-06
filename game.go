package goriot

import (
	"fmt"
)

//Game represents a League of Legends game
type Game struct {
	ChampionID    int        `json:"championId"`
	CreateDate    int64      `json:"createDate"`
	FellowPlayers []Player   `json:"fellowPlayers"`
	GameID        int64      `json:"gameId"`
	GameMode      string     `json:"gameMode"`
	GameType      string     `json:"gameType"`
	Invalid       bool       `json:"invalid"`
	Level         int        `json:"level"`
	MapID         int        `json:"mapId"`
	Spell1        int        `json:"spell1"`
	Spell2        int        `json:"spell2"`
	Statistics    []GameStat `json:"statistics"`
	SubType       string     `json:"subType"`
	TeamID        int        `json:"teamId"`
}

//Player represents a League of Legends player that was in the requested game
type Player struct {
	ChampionID int   `json:"championId"`
	SummonerID int64 `json:"summonerId"`
	TeamID     int   `json:"teamId"`
}

//GameStat represents a stat for the assosiated Game
type GameStat struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type gamesList struct {
	Games []Game `json:"games"`
}

//RecentGameBySummoner retrieves the current list of recent games from the Riot Games API.
//It returns an array of Game(s) and any errors that occured from the server
//The global API key must be set before use
func RecentGameBySummoner(region string, summonerID int64) (games []Game, err error) {
	var gameslist gamesList
	if !IsKeySet() {
		return games, ErrAPIKeyNotSet
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v1.2/game/by-summoner/%d/recent?%v", BaseURL, region, summonerID, args)
	err = requestAndUnmarshal(url, &gameslist)
	if err != nil {
		return
	}

	return gameslist.Games, err
}
