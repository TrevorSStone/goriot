package goriot

type Champion struct {
	Active            bool
	AttackRank        int
	BotEnabled        bool
	BotMmEnabled      bool
	DefenseRank       int
	DifficultyRank    int
	FreeToPlay        bool
	ID                int
	MagicRank         int
	Name              string
	RankedPlayEnabled bool
}

type ChampionList struct {
	Champions []Champion
}

func GetChampionList(region string, freetoplay bool) (champions []Champion, err error) {
	var champs ChampionList
	if !IsKeySet() {
		return champions, ErrAPIKeyNotSet
	}
	url := BaseURL + "lol/" + region + "/v1.1/champion"
	var args string
	if freetoplay {
		args = "freeToPlay=true&"
	}
	args += "api_key=" + apikey
	err = RequestAndUnmarshal(url+"?"+args, &champs)
	if err != nil {
		return
	}

	return champs.Champions, err
}
