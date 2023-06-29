package enums

type Race int

const (
	RaceGod Race = iota
	RaceHuman
	RaceDevil
)

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
