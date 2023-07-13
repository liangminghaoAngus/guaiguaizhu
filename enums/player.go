package enums

type Race int

const (
	RaceGod Race = iota
	RaceHuman
	RaceDevil
)

func GetRaceStr(raceInt Race) string {
	res := ""
	switch raceInt {
	case RaceGod:
		res = "god"
	case RaceHuman:
		res = "human"
	case RaceDevil:
		res = "devil"
	}
	return res
}

func GetRaceName(raceInt Race) string {
	res := ""
	switch raceInt {
	case RaceGod:
		res = "神"
	case RaceHuman:
		res = "人"
	case RaceDevil:
		res = "魔"
	}
	return res
}

func GetRaceText(raceInt Race) string {
	res := ""
	switch raceInt {
	case RaceGod:
		res = RaceGodText
	case RaceHuman:
		res = RaceHumanText
	case RaceDevil:
		res = RaceDevilText
	}
	return res
}
