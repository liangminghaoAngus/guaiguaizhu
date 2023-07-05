package system

import (
	"liangminghaoangus/guaiguaizhu/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		query: query.NewQuery(filter.Contains(component.Player, transform.Transform, component.Position, component.Control, component.Movement)),
	}
}

func (m *Control) Update(w donburi.World) {
	gameData := component.MustFindGame(w)
	m.query.Each(w, func(entry *donburi.Entry) {
		position := component.Position.Get(entry)
		movement := component.Movement.Get(entry)
		input := component.Control.Get(entry)
		isLeftPosition := false

		if ebiten.IsKeyPressed(input.Left) || ebiten.IsKeyPressed(input.Right) {
			movement.VelocityX += movement.AccelerationX
			if movement.VelocityX > movement.MaxSpeed {
				movement.VelocityX = movement.MaxSpeed
			}
		} else {
			movement.VelocityX -= movement.AccelerationX
			if movement.VelocityX < 0 {
				movement.VelocityX = 0
			}
		}

		if ebiten.IsKeyPressed(input.Left) {
			position.X -= movement.VelocityX
			isLeftPosition = true
			// fmt.Println(position)
		} else if ebiten.IsKeyPressed(input.Right) {
			position.X += movement.VelocityX
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
				stand.Disabled = true
			} else {
				stand.IsDirectionRight = !isLeftPosition
				stand.Disabled = false
			}
		}
		// 判断是否存在 spriteMovement 组件
		if entry.HasComponent(component.SpriteMovement) {
			move := component.SpriteMovement.Get(entry)
			move.IsDirectionRight = !isLeftPosition
			if movement.VelocityX > 0 {
				move.Disabled = false
			} else {
				move.Disabled = true
			}
		}

		// storeOpen

		if inpututil.IsKeyJustPressed(input.StoreKey) {
			gameData.IsPlayerStoreOpen = !gameData.IsPlayerStoreOpen
		}
	})
}
