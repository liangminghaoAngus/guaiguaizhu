package entity

import (
	"bytes"
	"image"
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/engine"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var PlayerEntity = []donburi.IComponentType{
	component.Player,
	transform.Transform,
	component.Health,
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
}

func NewPlayer(w donburi.World, raceInt enums.Race) *donburi.Entry {
	if name := enums.GetRaceName(raceInt); name == "" {
		panic("unknow race")
	}
	// todo 设计 player 的模型
	playerH, playerW := 80, 50
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
	playerCollision := engine.NewObject(20, 20, float64(playerW), float64(playerH), "player")
	component.Health.SetValue(player, component.NewPlayerHealthData(hp, mp))
	component.Race.SetValue(player, component.NewRaceData(raceInt))
	component.Level.SetValue(player, component.NewLevelData(playerLevel))
	component.Movement.SetValue(player, component.NewMovementData())
	component.Position.SetValue(player, component.NewPositionData())
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

	return player
}
