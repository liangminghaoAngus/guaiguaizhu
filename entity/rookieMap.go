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
	space := resolv.NewSpace(spaceW, spaceH, 4, 4)
	component.Sprite.SetValue(rookieMap, component.SpriteData{Image: bg})
	component.CollisionSpace.SetValue(rookieMap, component.CollisionSpaceData{Space: space})

	return nil
}
