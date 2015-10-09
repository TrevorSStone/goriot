package goriot

import (
	"fmt"
)

//Champion represents a League of Legends champion
type Champion struct {
	Active            bool `json:"active"`
	BotEnabled        bool `json:"botEnabled"`
	BotMmEnabled      bool `json:"botMmEnabled"`
	FreeToPlay        bool `json:"freeToPlay"`
	ID                int  `json:"id"`
	RankedPlayEnabled bool `json:"rankedPlayEnabled"`
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
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.2/champion?%v", region, BaseAPIURL, region, args)
	err = requestAndUnmarshal(url, &champs)
	if err != nil {
		return
	}

	return champs.Champions, err
}

//ChampsionByID retrieves the champion from the Riot Games API.
//It returns a single Champion and any errors that occured from the server
//The global API key must be set before use
func ChampionByID(region string, id int) (champion Champion, err error) {
	if !IsKeySet() {
		return champion, ErrAPIKeyNotSet
	}
	var args string
	args += "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.2/champion/%d?%v", region, BaseAPIURL, region, id, args)
	err = requestAndUnmarshal(url, &champion)
	if err != nil {
		return
	}

	return champion, err
}
