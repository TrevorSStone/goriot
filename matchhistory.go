package goriot

import (
	"fmt"
	"strconv"
)

type PlayerHistory struct {
	Matches []MatchSummary `json:"matches"`
}

type MatchSummary struct {
	MapID                 int                   `json:"mapId"`
	MatchCreation         int64                 `json:"matchCreation"`
	MatchDuration         int64                 `json:"matchDuration"`
	MatchVersion          string                `json:"matchVersion"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	QueueType             string                `json:"queueType"`
	Region                string                `json:"region"`
	Season                string                `json:"season"`
}

func MatchHistoryBySummonerID(
	region string,
	summonerID int64,
	championIDs []int64,
	rankedQueues []string,
	beginIndex int,
	endIndex int) (
	playerHistory PlayerHistory, err error) {

	if !IsKeySet() {
		return playerHistory, ErrAPIKeyNotSet
	}

	// Check to see if champions are being filtered.
	// There is 40 champion restriction when using
	// createSummonerIDString. This restriction does
	// not exist in the API
	var championIDStr string
	if championIDs == nil {
		championIDStr = ""
	} else {
		championIDStr, err = createSummonerIDString(championIDs)
		if err != nil {
			championIDStr = ""
		}
	}

	// check to see if specific queues are being filtered
	var rankedQueuesStr string
	if rankedQueues == nil {
		rankedQueuesStr = ""
	} else {
		rankedQueuesStr, err = createTeamIDString(rankedQueues)
		if err != nil {
			rankedQueuesStr = ""
		}
	}

	// check for indexing
	var beginIndexStr string
	if beginIndex == -1 {
		beginIndexStr = ""
	} else {
		beginIndexStr = strconv.Itoa(beginIndex)
	}

	var endIndexStr string
	if endIndex == -1 {
		endIndexStr = ""
	} else {
		endIndexStr = strconv.Itoa(endIndex)
	}

	// build argument string
	args := fmt.Sprintf(
		"api_key=%v&championIds=%v&rankedQueues=%v&beginIndex=%v&endIndex=%v",
		apikey,
		championIDStr,
		rankedQueuesStr,
		beginIndexStr,
		endIndexStr)

	url := fmt.Sprintf(
		"https://%v.%v/lol/%v/v2.2/matchhistory/%d?%v",
		region,
		BaseURL,
		region,
		summonerID,
		args)
	err = requestAndUnmarshal(url, &playerHistory)
	if err != nil {
		return playerHistory, err
	}

	return playerHistory, nil

}
