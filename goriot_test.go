package goriot

import (
	"fmt"
	"runtime"
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

func TestChampionsList(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := ChampionList(NA, false)
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
    summonerIds := []int64{2112}
	_, err := MasteriesBySummoner(NA, summonerIds)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestRunesBySummoner(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := RunesBySummoner(NA, []int64{2112,36080568})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestSummonerByName(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := SummonerByName(NA, []string{"manticorex"})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestSummonerByID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := SummonerByID(NA, []int64{2112})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestSummonerNamesByID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := SummonerNamesByID(NA, []int64{2112, 1111})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("done")
}

func TestTeamBySummonerID(t *testing.T) {
	SetAPIKey(personalkey)
	_, err := TeamBySummonerID(NA, 2112)
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

//ExampleSetSmallRateLimit shows the default way to set the smaller rate limit for developers
func ExampleSetSmallRateLimit() {
	SetSmallRateLimit(10, 10*time.Second)
}

//ExampleSetLongRateLimit shows the default way to set the larger rate limit for developers
func ExampleSetLongRateLimit() {
	SetLongRateLimit(500, 10*time.Minute)
}
