package system

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/component"
	"time"
)

type HeathSystem struct {
	query *query.Query
}

func NewHeath() *HeathSystem {
	return &HeathSystem{
		query: query.NewQuery(filter.And(
			filter.Contains(component.Health, component.Heal),
			filter.Or(filter.Contains(component.Control)),
		)),
	}
}

func (s *HeathSystem) Update(w donburi.World) {

	s.query.Each(w, func(entry *donburi.Entry) {
		health := component.Health.Get(entry)
		heal := component.Heal.Get(entry)
		nowTime := time.Now()
		if entry.HasComponent(component.Control) {
			input := component.Control.Get(entry)

			if inpututil.IsKeyJustPressed(input.HpKey) {
				if num := heal.UseHP(); num > 0 {
					health.ChangeHP(num, nowTime)
				}
			}

			if inpututil.IsKeyJustPressed(input.MpKey) {
				if num := heal.UseMP(); num > 0 {
					health.ChangeMP(num, nowTime)
				}
			}
		}

		// auto heal

		//
	})
}
