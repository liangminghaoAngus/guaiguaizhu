package entity

import (
	"bytes"
	"fmt"
	"image"
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/engine"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var GameMap = []donburi.IComponentType{
	transform.Transform,
	component.Intro,
	component.Sprite,
	component.Map,
	component.CollisionSpace,
	component.EnemyMaxCount,
}

func NewGameMap(world donburi.World, parent *donburi.Entry) []*donburi.Entry {
	r := make([]*donburi.Entry, 0)

	for _, mapInt := range enums.Maps {
		//
		npcs := data.GetNpcByMapID(int(mapInt))
		mapEntry := newMapEntry(world, parent, mapInt, npcs)
		r = append(r, mapEntry)
	}

	return r
}

func newMapEntry(w donburi.World, parent *donburi.Entry, mapInt enums.Map, npcIDs []int) *donburi.Entry {
	MapEntity := w.Create(GameMap...)
	Map := w.Entry(MapEntity)

	c := config.GetConfig()
	spaceW, spaceH := float64(c.ScreenWidth), float64(c.ScreenHeight)
	img, _, _ := image.Decode(bytes.NewReader(assetsImage.MapImage[mapInt]))
	bg := ebiten.NewImageFromImage(img)

	space := engine.NewSpace(spaceW, spaceH)
	// 制造地图边界
	top := createMapBound(0, 0, spaceW, 2)
	bot := createMapBound(0, spaceH-2, spaceW, 2)
	left := createMapBound(0, 0, 2, spaceH)
	right := createMapBound(spaceW-2, 0, 2, spaceH)
	space.AddObject(top, bot, left, right)

	component.Sprite.SetValue(Map, component.SpriteData{Image: bg})
	component.CollisionSpace.SetValue(Map, component.CollisionSpaceData{Space: space})
	component.Intro.SetValue(Map, component.IntroData{
		ID:    fmt.Sprintf("map_%d", mapInt),
		Name:  enums.MapName[mapInt],
		Intro: "",
	})
	component.EnemyMaxCount.SetValue(Map, component.EnemyMaxCountData{
		Max: enums.MapEnemyMax[mapInt],
		Cur: 0,
	})

	// 放置需要的 npc
	//npcIDs := []int{1, 2, 3}
	npcs := NewNPCs(w, npcIDs)
	for _, npc := range npcs {
		transform.AppendChild(parent, npc, false)
	}

	return Map
}

func createMapBound(x, y, w, h float64) *engine.Object {
	return engine.NewObject(x, y, w, h, "mapBound")
}
