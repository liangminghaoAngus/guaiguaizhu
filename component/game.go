package component

import (
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type GameData struct {
	Pause             bool
	PauseKey          ebiten.Key
	SaveGameID        int
	SaveGameKey       [2]ebiten.Key
	SystemUI          *ebiten.Image
	IsPlayerStoreOpen bool

	ConfigData *config.Config
}

func IsNewGame(gameData GameData) bool {
	return gameData.SaveGameID <= 0
}

var Game = donburi.NewComponentType[GameData](GameData{
	PauseKey:    ebiten.KeyEscape,
	SaveGameKey: [2]ebiten.Key{ebiten.KeyControl, ebiten.KeyS},
	ConfigData:  config.GetConfig(),
})

var Map = donburi.NewTag()

var MapActive = donburi.NewTag()

var Player = donburi.NewTag()

var Enemy = donburi.NewTag()

var Npc = donburi.NewTag()

var NotActive = donburi.NewTag()

func MustFindGame(w donburi.World) *GameData {
	game, ok := query.NewQuery(filter.Contains(Game)).First(w)
	if !ok {
		panic("game not found")
	}
	return Game.Get(game)
}

type TeleportData struct {
	ToMap      enums.Map
	ToPosition math.Vec2
}

var Teleport = donburi.NewComponentType[TeleportData](TeleportData{})
