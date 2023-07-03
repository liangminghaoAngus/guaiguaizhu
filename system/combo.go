package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/component"
)

type Combo struct {
	query *query.Query
}

func NewCombo() *Combo {
	return &Combo{
		query: query.NewQuery(filter.Contains(transform.Transform, component.Position, component.Control, component.Movement)),
	}
}

func (c *Combo) Update(w donburi.World) {

}
