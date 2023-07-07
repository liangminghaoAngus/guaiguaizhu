package enums

type Map int

const (
	MapRookie Map = iota
	MapFirstout
)

var Maps = []Map{MapRookie}

var MapImages = map[string]Map{
	"rookie_map.png": MapRookie,
}

var MapName = map[Map]string{
	MapRookie: "新手村",
}
