package goriot

import (
	"testing"
	"time"
)

const (
	personalkey = "your-key-here"
)

func TestSetup(t *testing.T) {
	SetSmallRateLimit(10, 10*time.Second)
	SetLongRateLimit(500, 10*time.Minute)
}

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
	_, err := GetStatSummariesBySummoner(NA, 2112, SEASON3)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRankedStatsBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetRankedStatsBySummoner(NA, 2112, SEASON3)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetMasteriesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetMasteriesBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetRunesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetRunesBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetSummonerByName(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetSummonerByName(NA, "manticorex")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetSummonerByID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetSummonerByID(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetSummonerNamesByID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetSummonerNamesByID(NA, 2112, 1111)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetTeamBySummonerID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := GetTeamBySummonerID(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRateLimits(t *testing.T) {
	if smallRateChan == nil {
		SetSmallRateLimit(10, 10*time.Second)
	}
	if longRateChan == nil {
		SetLongRateLimit(500, 10*time.Minute)
	}
	SetAPIKey(personalkey)
	for i := 0; i < 24; i++ {
		_, err := GetChampionList(NA, false)
		if err != nil {
			t.Error(err.Error())
		}
	}

}
