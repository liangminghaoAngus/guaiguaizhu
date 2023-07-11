package entity

import (
	"bytes"
	"fmt"
	"github.com/fishtailstudio/imgo"
	"image"
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/engine"
	"liangminghaoangus/guaiguaizhu/enums"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var EnemyEntity = []donburi.IComponentType{
	component.Enemy,
	component.Box,
	transform.Transform,
	component.Position,
	component.Health,
	component.Level,
	component.Intro,
	component.Movement,
	component.SpriteStand,
	//component.SpriteMovement,
	//component.Collision,
}

func NewEnemyByMap(w donburi.World, parent *donburi.Entry, mapInt enums.Map) []*donburi.Entry {
	res := make([]*donburi.Entry, 0)
	list := data.GetEnemyByMap(mapInt)

	maxEnemyNum := 6
	enemyArea := [2]math.Vec2{{X: 500, Y: 0}, {X: 1000, Y: 0}}
	for _, item := range list {
		l := NewEnemyEntity(w, parent, item.ID, maxEnemyNum, enemyArea)
		res = append(res, l...)
	}

	return res
}

func NewEnemyEntity(w donburi.World, parent *donburi.Entry, enemyID int, num int, area [2]math.Vec2) []*donburi.Entry {
	entitys := make([]*donburi.Entry, num)

	enemyEntitys := w.CreateMany(num, EnemyEntity...)

	// get enemy info from data
	enemyItem := data.GetEnemyByID(enemyID)
	enemyData := enemyItem.TransLate2EnemyItem()

	for ind, entity := range enemyEntitys {
		entry := w.Entry(entity)
		healthParams := basicHealthParams{} // todo

		// random area  just use x now
		x := randomNum(int(area[0].X), int(area[1].X))
		component.Position.SetValue(entry, component.PositionData{
			X:             float64(x),
			Y:             0,
			Map:           enums.Map(enemyData.Map),
			FaceDirection: 0,
		})
		component.Box.SetValue(entry, component.BoxData{
			Width:  enemyData.Box.Width,
			Height: enemyData.Box.Height,
		})

		component.Health.SetValue(entry, newEnemyBasicHealth(healthParams))
		component.Level.SetValue(entry, component.LevelData{
			LevelNum: enemyData.Level,
			Exp:      enemyData.BeatExp,
		})
		component.Intro.SetValue(entry, component.IntroData{
			ID:    fmt.Sprintf("enemy_%d:index_%d", enemyData.ID, ind),
			Type:  0,
			Name:  enemyData.Name,
			Intro: enemyData.Intro,
		})
		component.Movement.SetValue(entry, component.NewMovementData())
		//  根据 id 匹配对应的贴图
		enemyPngFile, _ := assetImages.EnemyImageDir.ReadDir(fmt.Sprintf("enemy/%d", enemyData.ID))
		l1, r1 := make([]*ebiten.Image, 0), make([]*ebiten.Image, 0)
		for _, item := range enemyPngFile {
			raw, _ := assetImages.EnemyImageDir.ReadFile(fmt.Sprintf("enemy/%d/%s", enemyData.ID, item.Name()))
			img, _, _ := image.Decode(bytes.NewReader(raw))
			filpImg := imgo.LoadFromImage(img).Flip(imgo.Horizontal).ToImage()
			r := ebiten.NewImageFromImage(img)
			r1 = append(r1, r)
			box := ebiten.NewImageFromImage(filpImg)
			l1 = append(l1, box)
		}
		component.SpriteStand.SetValue(entry, component.SpriteStandData{
			IsDirectionRight: true,
			Disabled:         false,
			Images:           l1,
			ImagesRight:      r1,
		})

		//	component.SpriteMovement,
		//	component.Collision,

		transform.AppendChild(parent, entry, false)
		entitys[ind] = entry
	}

	return entitys
}

type basicHealthParams struct {
	Hp, Mp, HpMax, MpMax int
	HpPos, MpPos         math.Vec2
	HpUI, MpUI           *ebiten.Image
}

func newEnemyBasicHealth(params basicHealthParams) component.HealthData {
	res := component.HealthData{
		HP:                   params.Hp,
		HPMax:                params.HpMax,
		MP:                   params.Mp,
		MPMax:                params.MpMax,
		HPuiPos:              params.HpPos,
		MPuiPos:              params.MpPos,
		HPui:                 params.HpUI,
		MPui:                 params.MpUI,
		AnimationTime:        time.Second,
		JustDamage:           false,
		DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
		LastAnimationTime:    make(map[component.AnimateHealthItem]time.Time),
		LastAnimationItem:    make(map[component.AnimateHealthItem]int),
	}
	return res
}

func randomNum(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	if b-a == 0 {
		return 0
	}
	return rand.Intn(b-a) + a
}
