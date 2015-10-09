package goriot

import (
	"fmt"
)

var (
	RegionToPlatformId = map[string]string{
		"br":   "BR1",
		"eune": "EUN1",
		"euw":  "EUW1",
		"kr":   "KR",
		"lan":  "LA1",
		"las":  "LA2",
		"na":   "NA1",
		"oce":  "OC1",
		"ru":   "RU",
		"tr":   "TR1",
	}
)

type FeaturedGame struct {
	GameID            int           `json:"gameId"`
	MapID             int           `json:"mapId"`
	GameMode          string        `json:"gameMode"`
	GameType          string        `json:"gameType"`
	GameQueueConfigID int           `json:"gameQueueConfigId"`
	Participants      []Participant `json:"participants"`
	Observers         struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	PlatformID      string `json:"platformId"`
	BannedChampions []struct {
		ChampionID int `json:"championId"`
		TeamID     int `json:"teamId"`
		PickTurn   int `json:"pickTurn"`
	} `json:"bannedChampions"`
	GameStartTime int64 `json:"gameStartTime"`
	GameLength    int   `json:"gameLength"`
}

type Games struct {
	ClientRefreshInterval int64          `json:"clientRefreshInterval"`
	GameList              []FeaturedGame `json:"gameList"`
}

// FeaturedGames requests a list of FeaturedGames from Riot
// API key needs to be set prior to use
func FeaturedGames(region string) ([]FeaturedGame, error) {
	if !IsKeySet() {
		return nil, ErrAPIKeyNotSet
	}

	var games Games
	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"https://%v.%v/featured?%v",
		region,
		BaseObserverURL,
		args)

	if err := requestAndUnmarshal(url, &games); err != nil {
		return nil, err
	}

	return games.GameList, nil
}

// FeaturedGameBySummonerID requests a FeaturedGame from Riot
// API key needs to be set prior to use
func FeaturedGameBySummonerID(region, summonerId string) (FeaturedGame, error) {
	if !IsKeySet() {
		return FeaturedGame{}, ErrAPIKeyNotSet
	}

	var game FeaturedGame
	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"https://%v.%v/consumer/getSpectatorGameInfo/%v/%v?%v",
		region,
		BaseObserverURL,
		RegionToPlatformId[region],
		summonerId,
		args,
	)

	if err := requestAndUnmarshal(url, &game); err != nil {
		return FeaturedGame{}, err
	}

	return game, nil
}
