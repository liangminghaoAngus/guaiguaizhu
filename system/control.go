package system

import (
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
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
		x, y := ebiten.CursorPosition()
		if gameData.IsPlayerStoreOpen {
			store := component.Store.Get(entry)
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				store.SetSelect(math.NewVec2(float64(x), float64(y)))
			} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
				// fmt.Println("release")
				if store.DragItem != nil {
					cur := math.NewVec2(float64(x)-store.RenderPoint.X, float64(y)-store.RenderPoint.Y)
					curGrid, ind := store.GetItem(cur.X, cur.Y)
					if curGrid == nil {
						curGrid = store.Cap[ind[0]][ind[1]]
					}

					store.Cap[store.DragIndex[0]][store.DragIndex[1]].Drag = false
					if ok := store.SwitchItems(curGrid, store.DragItem); ok {
						// 重置选区
						store.DragItem = nil
						store.DragIndex = [2]int{-1, -1}
					}
				}
			}
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
