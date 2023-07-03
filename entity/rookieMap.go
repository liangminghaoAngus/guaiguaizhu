package entity

import (
	"bytes"
	"image"
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

// type RookieMap struct{

// }

var RookieMap = []donburi.IComponentType{
	transform.Transform,
	component.Sprite,
	component.Map,
	// component.CollisionSpace,
}

func NewRookieMap(w donburi.World) *donburi.Entry {
	rookieMapEntity := w.Create(RookieMap...)
	rookieMap := w.Entry(rookieMapEntity)

	// c := config.GetConfig()
	// spaceW, spaceH := c.ScreenWidth, c.ScreenHeight
	img, _, _ := image.Decode(bytes.NewReader(assetsImage.MapImage[enums.MapRookie]))
	bg := ebiten.NewImageFromImage(img)
	// cellSize := 8
	// space := resolv.NewSpace(spaceW*cellSize, spaceH*cellSize, cellSize, cellSize)
	// 制造地图边界
	// top := createMapBound(0, 0, 1280, 1)
	// left := createMapBound(0, 0, float64(cellSize), 640)
	//_ := createMapBound(0, 640-float64(cellSize*2), 1280, 1)
	// right := createMapBound(float64(space.Width()), 0, 16, 640)
	// space.Add(left, right)

	component.Sprite.SetValue(rookieMap, component.SpriteData{Image: bg})
	// component.CollisionSpace.SetValue(rookieMap, component.CollisionSpaceData{Space: space})

	// 放置需要的 npc
	// npcIDs := []int{1, 2, 3}
	// NewNPCs(w, npcIDs)

	return rookieMap
}

func createMapBound(x, y, w, h float64) *resolv.Object {
	return resolv.NewObject(x, y, w, h, "mapBound")
}
