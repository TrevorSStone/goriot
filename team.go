package goriot

import (
	"errors"
	"fmt"
)

//Team is a League of Legends Ranked Team
type Team struct {
	CreateDate                    int64                 `json:"createDate"`
	FullID                        string                `json:"fullId"`
	LastGameDate                  int64                 `json:"lastGameDate"`
	LastJoinDate                  int64                 `json:"lastJoinDate"`
	LastJoinedRankedTeamQueueDate int64                 `json:"lastJoinedRankedTeamQueueDate"`
	MatchHistory                  []MatchHistorySummary `json:"matchHistory"`
	MessageOfDay                  MessageOfDay          `json:"messageOfDay"`
	ModifyDate                    int64                 `json:"modifyDate"`
	Name                          string                `json:"name"`
	Roster                        Roster                `json:"roster"`
	SecondLastJoinDate            int64                 `json:"secondLastJoinDate"`
	Status                        string                `json:"status"`
	Tag                           string                `json:"tag"`
	TeamStatSummary               TeamStatSummary       `json:"teamStatSummary"`
	ThirdJoinDate                 int64                 `json:"thirdLastJoinDate"`
}

//MatchHistorySummary is a summary of a matches played by a Team
type MatchHistorySummary struct {
	Assists           int    `json:"assists"`
	Deaths            int    `json:"deaths"`
	GameID            int64  `json:"gameId"`
	GameMode          string `json:"gameMode"`
	Invalid           bool   `json:"inalid"`
	Kills             int    `json:"kills"`
	MapID             int    `json:"mapId"`
	OpposingTeamKills int    `json:"opposingTeamKills"`
	OpposingTeamName  string `json:"opposingTeamName"`
	Win               bool   `json:"win"`
}

//MessageOfDay is a message of the day for a ranked team
type MessageOfDay struct {
	CreateDate int64  `json:"createDate"`
	Message    string `json:"message"`
	Version    int    `json:"version"`
}

//Roster represents the roster of a League of Legends ranked team
type Roster struct {
	MemberList []TeamMemberInfo `json:"memberList"`
	OwnerID    int64            `json:"ownerId"`
}

//TeamStatSummary is a summary of the statistics for a ranked team
type TeamStatSummary struct {
	FullID          string           `json:"fullId"`
	TeamStatDetails []TeamStatDetail `json:"teamStatDetails"`
}

//TeamMemberInfo is the individual information for a player on a ranked team
type TeamMemberInfo struct {
	InviteDate int64  `json:"inviteDate"`
	JoinDate   int64  `json:"joinDate"`
	PlayerID   int64  `json:"playerId"`
	Status     string `json:"status"`
}

//TeamStatDetail is the statistics for a ranked team
type TeamStatDetail struct {
	AverageGamesPlayed int    `json:"averageGamesPlayed"`
	FullID             string `json:"fullId"`
	Losses             int    `json:"losses"`
	TeamStatType       string `json:"teamStatType"`
	Wins               int    `json:"wins"`
}

//TeamBySummonerID retrieves a summoner's assosiated teams from Riot Games API.
//It returns an array of Team and any errors that occured from the server
//The global API key must be set before use
func TeamBySummonerID(region string, summonerID int64) (team []Team, err error) {
	if !IsKeySet() {
		return team, ErrAPIKeyNotSet
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v2.2/team/by-summoner/%d?%v", region, BaseURL, region, summonerID, args)

	err = requestAndUnmarshal(url, &team)
	if err != nil {
		return
	}
	return
}

func createTeamIDString(teamID []string) (teamIDstr string, err error) {

	if len(teamID) > 10 {
		return teamIDstr, errors.New("A Maximum of 10 TeamIDs are allowed")
	}

	for k, v := range teamID {
		teamIDstr += v
		if k != len(teamID)-1 {
			teamIDstr += ","
		}
	}

	return
}
