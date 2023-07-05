package scene

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"io"
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/assets/sound"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/entity"
	"liangminghaoangus/guaiguaizhu/enums"
	"liangminghaoangus/guaiguaizhu/system"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
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

	g.systems = []System{
		render,
		system.NewControl(),
		system.NewSound(),
		system.NewHeath(),
		mapRender,
	}

	g.drawables = []Drawable{
		mapRender,
		render,
	}

	g.world = g.createWorld(raceInt)

}

func (g *Game) createWorld(raceInt enums.Race) donburi.World {
	world := donburi.NewWorld()
	parent := world.Entry(world.Create(component.Game, transform.Transform))
	transform.SetWorldPosition(parent, math.Vec2{X: 0, Y: 300})

	soundEntity := world.Entry(world.Create(component.Sound, component.BgSound))

	sampleRate := 11025
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
		Total:        time.Second * time.Duration(s.Length()) / time.Duration(sampleRate),
		AudioContext: audioContext,
		AudioPlayer:  p,
		Mp3Byte:      sound.Intro,
		Volume:       10,
	})

	systemUI, _, _ := image.Decode(bytes.NewReader(assetImages.SystemUI))
	// 添加系统 ui
	gameC := component.Game.Get(parent)
	gameC.SystemUI = ebiten.NewImageFromImage(systemUI)

	rookieMap := entity.NewRookieMap(world)
	// entity.NewRookieMap(world)
	player := entity.NewPlayer(world, raceInt)
	transform.AppendChild(parent, player, false)

	// 将 player 添加至 rookie map bound
	rSpace := component.CollisionSpace.Get(rookieMap)
	pCollision := component.Collision.Get(player)
	rSpace.Space.AddObject(pCollision.Items...)

	return world
}

func (g *Game) Update() {
	gameData := component.MustFindGame(g.world)
	bgSound := component.FindBgSound(g.world)
	playForSave := entity.MustFindPlayerEntry(g.world)

	// pause game
	if inpututil.IsKeyJustPressed(gameData.PauseKey) {
		gameData.Pause = !gameData.Pause
		bgSound.AudioPlayer.Pause()
	} else {
		if !bgSound.AudioPlayer.IsPlaying() {
			_ = bgSound.AudioPlayer.Rewind()
			bgSound.AudioPlayer.Play()
		}
	}

	// save game
	if inpututil.IsKeyJustPressed(gameData.SaveGameKey[0]) && inpututil.IsKeyJustPressed(gameData.SaveGameKey[1]) {
		saveMap := make(map[string]string)
		if t, err := marshalComponentData(component.Player.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["player"] = t
		}

		if t, err := marshalComponentData(transform.Transform.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["transform"] = t
		}

		if t, err := marshalComponentData(component.Health.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["health"] = t
		}

		if t, err := marshalComponentData(component.Race.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["race"] = t
		}

		if t, err := marshalComponentData(component.Level.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["level"] = t
		}

		if t, err := marshalComponentData(component.Ability.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["ability"] = t
		}

		if t, err := marshalComponentData(component.Movement.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["movement"] = t
		}

		if t, err := marshalComponentData(component.Position.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["position"] = t
		}

		if t, err := marshalComponentData(component.Store.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["store"] = t
		}
		saveMapRaw, _ := json.Marshal(saveMap)

		var err error
		// todo 暂时不做额外存档处理
		if component.IsNewGame(*gameData) {
			id, err := data.SavePlayerData2Local(string(saveMapRaw))
			if err == nil {
				gameData.SaveGameID = id
			}
		} else {
			err = data.SaveGameChangeOrigin(gameData.SaveGameID, string(saveMapRaw))
		}

		if err != nil {
			fmt.Printf("save game db err :%+v\n", err)
		}

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

func marshalComponentData(componentData interface{}) (string, error) {
	raw, err := json.Marshal(componentData)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}
