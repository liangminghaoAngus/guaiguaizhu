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
	component.CollisionSpace,
}

func NewRookieMap(w donburi.World) *donburi.Entry {
	rookieMapEntity := w.Create(RookieMap...)
	rookieMap := w.Entry(rookieMapEntity)

	spaceW, spaceH := 1280, 640
	img, _, _ := image.Decode(bytes.NewReader(assetsImage.MapImage[enums.MapRookie]))
	bg := ebiten.NewImageFromImage(img)
	cellSize := 1
	space := resolv.NewSpace(spaceW, spaceH, cellSize, cellSize)
	// 制造地图边界
	_, left, _, right := createMapBound(0, 0, 1280, 1), createMapBound(0, 0, 1, 640), createMapBound(0, 640-float64(cellSize*2), 1280, 1), createMapBound(1280-float64(cellSize*2), 0, 1, 640)
	space.Add(left, right)

	component.Sprite.SetValue(rookieMap, component.SpriteData{Image: bg})
	component.CollisionSpace.SetValue(rookieMap, component.CollisionSpaceData{Space: space})

	return rookieMap
}

func createMapBound(x, y, w, h float64) *resolv.Object {
	return resolv.NewObject(x, y, w, h, "mapBound")
}
