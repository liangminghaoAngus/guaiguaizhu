package system

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Render struct {
	query     *query.Query
	playerUI  *query.Query
	offscreen *ebiten.Image
}

func NewRender() *Render {
	r := &Render{
		query: query.NewQuery(
			filter.And(
				filter.Contains(transform.Transform),
				filter.Or(filter.Contains(component.Sprite), filter.Contains(component.SpriteStand)),
				filter.Not(filter.Contains(component.Map)))),
		playerUI:  query.NewQuery(filter.Contains(component.Health, component.Player, component.Level, component.Store)),
		offscreen: ebiten.NewImage(3000, 3000),
	}
	return r
}

func (r *Render) Update(w donburi.World) {

	// 修改 sprite 渲染
	r.query.Each(w, func(entry *donburi.Entry) {
		// 判断是否实体存在 spriteStand
		if entry.HasComponent(component.SpriteStand) {
			standImages := component.SpriteStand.Get(entry)
			if !standImages.Disabled {
				index := (standImages.Count / 5) % 8
				if index > len(standImages.Images)-1 {
					standImages.Count = 0
					index = 0
				}
				standImages.Count++
			} else {
				standImages.Count = 0 // 重置动画
			}
		}

		// 判断是否实体存在 spriteMovement
		if entry.HasComponent(component.SpriteMovement) {
			move := component.SpriteMovement.Get(entry)
			if !move.Disabled {
				index := (move.Count / 5) % 8
				if index > len(move.LeftImages)-1 {
					move.Count = 0
					index = 0
				}
				move.Count++
			} else {
				move.Count = 0 // 重置动画
			}
		}

	})

}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {

	//nowTime := time.Now()
	gameData := component.MustFindGame(w)
	var entries []*donburi.Entry

	r.query.Each(w, func(entry *donburi.Entry) {
		entries = append(entries, entry)
		pos := transform.WorldPosition(entry)
		position := component.Position.Get(entry)

		if entry.HasComponent(component.Collision) && entry.HasComponent(component.Position) {
			collision := component.Collision.Get(entry)
			for _, object := range collision.Items {
				object.Position.X = pos.X + position.X
				object.Position.Y = pos.Y + position.Y
				if collision.Debug {
					debugBounds := ebiten.NewImage(int(object.Width), int(object.Height))
					debugBounds.Fill(color.Black)
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(pos.X+position.X, pos.Y+position.Y)
					screen.DrawImage(debugBounds, op)
				}
			}
		}

		if entry.HasComponent(component.SpriteMovement) && entry.HasComponent(component.Position) {
			movementImages := component.SpriteMovement.Get(entry)
			if !movementImages.Disabled {
				index := (movementImages.Count / 5) % 8
				// 判断是否需要翻转贴图方向
				targetImage := &ebiten.Image{}
				if movementImages.IsDirectionRight {
					targetImage = movementImages.RightImages[index]
				} else {
					targetImage = movementImages.LeftImages[index]
				}

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(position.X+pos.X, position.Y+pos.Y)
				screen.DrawImage(targetImage, op)
			}
		}

		if entry.HasComponent(component.SpriteStand) && entry.HasComponent(component.Position) {
			// position := component.Position.Get(entry)
			standImages := component.SpriteStand.Get(entry)
			if !standImages.Disabled {
				index := (standImages.Count / 5) % 8
				// 判断是否需要翻转贴图方向
				targetImage := &ebiten.Image{}
				if standImages.IsDirectionRight {
					targetImage = standImages.Images[index]
				} else {
					targetImage = standImages.ImagesRight[index]
				}

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(position.X+pos.X, position.Y+pos.Y)
				screen.DrawImage(targetImage, op)
			}
		}

	})

	playerEntity, ok := r.playerUI.First(w)

	// 绘制 UI
	{
		x := screen.Bounds().Max.X/2 - gameData.SystemUI.Bounds().Dx()/2
		y := screen.Bounds().Max.Y - gameData.SystemUI.Bounds().Dy()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(gameData.SystemUI, op)
		if ok {
			health := component.Health.Get(playerEntity)
			level := component.Level.Get(playerEntity)
			{ // hp
				health.DrawPlayerHPImage(screen, health.HPui, x, y, health.HP, 1)
				// get last time hp info ,do draw
				lastTimeHP, ok := health.LastAnimationItem[component.AnimateHp]
				lastTime, ok := health.LastAnimationTime[component.AnimateHp]
				animateTime := health.AnimationTime
				if elapsed := time.Since(lastTime); elapsed.Seconds() < animateTime.Seconds() && ok {
					// 绘制上一次的内容
					health.DrawPlayerHPImage(screen, health.HPui, x, y, lastTimeHP, 0.65)
				}

				//hpImage := health.HPui
				//percent := float64(health.HP) / float64(health.HPMax)
				//x0 := 0
				//y0 := hpImage.Bounds().Dy()
				//x1 := hpImage.Bounds().Dx()
				//y1 := float64(hpImage.Bounds().Dy()) * (float64(1) - percent)
				//
				//
				//op := &ebiten.DrawImageOptions{}
				//op.GeoM.Translate(float64(x+10), float64(y+12+int(y1)))
				//screen.DrawImage(hpImage.SubImage(image.Rect(x0, y0, x1, int(y1))).(*ebiten.Image), op)
			}
			{ // mp
				health.DrawPlayerMPImage(screen, health.MPui, x+gameData.SystemUI.Bounds().Dx(), y, health.MP, 1)
				// get last time hp info ,do draw
				lastTimeMP, ok := health.LastAnimationItem[component.AnimateMp]
				lastTime, ok := health.LastAnimationTime[component.AnimateMp]
				animateTime := health.AnimationTime
				if elapsed := time.Since(lastTime); elapsed.Seconds() < animateTime.Seconds() && ok {
					// 绘制上一次的内容
					health.DrawPlayerMPImage(screen, health.MPui, x+gameData.SystemUI.Bounds().Dx(), y, lastTimeMP, 0.65)
				}
				//mpImage := health.MPui
				//percent := float64(health.MP) / float64(health.MPMax)
				//x0 := 0
				//y0 := mpImage.Bounds().Dy()
				//x1 := mpImage.Bounds().Dx()
				//y1 := float64(mpImage.Bounds().Dy()) * (float64(1) - percent)
				//
				//transx := x + gameData.SystemUI.Bounds().Dx() - mpImage.Bounds().Dx()
				//
				//op := &ebiten.DrawImageOptions{}
				//op.GeoM.Translate(float64(transx-16), float64(y+12+int(y1)))
				//screen.DrawImage(mpImage.SubImage(image.Rect(x0, y0, x1, int(y1))).(*ebiten.Image), op)
			}
			{
				percent := float64(level.Exp) / float64(level.ExpNextLevel)
				img, _, _ := image.Decode(bytes.NewReader(assetImages.ExpBackground))
				x0 := 0
				y0 := img.Bounds().Dy()
				x1 := img.Bounds().Dx()
				y1 := float64(img.Bounds().Dy()) * (float64(1) - percent)
				fixedY := float64(img.Bounds().Dy())

				bgImg := ebiten.NewImageFromImage(img)
				font := config.GetSystemFontSize(12)
				bText := text.BoundString(font, fmt.Sprint(level.LevelNum))
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x+health.HPui.Bounds().Dx()+11), float64(screen.Bounds().Max.Y)-float64(fixedY)-12+float64(y1))
				// screen.DrawImage(bgImg, op)
				screen.DrawImage(bgImg.SubImage(image.Rect(x0, y0, x1, int(y1))).(*ebiten.Image), op)
				text.Draw(screen, fmt.Sprint(level.LevelNum), font, health.HPui.Bounds().Dx()+x+22-bText.Dx()/2, screen.Bounds().Max.Y-int(fixedY)+2, color.White)
				// fmt.Println(level.LevelNum)
			}
		}
	}

	if gameData.IsPlayerStoreOpen {
		if ok {
			backpack := component.Store.Get(playerEntity)
			// draw store
			backpack.DrawBackpackUI(screen)
		}
	}
}
