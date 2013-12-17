package goriot

import (
	"strconv"
)

type Game struct {
	ChampionID    int
	CreateDate    int64
	FellowPlayers []Player
	GameID        int64
	GameMode      string
	GameType      string
	Invalid       bool
	Level         int
	MapID         int
	Spell1        int
	Spell2        int
	Statistics    []GameStat
	SubType       string
	TeamID        int
}

type Player struct {
	ChampionID int
	SummonerID int64
	TeamID     int
}

type GameStat struct {
	ID    int
	Name  string
	Value int
}

type GamesList struct {
	Games []Game
}

func RecentGameBySummoner(region string, summonerID int64) (games []Game, err error) {
	var gameslist GamesList
	if !IsKeySet() {
		return games, ErrAPIKeyNotSet
	}
	url := BaseURL + "lol/" + region + "/v1.1/game/by-summoner/" + strconv.FormatInt(summonerID, 10) + "/recent"
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &gameslist)
	if err != nil {
		return
	}

	return gameslist.Games, err
}
