package entity

import (
	"bytes"
	"image"
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
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
}

func NewRookieMap(w donburi.World) *donburi.Entry {
	rookieMapEntity := w.Create(RookieMap...)
	rookieMap := w.Entry(rookieMapEntity)

	img, _, _ := image.Decode(bytes.NewReader(assetsImage.MapImage[enums.MapRookie]))
	bg := ebiten.NewImageFromImage(img)
	component.Sprite.SetValue(rookieMap, component.SpriteData{Image: bg})

	return nil
}
