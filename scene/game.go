package scene

import (
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/entity"
	"liangminghaoangus/guaiguaizhu/enums"
	"liangminghaoangus/guaiguaizhu/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
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

	g.initGame(raceInt)
	return g
}

func (g *Game) initGame(raceInt enums.Race) {
	render := system.NewRender()
	// todo append system
	g.systems = []System{
		render,
		system.NewControl(),
	}

	g.drawables = []Drawable{
		render,
	}

	g.world = g.createWorld(raceInt)

}

func (g *Game) createWorld(raceInt enums.Race) donburi.World {
	world := donburi.NewWorld()
	world.Entry(world.Create(component.Game))
	entity.NewPlayer(world, raceInt)
	// todo
	entity.NewRookieMap(world)

	// create base layer
	//playerEntity := entity.NewPlayer(world, raceInt)

	//levelEntry := world.Entry(
	//	world.Create(transform.Transform, component.Sprite),
	//)
	//
	//component.Sprite.SetValue(levelEntry, component.SpriteData{
	//	Image: levelAsset.Background,
	//	Layer: component.SpriteLayerBackground,
	//	Pivot: component.SpritePivotTopLeft,
	//})

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
	screen.Clear()
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
