package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/component"
)

type BaseGround struct {
	query *query.Query
}

func NewBaseGround() *BaseGround {
	return &BaseGround{
		query: query.NewQuery(filter.Contains(component.Position, transform.Transform, component.Box)),
	}
}

var BaseGroundHeight float64 = 380

func (b *BaseGround) Update(world donburi.World) {
	// calculate position y + height
	b.query.Each(world, func(entry *donburi.Entry) {
		pos := component.Position.Get(entry)
		position := transform.WorldPosition(entry)
		box := component.Box.Get(entry)

		if pos.Y+position.Y+float64(box.Height) < BaseGroundHeight {
			pos.Y += BaseGroundHeight - (pos.Y + position.Y + float64(box.Height))
		} else if pos.Y+position.Y+float64(box.Height) > BaseGroundHeight {
			pos.Y -= (pos.Y + position.Y + float64(box.Height)) - BaseGroundHeight
		}

	})
}
