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
	mapRender := system.NewMap()
	animateScene := system.NewAnimation()

	g.systems = []System{
		render,
		system.NewControl(),
		system.NewSound(),
		system.NewHeath(),
		mapRender,
		animateScene,
	}

	g.drawables = []Drawable{
		mapRender,
		render,
		animateScene,
	}

	g.world = g.createWorld(raceInt)

}

func (g *Game) createWorld(raceInt enums.Race) donburi.World {
	world := donburi.NewWorld()
	parent := world.Entry(world.Create(entity.GameEntity...))
	transform.SetWorldPosition(parent, math.Vec2{X: 0, Y: 300})

	cfg := config.GetConfig()
	ot := ebiten.NewImage(cfg.ScreenWidth, cfg.ScreenHeight)
	component.Animation.SetValue(parent, component.NewAnimation(ot, cfg.ScreenWidth, cfg.ScreenHeight))

	soundEntity := world.Entry(world.Create(component.Sound, component.BgSound))

	soundData := ChangeMusic("body")

	component.Sound.SetValue(soundEntity, *soundData)

	systemUI, _, _ := image.Decode(bytes.NewReader(assetImages.SystemUI))
	// 添加系统 ui
	gameC := component.Game.Get(parent)
	gameC.SystemUI = ebiten.NewImageFromImage(systemUI)

	mapsEntry := entity.NewGameMap(world, parent)

	player := entity.NewPlayer(world, raceInt)
	transform.AppendChild(parent, player, false)

	teleports := entity.NewTeleports(world)
	for _, teleport := range teleports {
		transform.AppendChild(parent, teleport, false)
	}

	// 将 player 添加至 map bound
	if len(mapsEntry) > 0 {
		rSpace := component.CollisionSpace.Get(mapsEntry[0])
		pCollision := component.Collision.Get(player)
		rSpace.Space.AddObject(pCollision.Items...)
	}

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
		if !gameData.Pause && !bgSound.AudioPlayer.IsPlaying() {
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

		if t, err := marshalComponentData(component.Attribute.Get(playForSave)); err != nil {
			fmt.Printf("save data err:%+v\n", err)
		} else {
			saveMap["attribute"] = t
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

func ChangeMusic(screen string) *component.SoundData {
	mp3Byte := []byte("")
	switch screen {
	case "body":
		mp3Byte = sound.Body
	case "intro":
		mp3Byte = sound.Intro
	case "boss":
		mp3Byte = sound.Boss
	}
	sampleRate := 11025
	s, err := wav.DecodeWithoutResampling(bytes.NewReader(mp3Byte))
	if err != nil {
		println("music err")
	}
	audioContext := audio.NewContext(sampleRate)
	m, err := io.ReadAll(s)
	if err != nil {
		println("music err")
	}
	p := audioContext.NewPlayerFromBytes(m)
	return &component.SoundData{
		Loop:         true,
		AudioContext: audioContext,
		AudioPlayer:  p,
		//Mp3Byte:      mp3Byte,
		Volume: 10,
	}
}
