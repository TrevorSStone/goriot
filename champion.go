package goriot

import (
	"fmt"
)

//Champion represents a League of Legends champion
type Champion struct {
	Active            bool   `json:"active"`
	AttackRank        int    `json:"attackRank"`
	BotEnabled        bool   `json:"botEnabled"`
	BotMmEnabled      bool   `json:"botMmEnabled"`
	DefenseRank       int    `json:"defenseRank"`
	DifficultyRank    int    `json:"difficultyRank"`
	FreeToPlay        bool   `json:"freeToPlay"`
	ID                int    `json:"id"`
	MagicRank         int    `json:"magicRank"`
	Name              string `json:"name"`
	RankedPlayEnabled bool   `json:"rankedPlayEnabled"`
}

type championList struct {
	Champions []Champion
}

//ChampionList retrieves the current list of champions from the Riot Games API.
//If freetoplay is set to true, only champions currently free to play are returned
//It returns an array of Champion(s) and any errors that occured from the server
//The global API key must be set before use
func ChampionList(region string, freetoplay bool) (champions []Champion, err error) {
	var champs championList
	if !IsKeySet() {
		return champions, ErrAPIKeyNotSet
	}
	var args string
	if freetoplay {
		args = "freeToPlay=true&"
	}
	args += "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v1.1/champion?%v", BaseURL, region, args)
	err = requestAndUnmarshal(url, &champs)
	if err != nil {
		return
	}

	return champs.Champions, err
}
