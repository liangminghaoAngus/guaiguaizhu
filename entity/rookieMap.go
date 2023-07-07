package entity

import (
	"bytes"
	"image"
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/engine"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

// type RookieMap struct{

// }

var RookieMap = []donburi.IComponentType{
	transform.Transform,
	component.Sprite,
	component.Map,
	component.CollisionSpace,
}

func NewRookieMap(w donburi.World, parent *donburi.Entry) *donburi.Entry {
	rookieMapEntity := w.Create(RookieMap...)
	rookieMap := w.Entry(rookieMapEntity)

	c := config.GetConfig()
	spaceW, spaceH := float64(c.ScreenWidth), float64(c.ScreenHeight)
	img, _, _ := image.Decode(bytes.NewReader(assetsImage.MapImage[enums.MapRookie]))
	bg := ebiten.NewImageFromImage(img)

	space := engine.NewSpace(spaceW, spaceH)
	// 制造地图边界
	top := createMapBound(0, 0, spaceW, 2)
	bot := createMapBound(0, spaceH-2, spaceW, 2)
	left := createMapBound(0, 0, 2, spaceH)
	right := createMapBound(spaceW-2, 0, 2, spaceH)
	space.AddObject(top, bot, left, right)

	component.Sprite.SetValue(rookieMap, component.SpriteData{Image: bg})
	component.CollisionSpace.SetValue(rookieMap, component.CollisionSpaceData{Space: space})

	// 放置需要的 npc
	npcIDs := []int{1, 2, 3}
	npcs := NewNPCs(w, npcIDs)
	for _, npc := range npcs {
		transform.AppendChild(parent, npc, false)
	}

	return rookieMap
}

func createMapBound(x, y, w, h float64) *engine.Object {
	return engine.NewObject(x, y, w, h, "mapBound")
}
