package system

import (
	"fmt"
	"image/color"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/enums"
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Map struct {
	query    *query.Query // 地图底层
	children *query.Query // 地图内的碰撞实体
}

func NewMap(mapInt enums.Map) *Map {
	return &Map{
		query:    query.NewQuery(filter.Contains(transform.Transform, component.Sprite, component.Map)),
		children: query.NewQuery(filter.Contains(transform.Transform, component.Position, component.Collision)),
	}
}

func (m *Map) Update(w donburi.World) {
	// 计算 collision space
	m.children.Each(w, func(e *donburi.Entry) {
		wPos := transform.WorldPosition(e)
		position := component.Position.Get(e)
		c := component.Collision.Get(e)
		for _, item := range c.Items {
			x, y := wPos.X+position.X, wPos.Y+position.Y
			x, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", x), 64)
			y, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", y), 64)
			if happen := item.Check(x, y); happen != nil {
				c := happen.Objects[0]
				// rawX, rawY := x-c.Position.X, y-c.Position.Y
				// engine.SeparateRectangles(item, c)
				if position.X < 0 {
					position.X = math.Abs(float64(c.Position.X + c.Width/2 - item.Width/2))
				} else if position.X > 0 && position.X+item.Width > item.Space.Width {
					position.X = item.Space.Width - item.Width - 10
				} else {
					position.X = math.Abs(float64(c.Position.X + c.Width/2 - item.Width/2))
				}

				// position.Y = math.Abs(c.Position.Y + c.Height/2 - item.Height/2)

				// position.Y = item.Position.Y
				// if rawX > 0 {
				// 	position.X = (c.Position.X - c.Width) - 1
				// } else {
				// 	position.X = (c.Position.X + c.Width) + 1
				// }

				// if rawY > 0 {
				// 	position.Y = (c.Position.Y - c.Height) - 1
				// } else {
				// 	position.Y = (c.Position.Y + c.Height) + 1
				// }
			}
		}

	})

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

		// just debug collision
		if entry.HasComponent(component.CollisionSpace) {
			space := component.CollisionSpace.Get(entry)
			// space := spaceC.Space

			for _, v := range space.Space.Objects {
				d := ebiten.NewImage(int(v.Width), int(v.Height))
				d.Fill(color.White)
				op := ebiten.DrawImageOptions{}
				text.Draw(screen, strings.Join(v.Tags(), ","), config.GetSystemFont(), int(v.Position.X), int(v.Position.Y), color.Black)
				op.GeoM.Translate(v.Position.X, v.Position.Y)
				screen.DrawImage(d, &op)
			}

			// d := ebiten.NewImage(space.Space.Width(), space.Space.Height())
			// d.Fill(color.White)
			// screen.DrawImage(d, &ebiten.DrawImageOptions{})
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
