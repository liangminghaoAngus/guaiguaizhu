package entity

import (
	"bytes"
	"fmt"
	"image"
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/engine"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

var PlayerEntity = []donburi.IComponentType{
	component.Player,
	component.Box,
	component.Attribute,
	transform.Transform,
	component.Health,
	component.Heal,
	component.Race,
	component.Level,
	component.Ability,
	component.Movement,
	component.Position,
	component.SpriteStand,
	component.SpriteMovement,
	component.Control,
	component.Collision,
	component.Store,
	component.WeaponHandler,
}

func NewPlayer(w donburi.World, raceInt enums.Race) *donburi.Entry {
	if name := enums.GetRaceName(raceInt); name == "" {
		panic("unknow race")
	}

	playerPositionX, playerPositionY := 20, 20

	playerLevel := 1

	standImages := make([]*ebiten.Image, 0)
	standImagesLeft := make([]*ebiten.Image, 0)
	movementLeftImages := make([]*ebiten.Image, 0)
	movementRightImages := make([]*ebiten.Image, 0)
	switch raceInt {
	case enums.RaceGod:
	case enums.RaceHuman:
		standImages = assetImages.HumanStandImgs
		standImagesLeft = assetImages.HumanStandImgsLeft
		movementLeftImages = assetImages.HumanMovementLeftImgs
		movementRightImages = assetImages.HumanMovementRightImgs
	case enums.RaceDevil:
	}

	HPuiImage, _, _ := image.Decode(bytes.NewReader(assetImages.SystemHP))
	MPuiImage, _, _ := image.Decode(bytes.NewReader(assetImages.SystemMP))
	hp := ebiten.NewImageFromImage(HPuiImage)
	mp := ebiten.NewImageFromImage(MPuiImage)

	playerEntity := w.Create(PlayerEntity...)
	player := w.Entry(playerEntity)

	component.Box.SetValue(player, component.NewPlayerBox(raceInt))
	box := component.Box.Get(player)
	playerH, playerW := box.Height, box.Width

	playerCollision := engine.NewObject(float64(playerPositionX), float64(playerPositionY), float64(playerW), float64(playerH), "player")
	component.Health.SetValue(player, component.NewPlayerHealthData(hp, mp))
	component.Race.SetValue(player, component.NewRaceData(raceInt))
	component.Level.SetValue(player, component.NewLevelData(playerLevel))
	//component.Movement.SetValue(player, component.NewMovementData())
	component.Position.SetValue(player, component.NewPlayerPositionData())
	component.SpriteStand.SetValue(player, component.SpriteStandData{
		IsDirectionRight: true,
		Disabled:         false,
		Images:           standImages,
		ImagesRight:      standImagesLeft,
	})
	component.SpriteMovement.SetValue(player, component.SpriteMovementData{
		IsDirectionRight: true,
		Disabled:         true,
		LeftImages:       movementLeftImages,
		RightImages:      movementRightImages,
	})
	component.Control.SetValue(player, component.NewPlayerControl())
	component.Collision.SetValue(player, component.CollisionData{
		Debug:     true,
		Items:     []*engine.Object{playerCollision},
		TagsOrder: []string{"player"},
	})
	component.Ability.SetValue(player, component.NewAbility(raceInt))
	store := component.MustFindStore(w)
	store.DrawUI()

	armer := ebiten.NewImage(20, 20)
	// armer.Fill(color.White)
	humanHand, _ := assetImages.WeaponDir.ReadFile("weapon/666.png")
	humanHandImg, _, _ := image.Decode(bytes.NewReader(humanHand))
	hImg := ebiten.NewImageFromImage(humanHandImg)
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(0.6, 0.6)
	armer.DrawImage(hImg, ops)

	// todo test weapon
	weaponData := data.GetWeaponByID(1)
	wi, _ := assetImages.WeaponDir.ReadFile(fmt.Sprintf("weapon/%s", weaponData.Image))
	weaponImg, _, _ := image.Decode(bytes.NewReader(wi))

	component.WeaponHandler.SetValue(player, component.WeaponHandlerData{
		Image:       armer,
		WeaponPoint: math.NewVec2(20, 20),
		Weapon: &component.WeaponData{
			Image:  ebiten.NewImageFromImage(weaponImg),
			Width:  80,
			Height: 33,
		},
		Point:  math.NewVec2(20, 34),
		Width:  float64(armer.Bounds().Dx()),
		Height: float64(armer.Bounds().Dy()),
	})

	return player
}

func MustFindPlayerEntry(w donburi.World) *donburi.Entry {
	entry, ok := query.NewQuery(filter.Contains(PlayerEntity...)).First(w)
	if ok {
		return entry
	}
	return nil
}
