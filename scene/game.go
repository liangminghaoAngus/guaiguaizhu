package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/enums"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Game struct {
	world     donburi.World
	systems   []System
	drawables []Drawable
}

func NewGame(raceInt enums.Race) *Game {
	g := &Game{}

	g.initGame()
	return g
}

func (g *Game) initGame() {

	g.world = g.createWorld()
}

func (g *Game) createWorld() donburi.World {
	world := donburi.NewWorld()
	world.Entry(world.Create(component.Game))

	return world
}

func (g *Game) Update() {
	gameData := component.MustFindGame(g.world)
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		gameData.Pause = !gameData.Pause
	}

	if gameData.Pause {
		ebiten.SetWindowTitle("game pause press Esc continue")
		return
	} else {
		ebiten.SetWindowTitle(config.GetConfig().GameName)
	}

	for _, s := range g.systems {
		s.Update(g.world)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	//text.Draw(screen, "m110's Airplanes", assets.NarrowFont, t.screenWidth/4, 100, color.White)
	//text.Draw(screen, "Player 1: WASD + Space", assets.NarrowFont, t.screenWidth/6, 250, color.White)
	//text.Draw(screen, "Player 2: Arrows + Enter", assets.NarrowFont, t.screenWidth/6, 350, color.White)
	//text.Draw(screen, "Press space to start", assets.NarrowFont, t.screenWidth/5, 500, color.White)
}
