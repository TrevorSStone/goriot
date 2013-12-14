package goriot

import (
	"strconv"
)

type Team struct {
	CreateDate                    int64
	LastGameDate                  int64
	LastJoinDate                  int64
	LastJoinedRankedTeamQueueDate int64
	MatchHistory                  []MatchHistorySummary
	MessageOfDay                  MessageOfDay
	ModifyDate                    int64
	Name                          string
	Roster                        Roster
	SecondLastJoinDate            int64
	Status                        string
	Tag                           string
	TeamID                        TeamID
	TeamStatSummary               TeamStatSummary
	ThirdJoinDate                 int64
	Timestamp                     int64
}

type MatchHistorySummary struct {
	Assists           int
	Date              int64
	Deaths            int
	GameID            int64
	GameMode          string
	Invalid           bool
	Kills             int
	MapID             int
	OpposingTeamKills int
	OpposingTeamName  string
	Win               bool
}

type MessageOfDay struct {
	CreateDate int64
	Message    string
	Version    int
}

type Roster struct {
	MemberList []TeamMemberInfo
	OwnerID    int64
}

type TeamID struct {
	FullID string
}

type TeamStatSummary struct {
	TeamID          TeamID
	TeamStatDetails []TeamStatDetail
}

type TeamMemberInfo struct {
	InviteDate int64
	JoinDate   int64
	PlayerID   int64
	Status     string
}

type TeamStatDetail struct {
	AverageGamesPlayed int
	Losses             int
	MaxRating          int
	Rating             int
	SeedRating         int
	TeamID             TeamID
	TeamStatType       string
	Wins               int
}

func GetTeamBySummonerID(region string, summonerID int64) (team []Team, err error) {
	if !IsKeySet() {
		return team, ErrAPIKeyNotSet
	}
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	url := BaseURL + region + "/v2.1/team/by-summoner/" + summonerIDstr
	args := "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &team)
	if err != nil {
		return
	}
	return
}
