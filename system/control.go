package system

import (
	"liangminghaoangus/guaiguaizhu/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Control struct {
	query *query.Query
}

func NewControl() *Control {
	return &Control{
		query: query.NewQuery(filter.Contains(transform.Transform, component.Position, component.Control, component.Movement)),
	}
}

func (m *Control) Update(w donburi.World) {
	gameData := component.MustFindGame(w)
	m.query.Each(w, func(entry *donburi.Entry) {
		position := component.Position.Get(entry)
		movement := component.Movement.Get(entry)
		input := component.Control.Get(entry)
		isLeftPosition := false

		if ebiten.IsKeyPressed(input.Left) {
			position.X -= movement.Speed
			isLeftPosition = true
			// fmt.Println(position)
		} else if ebiten.IsKeyPressed(input.Right) {
			position.X += movement.Speed
			// fmt.Println(position)
		}

		if ebiten.IsKeyPressed(input.EnterKey) {
			print("input.EnterKey")
		}
		if ebiten.IsKeyPressed(input.UeKey) {
			print("input.UeKey")
		}

		// 判断是否存在 spriteStand 组件，根据前进方向修改贴图方向
		if entry.HasComponent(component.SpriteStand) {
			stand := component.SpriteStand.Get(entry)
			// 判断是否进行了移动操作
			if ebiten.IsKeyPressed(input.Left) || ebiten.IsKeyPressed(input.Right) {
				stand.IsDirectionRight = !isLeftPosition
				stand.Disabled = true
			} else {
				stand.Disabled = false
			}
		}
		// 判断是否存在 spriteMovement 组件
		if entry.HasComponent(component.SpriteMovement) {
			move := component.SpriteMovement.Get(entry)
			if ebiten.IsKeyPressed(input.Left) || ebiten.IsKeyPressed(input.Right) {
				move.Disabled = false
				move.IsDirectionRight = !isLeftPosition
			} else {
				move.Disabled = true
			}
		}

		// storeOpen
		if ebiten.IsKeyPressed(input.StoreKey) {
			gameData.IsPlayerStoreOpen = !gameData.IsPlayerStoreOpen
		}
	})
}
