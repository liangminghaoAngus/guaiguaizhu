package scene

import (
	"bytes"
	"io"
	"liangminghaoangus/guaiguaizhu/assets/sound"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/entity"
	"liangminghaoangus/guaiguaizhu/enums"
	"liangminghaoangus/guaiguaizhu/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
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
	mapRender := system.NewMap(enums.MapRookie)
	// todo append system
	g.systems = []System{
		render,
		system.NewControl(),
		system.NewSound(),
		mapRender,
	}

	g.drawables = []Drawable{
		render,
		mapRender,
	}

	g.world = g.createWorld(raceInt)

}

func (g *Game) createWorld(raceInt enums.Race) donburi.World {
	world := donburi.NewWorld()
	world.Entry(world.Create(component.Game))

	soundEntity := world.Entry(world.Create(component.Sound, component.BgSound))

	// todo need to do switch music
	s, err := wav.DecodeWithoutResampling(bytes.NewReader(sound.Intro))
	if err != nil {
		println("music err")
	}
	audioContext := audio.NewContext(11025)
	m, err := io.ReadAll(s)
	if err != nil {
		println("music err")
	}
	p := audioContext.NewPlayerFromBytes(m)
	component.Sound.SetValue(soundEntity, component.SoundData{
		Loop:         true,
		AudioContext: audioContext,
		AudioPlayer:  p,
		Mp3Byte:      sound.Intro,
		Volume:       10,
	})

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
	bgSound := component.FindBgSound(g.world)
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		gameData.Pause = !gameData.Pause
		bgSound.Paused = !bgSound.Paused
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
