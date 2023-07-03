package entity

import (
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var PlayerEntity = []donburi.IComponentType{
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

	playerEntity := w.Create(PlayerEntity...)
	player := w.Entry(playerEntity)
	playerCollision := resolv.NewObject(20, 20, float64(playerW/2), float64(playerH/2), "player")
	component.Health.SetValue(player, component.NewPlayerHealthData())
	component.Race.SetValue(player, component.NewRaceData(raceInt))
	component.Level.SetValue(player, component.NewLevelData())
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
		Items:     []*resolv.Object{playerCollision},
		TagsOrder: []string{"player"},
	})
	component.Ability.SetValue(player, component.NewAbility(raceInt))

	return player
}
