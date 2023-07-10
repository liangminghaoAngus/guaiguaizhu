package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/component"
)

type Animation struct {
	query *query.Query
}

func (a *Animation) Update(world donburi.World) {
	a.query.Each(world, func(entry *donburi.Entry) {
		animation := component.Animation.Get(entry)
		if animation.OutOfBound() {
			world.Remove(entry.Entity())
		} else {
			_ = animation.Update()
		}
	})
}

func (a *Animation) Draw(world donburi.World, screen *ebiten.Image) {
	a.query.Each(world, func(entry *donburi.Entry) {
		animation := component.Animation.Get(entry)
		animation.Draw(screen)
	})
}

func NewAnimation() *Animation {
	return &Animation{query: query.NewQuery(filter.Contains(component.Animation))}
}
