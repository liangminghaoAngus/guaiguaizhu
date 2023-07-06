package entity

import (
	"bytes"
	"image"
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/engine"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var PlayerEntity = []donburi.IComponentType{
	component.Player,
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
}

func NewPlayer(w donburi.World, raceInt enums.Race) *donburi.Entry {
	if name := enums.GetRaceName(raceInt); name == "" {
		panic("unknow race")
	}

	playerPositionX, playerPositionY := 20, 20
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

	return player
}

func MustFindPlayerEntry(w donburi.World) *donburi.Entry {
	entry, ok := query.NewQuery(filter.Contains(PlayerEntity...)).First(w)
	if ok {
		return entry
	}
	return nil
}
