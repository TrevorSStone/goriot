package goriot

import (
	"fmt"
	"strconv"
)

type PlayerStatsSummaryList struct {
	PlayerStatSummaries []PlayerStatsSummary
	SummonerID          int64
}

type PlayerStatsSummary struct {
	AggregatedStats       []AggregatedStat
	Losses                int
	ModifyDate            int64
	PlayerStatSummaryType string
	Wins                  int
}

type AggregatedStat struct {
	Count int
	ID    int
	Name  string
}

type RankedStats struct {
	Champions  []ChampionStats
	ModifyDate int64
	SummonerID int64
}

type ChampionStats struct {
	ID    int
	Name  string
	Stats []ChampionStat
}

type ChampionStat struct {
	Count int `json:"c"`
	ID    int
	Name  string
	Value int
}

func StatSummariesBySummoner(region string, summonerID int64, season string) (stats []PlayerStatsSummary, err error) {
	var list PlayerStatsSummaryList
	if !IsKeySet() {
		return stats, ErrAPIKeyNotSet
	}
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	url := BaseURL + "lol/" + region + "/v1.1/stats/by-summoner/" + summonerIDstr + "/summary"
	var args string
	if season != "" {
		args = fmt.Sprintf("season=%s&", season)
	}
	args += "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &list)
	if err != nil {
		return
	}
	return list.PlayerStatSummaries, err
}

func RankedStatsBySummoner(region string, summonerID int64, season string) (stats RankedStats, err error) {
	if !IsKeySet() {
		return stats, ErrAPIKeyNotSet
	}
	summonerIDstr := strconv.FormatInt(summonerID, 10)
	url := BaseURL + "lol/" + region + "/v1.1/stats/by-summoner/" + summonerIDstr + "/ranked"
	var args string
	if season != "" {
		args = fmt.Sprintf("season=%s&", season)
	}
	args += "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &stats)
	if err != nil {
		return
	}
	return
}
