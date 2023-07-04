package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/config"
)

type GameData struct {
	Pause             bool
	SystemUI          *ebiten.Image
	IsPlayerStoreOpen bool
	SaveData          []interface{}
	ConfigData        config.Config
}

var Game = donburi.NewComponentType[GameData](GameData{})

var Map = donburi.NewTag()

var Player = donburi.NewTag()

func MustFindGame(w donburi.World) *GameData {
	game, ok := query.NewQuery(filter.Contains(Game)).First(w)
	if !ok {
		panic("game not found")
	}
	return Game.Get(game)
}
