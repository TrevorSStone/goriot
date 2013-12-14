package goriot

import (
	"fmt"
	"testing"
)

const (
	personalkey = "your-key-here"
)

func TestGetChampionsList(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetChampionList(NA, false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetRecentGameBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetRecentGameBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetLeagueBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetLeagueBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetStatSummariesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	stats, err := GetStatSummariesBySummoner(NA, 2112, SEASON3)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(stats)
}

func TestRankedStatsBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	stats, err := GetRankedStatsBySummoner(NA, 2112, SEASON3)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(stats)
}

func TestGetMasteriesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	stats, err := GetMasteriesBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(stats)
}

func TestGetRunesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	stats, err := GetRunesBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(stats)
}

func TestGetSummonerByName(t *testing.T) {
	SetAPIKey(personalkey)
	summoner, err := GetSummonerByName(NA, "manticorex")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(summoner)
}

func TestGetSummonerByID(t *testing.T) {
	SetAPIKey(personalkey)
	summoner, err := GetSummonerByID(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(summoner)
}

func TestGetSummonerNamesByID(t *testing.T) {
	SetAPIKey(personalkey)
	summoners, err := GetSummonerNamesByID(NA, 2112, 1111)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(summoners)
}

func TestGetTeamBySummonerID(t *testing.T) {
	SetAPIKey(personalkey)
	teams, err := GetTeamBySummonerID(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(teams)
}
