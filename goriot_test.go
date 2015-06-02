package goriot

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

var (
	personalkey = os.Getenv("RIOTAPIKEY_TEST")
)

func TestSetup(t *testing.T) {
	SetSmallRateLimit(10, 10*time.Second)
	SetLongRateLimit(500, 10*time.Minute)
}

func TestChampionsList(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := ChampionList(NA, false)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestChampionByID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := ChampionByID(NA, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestRecentGameBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := RecentGameBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestLeagueBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := LeagueBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestLeagueByChallenger(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := LeagueByChallenger(NA, RANKED_SOLO_5x5)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestLeagueEntryBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := LeagueEntryBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestLeagueByTeam(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := LeagueByTeam(NA, "TEAM-74daa170-a1b3-11e3-b97f-782bcb4d0bb2")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestLeagueEntryByTeam(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := LeagueEntryByTeam(NA, "TEAM-74daa170-a1b3-11e3-b97f-782bcb4d0bb2")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestStatSummariesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := StatSummariesBySummoner(NA, 2112, SEASON3)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestRankedStatsBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := RankedStatsBySummoner(NA, 2112, SEASON3)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestMasteriesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := MasteriesBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestRunesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := RunesBySummoner(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestSummonerByName(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := SummonerByName(NA, "manticorex")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestSummonerByID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := SummonerByID(NA, 2112)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestSummonerNamesByID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := SummonerNamesByID(NA, 2112, 1111)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestTeamBySummonerID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := TeamBySummonerID(NA, 24199871)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestTeamByTeamID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := TeamByTeamID(NA, "TEAM-9179f610-7a48-11e3-b350-782bcb4d0bb2")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestMatchByMatchIDNoTimeline(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := MatchByMatchID(NA, false, 1495236003)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestMatchByMatchIDWithTimeline(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := MatchByMatchID(NA, true, 1495236003)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestMatchHistoryBySummonerID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := MatchHistoryBySummonerID(
		NA, 24199871, nil, nil, -1, -1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestMatchHistoryBySummonerIDWithFilters(t *testing.T) {
	SetAPIKey(personalkey)
	champions := []int64{11, 45}
	queues := []string{"RANKED_SOLO_5x5", "RANKED_TEAM_3x3", "RANKED_TEAM_5x5"}

	_, err := MatchHistoryBySummonerID(
		NA, 24199871, champions, queues, 1, -1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestRateLimits(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rateChecks := 100
	if smallRateChan.RateQueue == nil {
		SetSmallRateLimit(10, 10*time.Second)
	}
	if longRateChan.RateQueue == nil {
		SetLongRateLimit(500, 10*time.Minute)
	}
	SetAPIKey(personalkey)
	returnchan := make(chan bool)
	for i := 0; i < rateChecks; i++ {
		go func() {
			_, err := ChampionList(NA, false)

			if err != nil {
				fmt.Println(err.Error())
			}
			returnchan <- true
		}()
	}

	for i := 0; i < rateChecks; i++ {
		<-returnchan
		fmt.Println(i)
	}
}

func TestNormalizeSummonerName(t *testing.T) {
	name := "MAnTi Co   RE x"
	name2 := "Ç4 SextacyDragon"
	name3 := "잘 못"
	name = NormalizeSummonerName(name)[0]
	if name != "manticorex" {
		t.Fatalf("SummonerName Not Normallized: %s", name)
	}
	name2 = NormalizeSummonerName(name2)[0]
	if name2 != "ç4sextacydragon" {
		t.Fatalf("SummonerName Not Normallized: %s", name2)
	}
	name3 = NormalizeSummonerName(name3)[0]
	if name3 != "잘못" {
		t.Fatalf("SummonerName Not Normallized: %s", name3)
	}
	names := []string{"MantIc oreX", "Ç4 Sexta cyDragon", "잘 못"}
	NormalizeSummonerName(names...)
	if names[0] != "manticorex" {
		t.Fatalf("SummonerName Not Normallized: %s", name)
	}
	if names[1] != "ç4sextacydragon" {
		t.Fatalf("SummonerName Not Normallized: %s", name2)
	}
	if names[2] != "잘못" {
		t.Fatalf("SummonerName Not Normallized: %s", name3)
	}
}

//ExampleSetSmallRateLimit shows the default way to set the smaller rate limit for developers
func ExampleSetSmallRateLimit() {
	SetSmallRateLimit(10, 10*time.Second)
}

//ExampleSetLongRateLimit shows the default way to set the larger rate limit for developers
func ExampleSetLongRateLimit() {
	SetLongRateLimit(500, 10*time.Minute)
}
