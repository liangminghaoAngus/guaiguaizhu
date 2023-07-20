package enums

type Map int

const (
	MapRookie Map = iota
	MapEastCountryOne
	MapEastCountryTwo
	MapEastCountryThree
	MapEastCountryFour
	MapEastCountryFive
)

var Maps = []Map{MapRookie, MapEastCountryOne, MapEastCountryTwo, MapEastCountryThree, MapEastCountryFour, MapEastCountryFive}

var MapImages = map[string]Map{
	"rookie_map.png": MapRookie,
	"first_out.png":  MapEastCountryOne,
	"east2.png":      MapEastCountryTwo,
	"east3.png":      MapEastCountryThree,
	"east4.png":      MapEastCountryFour,
	"east5.png":      MapEastCountryFive,
}

var MapName = map[Map]string{
	MapRookie:           "新手村",
	MapEastCountryOne:   "乖乖城东郊1",
	MapEastCountryTwo:   "乖乖城东郊2",
	MapEastCountryThree: "乖乖城东郊3",
	MapEastCountryFour:  "乖乖城东郊4",
	MapEastCountryFive:  "乖乖城东郊5",
}

var MapEnemyMax = map[Map]int{
	MapRookie:           0,
	MapEastCountryOne:   6,
	MapEastCountryTwo:   6,
	MapEastCountryThree: 6,
	MapEastCountryFour:  6,
	MapEastCountryFive:  6,
}
