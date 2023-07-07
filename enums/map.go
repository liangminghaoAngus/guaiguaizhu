package enums

type Map int

const (
	MapRookie Map = iota
	MapFirstout
)

var Maps = []Map{MapRookie, MapFirstout}

var MapImages = map[string]Map{
	"rookie_map.png": MapRookie,
	"first_out.png":  MapFirstout,
}

var MapName = map[Map]string{
	MapRookie:   "新手村",
	MapFirstout: "城门外",
}
