package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/component"
)

type HeathSystem struct {
	query *query.Query
}

func NewHeathSystem() *HeathSystem {
	return &HeathSystem{
		query: query.NewQuery(filter.And(
			filter.Contains(component.Health, component.Heal),
			filter.Or(filter.Contains(component.Control)),
		)),
	}
}

func (s *HeathSystem) Update(w donburi.World) {

}
