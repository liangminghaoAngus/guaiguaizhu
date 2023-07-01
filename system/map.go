package system

import (
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/enums"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Map struct {
	query *query.Query
}

func NewMap(mapInt enums.Map) *Map {
	return &Map{
		query: query.NewQuery(filter.Contains(transform.Transform, component.Sprite, component.Map)),
	}
}

func (m *Map) Update(w donburi.World) {

}

func (m *Map) Draw(w donburi.World, screen *ebiten.Image) {

	// var entries []*donburi.Entry
	m.query.Each(w, func(entry *donburi.Entry) {
		// entries = append(entries, entry)
		mapData := component.Sprite.Get(entry)
		// 计算尺寸
		img := mapData.Image.Bounds().Size()
		scr := screen.Bounds().Size()
		op := &ebiten.DrawImageOptions{}
		scX, scY := float64(scr.X)/float64(img.X), float64(scr.Y)/float64(img.Y)
		op.GeoM.Scale(scX, scY)

		screen.DrawImage(mapData.Image, op)
	})

	// byLayer := lo.GroupBy(entries, func(entry *donburi.Entry) int {
	// 	return int(component.Sprite.Get(entry).Layer)
	// })
	// layers := lo.Keys(byLayer)
	// sort.Ints(layers)

	// for _, layer := range layers {
	// 	for _, entry := range byLayer[layer] {
	// 		sprite := component.Sprite.Get(entry)

	// 		if sprite.HiddenED {
	// 			continue
	// 		}

	// 		w, h := sprite.Image.Size()
	// 		halfW, halfH := float64(w)/2, float64(h)/2

	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(-halfW, -halfH)
	// 		op.GeoM.Rotate(float64(int(transform.WorldRotation(entry)-sprite.OriginalRotation)%360) * 2 * math.Pi / 360)
	// 		op.GeoM.Translate(halfW, halfH)

	// 		position := transform.WorldPosition(entry)

	// 		x := position.X
	// 		y := position.Y

	// 		switch sprite.Pivot {
	// 		case component.SpritePivotCenter:
	// 			x -= halfW
	// 			y -= halfH
	// 		}

	// 		scale := transform.WorldScale(entry)
	// 		op.GeoM.Translate(-halfW, -halfH)
	// 		op.GeoM.Scale(scale.X, scale.Y)
	// 		op.GeoM.Translate(halfW, halfH)

	// 		//if sprite.ColorOverride != nil {
	// 		//	op.ColorM.Scale(0, 0, 0, sprite.ColorOverride.A)
	// 		//	op.ColorM.Translate(sprite.ColorOverride.R, sprite.ColorOverride.G, sprite.ColorOverride.B, 0)
	// 		//}

	// 		op.GeoM.Translate(x, y)

	// 		r.offscreen.DrawImage(sprite.Image, op)
	// 	}
	// }

	// op := &ebiten.DrawImageOptions{}
	// screen.DrawImage(r.offscreen, op)
}
