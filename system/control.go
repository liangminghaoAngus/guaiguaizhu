package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/entity"
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
			if !position.IsFaceLeft() {
				position.ChangeFaceDirection()
			}
			position.X -= movement.VelocityX
		} else if ebiten.IsKeyPressed(input.Right) {
			if position.IsFaceLeft() {
				position.ChangeFaceDirection()
			}
			position.X += movement.VelocityX
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
			if movement.VelocityX > 0 {
				stand.Disabled = true
			} else {
				stand.IsDirectionRight = !position.IsFaceLeft()
				stand.Disabled = false
			}
		}
		// 判断是否存在 spriteMovement 组件
		if entry.HasComponent(component.SpriteMovement) {
			move := component.SpriteMovement.Get(entry)
			move.IsDirectionRight = !position.IsFaceLeft()
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

		// teleport
		if inpututil.IsKeyJustPressed(input.TeleportKey) {
			if teleport := entity.InTeleport(w, entry, position.Map); teleport != nil {
				// 传送 entity 修改 entity 的 position
				// 切换地图
				toPos := component.Teleport.Get(teleport)
				pos := component.Position.Get(entry)
				pos.Map = toPos.ToMap
				pos.X = toPos.ToPosition.X
				pos.Y = toPos.ToPosition.Y
			}
		}
	})
}
