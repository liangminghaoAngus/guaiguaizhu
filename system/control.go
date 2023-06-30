package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/component"
	"sync"
)

type Control struct {
	query *query.Query
}

func NewControl() *Control {
	return &Control{
		query: query.NewQuery(filter.Contains(transform.Transform, component.Position, component.Control, component.Movement)),
	}
}

var once sync.Once

func (m *Control) Update(w donburi.World) {
	m.query.Each(w, func(entry *donburi.Entry) {
		position := component.Position.Get(entry)
		movement := component.Movement.Get(entry)
		input := component.Control.Get(entry)

		if ebiten.IsKeyPressed(input.Left) {
			position.X -= movement.Speed
			fmt.Println(position)
		} else if ebiten.IsKeyPressed(input.Right) {
			position.X += movement.Speed
			fmt.Println(position)
		}

		if ebiten.IsKeyPressed(input.EnterKey) {
			print("input.EnterKey")
		}
		if ebiten.IsKeyPressed(input.UeKey) {
			print("input.UeKey")
		}

	})
}
