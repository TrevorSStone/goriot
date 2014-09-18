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

	// create a filter for specific champions
	var championIDStr string = intURLParameter(championIDs).String()

	// check to see if specific queues are being filtered
	var rankedQueuesStr string = strURLParameter(rankedQueues).String()

	// check for indexing
	var beginIndexStr string
	if beginIndex <= -1 {
		beginIndexStr = ""
	} else {
		beginIndexStr = strconv.Itoa(beginIndex)
	}

	var endIndexStr string
	if endIndex <= -1 {
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

	// build url string, request, return payload
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
