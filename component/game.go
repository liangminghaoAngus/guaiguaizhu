package component

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type GameData struct {
	Pause    bool
	SaveData []interface{}
}

var Game = donburi.NewComponentType[GameData](GameData{})

var Map = donburi.NewTag()

func MustFindGame(w donburi.World) *GameData {
	game, ok := query.NewQuery(filter.Contains(Game)).First(w)
	if !ok {
		panic("game not found")
	}
	return Game.Get(game)
}
