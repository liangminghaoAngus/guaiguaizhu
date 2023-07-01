package system

import (
	"liangminghaoangus/guaiguaizhu/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Render struct {
	query     *query.Query
	offscreen *ebiten.Image
}

func NewRender() *Render {
	r := &Render{
		query: query.NewQuery(
			filter.And(
				filter.Contains(transform.Transform),
				filter.Or(filter.Contains(component.Sprite), filter.Contains(component.SpriteStand)),
				filter.Not(filter.Contains(component.Map)))),
		offscreen: ebiten.NewImage(3000, 3000),
	}
	return r
}

func (r *Render) Update(w donburi.World) {

	// 修改 sprite 渲染
	r.query.Each(w, func(entry *donburi.Entry) {
		// 判断是否实体存在 spriteStand 贴图和 position 数据
		if entry.HasComponent(component.SpriteStand) {
			// position := component.Position.Get(entry)
			standImages := component.SpriteStand.Get(entry)
			if standImages.Disabled {
				return
			}
			index := (standImages.Count / 5) % 8
			if index > len(standImages.Images)-1 {
				standImages.Count = 0
				index = 0
			}
			standImages.Count++
		}

	})

}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {

	var entries []*donburi.Entry
	r.query.Each(w, func(entry *donburi.Entry) {
		entries = append(entries, entry)
		pos := transform.WorldPosition(entry)

		if entry.HasComponent(component.SpriteStand) && entry.HasComponent(component.Position) {
			position := component.Position.Get(entry)
			standImages := component.SpriteStand.Get(entry)
			if standImages.Disabled {
				return
			}
			index := (standImages.Count / 5) % 8
			// 判断是否需要翻转贴图方向
			targetImage := &ebiten.Image{}
			if standImages.IsDirectionRight {
				targetImage = standImages.Images[index]
			} else {
				targetImage = standImages.ImagesRight[index]
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(position.X+pos.X, position.Y+pos.Y)
			screen.DrawImage(targetImage, op)
		}

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
