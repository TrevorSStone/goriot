package goriot

import (
	"errors"
	"fmt"
	"errors"
	"strconv"
)

//Team is a League of Legends Ranked Team
type Team struct {
	CreateDate                    int64                 `json:"createDate"`
	FullID                        string                `json:"fullId"`
	LastGameDate                  int64                 `json:"lastGameDate"`
	LastJoinDate                  int64                 `json:"lastJoinDate"`
	LastJoinedRankedTeamQueueDate int64                 `json:"lastJoinedRankedTeamQueueDate"`
	MatchHistory                  []MatchHistorySummary `json:"matchHistory"`
	ModifyDate                    int64                 `json:"modifyDate"`
	Name                          string                `json:"name"`
	Roster                        Roster                `json:"roster"`
	SecondLastJoinDate            int64                 `json:"secondLastJoinDate"`
	Status                        string                `json:"status"`
	Tag                           string                `json:"tag"`
	TeamStatDetails               []TeamStatDetail      `json:"teamStatDetails"`
	ThirdJoinDate                 int64                 `json:"thirdLastJoinDate"`
}

//MatchHistorySummary is a summary of a matches played by a Team
type MatchHistorySummary struct {
	Assists           int    `json:"assists"`
	Date              int64  `json:"date"`
	Deaths            int    `json:"deaths"`
	GameID            int64  `json:"gameId"`
	GameMode          string `json:"gameMode"`
	Invalid           bool   `json:"invalid"`
	Kills             int    `json:"kills"`
	MapID             int    `json:"mapId"`
	OpposingTeamKills int    `json:"opposingTeamKills"`
	OpposingTeamName  string `json:"opposingTeamName"`
	Win               bool   `json:"win"`
}

//Roster represents the roster of a League of Legends ranked team
type Roster struct {
	MemberList []TeamMemberInfo `json:"memberList"`
	OwnerID    int64            `json:"ownerId"`
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
	Losses             int    `json:"losses"`
	TeamStatType       string `json:"teamStatType"`
	Wins               int    `json:"wins"`
}

//TeamBySummonerID retrieves a summoner's assosiated teams from Riot Games API.
//It returns an array of Team and any errors that occured from the server
//The global API key must be set before use
func TeamBySummonerID(region string, summonerID ...int64) (teams map[int64][]Team, err error) {
	if !IsKeySet() {
		return nil, ErrAPIKeyNotSet
	}

	// Only 10 summoners are currently allowed to be looked up
	if len(summonerID) > 10 {
		return nil, errors.New("TeamBySummoner only allows 10 summoner IDs")
	}

	teams = make(map[int64][]Team)
	preTeams := make(map[string][]Team)

	summonerIdStr, err := createSummonerIDString(summonerID)
	if err != nil {
		return nil, err
	}

	args := "api_key=" + apikey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v2.4/team/by-summoner/%v?%v",
		region, BaseURL, region, summonerIdStr, args)

	err = requestAndUnmarshal(url, &preTeams)
	if err != nil {
		return nil, err
	}

	// Convert summoner IDs to int64
	for k, v := range preTeams {
		id, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil, err
		}

		teams[id] = v
	}

	return teams, nil
}

// Search for Team information using Team ID
// Returns Team match history and information or errors from server
// Global API key must be set prior to use
func TeamByTeamID(region string, teamID ...string) (teams map[string]Team, err error) {

	if !IsKeySet() {
		return nil, ErrAPIKeyNotSet
	}

	teams = make(map[string]Team)

	teamIdStr, err := createTeamIDString(teamID)
	if err != nil {
		return nil, err
	}

	args := "api_key=" + apikey
	url := fmt.Sprintf(
		"https://%v.%v/lol/%v/v2.4/team/%v?%v",
		region,
		BaseURL,
		region,
		teamIdStr,
		args)

	err = requestAndUnmarshal(url, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

// Creates a string of comma seperated Team IDs from
// a slice of Team IDs
// There is currently a hard limit of 10 Teams
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
