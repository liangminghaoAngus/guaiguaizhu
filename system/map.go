package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"liangminghaoangus/guaiguaizhu/entity"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Map struct {
	query    *query.Query // 地图底层
	children *query.Query // 地图内的碰撞实体
}

func NewMap() *Map {
	return &Map{
		query: query.NewQuery(
			filter.Contains(transform.Transform, component.Sprite, component.Map, component.MapActive)),
		children: query.NewQuery(filter.Contains(transform.Transform, component.Position, component.Collision)),
	}
}

func (m *Map) Update(w donburi.World) {

	// 判断角色所在的地图
	player := entity.MustFindPlayerEntry(w)
	playerPos := component.Position.Get(player)
	playTransform, ok := transform.GetParent(player)
	if !ok {
		return
	}

	// 设置切换场景地图
	query.NewQuery(filter.Contains(transform.Transform, component.Sprite, component.Intro, component.Map)).Each(w, func(entry *donburi.Entry) {
		intro := component.Intro.Get(entry)
		enemyMaxCount := component.EnemyMaxCount.Get(entry)
		if intro.ID == fmt.Sprintf("map_%d", playerPos.Map) {
			entry.AddComponent(component.MapActive)
			if enemyMaxCount.Cur < enemyMaxCount.Max {
				// 添加当前地图的怪物
				e := entity.NewEnemyByMap(w, playTransform, playerPos.Map)
				enemyMaxCount.Cur += len(e)
			}
		} else {
			entry.RemoveComponent(component.MapActive)
		}
	})

	// 不在当前地图的怪物移除
	component.Enemy.Each(w, func(entry *donburi.Entry) {
		p := component.Position.Get(entry)
		if p.Map != playerPos.Map {
			w.Remove(entry.Entity())
		}
	})

	component.Npc.Each(w, func(e *donburi.Entry) {
		position := component.Position.Get(e)
		if position.Map != playerPos.Map {
			e.AddComponent(component.NotActive)
		} else {
			e.RemoveComponent(component.NotActive)
		}
	})

	component.Teleport.Each(w, func(e *donburi.Entry) {
		position := component.Position.Get(e)
		if position.Map != playerPos.Map {
			e.AddComponent(component.NotActive)
		} else {
			e.RemoveComponent(component.NotActive)
		}
	})

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
				if position.X < 0 {
					position.X = math.Abs(float64(c.Position.X + c.Width/2 - item.Width/2))
				} else if position.X > 0 && position.X+item.Width > item.Space.Width {
					position.X = item.Space.Width - item.Width - 10
				} else {
					position.X = math.Abs(float64(c.Position.X + c.Width/2 - item.Width/2))
				}
			}
		}

	})

}

func (m *Map) Draw(w donburi.World, screen *ebiten.Image) {

	entry, ok := m.query.First(w)
	if !ok {
		return
	}
	mapData := component.Sprite.Get(entry)
	mapInfo := component.Intro.Get(entry)
	img := mapData.Image.Bounds().Size()
	scr := screen.Bounds().Size()
	op := &ebiten.DrawImageOptions{}
	scX, scY := float64(scr.X)/float64(img.X), float64(scr.Y)/float64(img.Y)
	op.GeoM.Scale(scX, scY)

	screen.DrawImage(mapData.Image, op)

	// 绘制地图名称
	fontSize := config.GetSystemFont()
	nameText := fmt.Sprintf("当前地图:%s", mapInfo.Name)
	nameTextBound := text.BoundString(fontSize, nameText)
	text.Draw(screen, nameText, config.GetSystemFont(), 0, screen.Bounds().Dy()-nameTextBound.Bounds().Dy(), color.White)

	//m.query.Each(w, func(entry *donburi.Entry) {
	//	// entries = append(entries, entry)
	//	mapData := component.Sprite.Get(entry)
	//	// 计算尺寸
	//	img := mapData.Image.Bounds().Size()
	//	scr := screen.Bounds().Size()
	//	op := &ebiten.DrawImageOptions{}
	//	scX, scY := float64(scr.X)/float64(img.X), float64(scr.Y)/float64(img.Y)
	//	op.GeoM.Scale(scX, scY)
	//
	//	screen.DrawImage(mapData.Image, op)
	//
	//	// just debug collision
	//	//if entry.HasComponent(component.CollisionSpace) {
	//	//space := component.CollisionSpace.Get(entry)
	//	// space := spaceC.Space
	//
	//	//for _, v := range space.Space.Objects {
	//	//	d := ebiten.NewImage(int(v.Width), int(v.Height))
	//	//	d.Fill(color.White)
	//	//	op := ebiten.DrawImageOptions{}
	//	//	text.Draw(screen, strings.Join(v.Tags(), ","), config.GetSystemFont(), int(v.Position.X), int(v.Position.Y), color.Black)
	//	//	op.GeoM.Translate(v.Position.X, v.Position.Y)
	//	//	screen.DrawImage(d, &op)
	//	//}
	//
	//	// d := ebiten.NewImage(space.Space.Width(), space.Space.Height())
	//	// d.Fill(color.White)
	//	// screen.DrawImage(d, &ebiten.DrawImageOptions{})
	//	//}
	//
	//})

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
