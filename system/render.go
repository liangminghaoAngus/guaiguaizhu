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

// var implCount = 10
// var countFrame = 0

func NewRender() *Render {
	r := &Render{
		query: query.NewQuery(
			filter.And(
				filter.Contains(transform.Transform, component.Position),
				filter.Or(filter.Contains(component.Sprite), filter.Contains(component.SpriteStand)),
				filter.Not(filter.Contains(component.Map, component.NotActive)))),
		playerUI:  query.NewQuery(filter.Contains(component.Health, component.Heal, component.Player, component.Level, component.Store)),
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
				standImages.Count++
				index := (standImages.Count / 5) % len(standImages.Images)
				if index > len(standImages.Images)-1 {
					standImages.Count = 0
					index = 0
				}
			} else {
				standImages.Count = 0 // 重置动画
			}
		}

		// 判断是否实体存在 spriteMovement
		if entry.HasComponent(component.SpriteMovement) {
			move := component.SpriteMovement.Get(entry)
			if !move.Disabled {
				move.Count++
				index := (move.Count / 5) % len(move.LeftImages)
				if index > len(move.LeftImages)-1 {
					move.Count = 0
					index = 0
				}
			} else {
				move.Count = 0 // 重置动画
			}
		}

		// if entry.HasComponent(component.WeaponHandler) {
		// 	weaponHand := component.WeaponHandler.Get(entry)
		// 	countFrame++
		// 	if countFrame >= implCount {
		// 		countFrame = 0
		// 		weaponHand.Angle += 5
		// 	}
		// }

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

		if entry.HasComponent(component.Sprite) && !entry.HasComponent(component.NotActive) {
			sprite := component.Sprite.Get(entry)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(pos.X+position.X, pos.Y+position.Y)
			screen.DrawImage(sprite.Image, op)
		}

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
			box := component.Box.Get(entry)
			if !movementImages.Disabled {
				index := (movementImages.Count / 5) % len(movementImages.RightImages)
				boxImage := ebiten.NewImage(box.Width, box.Height)
				// 判断是否需要翻转贴图方向
				targetImage := &ebiten.Image{}
				if movementImages.IsDirectionRight {
					targetImage = movementImages.RightImages[index]
				} else {
					targetImage = movementImages.LeftImages[index]
				}
				ops := &ebiten.DrawImageOptions{}
				scale := float64(boxImage.Bounds().Dx()) / float64(targetImage.Bounds().Dx())
				ops.GeoM.Scale(scale, scale)
				boxImage.DrawImage(targetImage, ops)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(position.X+pos.X, position.Y+pos.Y)
				screen.DrawImage(boxImage, op)
			}
		}

		// 武器绘制
		if entry.HasComponent(component.WeaponHandler) && entry.HasComponent(component.Box) {
			weaponHand := component.WeaponHandler.Get(entry)
			// box := component.Box.Get(entry)
			boxImage := ebiten.NewImage(200, 200)
			drawX, drawY := weaponHand.GetRenderPoint().X, weaponHand.GetRenderPoint().Y
			ops := &ebiten.DrawImageOptions{}
			ops.GeoM.Rotate(weaponHand.Angle)
			ops.GeoM.Translate(drawX, drawY)
			boxImage.DrawImage(weaponHand.Image, ops)
			boxOps := &ebiten.DrawImageOptions{}
			boxOps.GeoM.Translate(position.X+pos.X, position.Y+pos.Y)
			if weaponHand.Weapon != nil {
				weaponBox := ebiten.NewImage(int(weaponHand.Weapon.Width), int(weaponHand.Weapon.Height))
				weaponBox.DrawImage(weaponHand.Weapon.Image, nil)
				ops := &ebiten.DrawImageOptions{}
				ops.GeoM.Translate(drawX+weaponHand.WeaponPoint.X, drawY+weaponHand.WeaponPoint.Y)
				boxImage.DrawImage(weaponBox, ops)
			}
			screen.DrawImage(boxImage, boxOps)
		}

		if entry.HasComponent(component.SpriteStand) && entry.HasComponent(component.Position) && entry.HasComponent(component.Box) {
			standImages := component.SpriteStand.Get(entry)
			box := component.Box.Get(entry)
			if !standImages.Disabled {
				index := (standImages.Count / 5) % len(standImages.Images)
				// 判断是否需要翻转贴图方向
				boxImage := ebiten.NewImage(box.Width, box.Height)
				targetImage := &ebiten.Image{}
				if standImages.IsDirectionRight {
					targetImage = standImages.Images[index]
				} else {
					targetImage = standImages.ImagesRight[index]
				}
				ops := &ebiten.DrawImageOptions{}
				scale := float64(boxImage.Bounds().Dx()) / float64(targetImage.Bounds().Dx())
				ops.GeoM.Scale(scale, scale)
				boxImage.DrawImage(targetImage, ops)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(position.X+pos.X, position.Y+pos.Y)
				screen.DrawImage(boxImage, op)
			}
		}

	})

	playerEntity, ok := r.playerUI.First(w)

	// 绘制技能栏
	{
		level := component.Level.Get(playerEntity)
		playerSkills := component.Ability.Get(playerEntity)
		skillImage := playerSkills.DrawAbilityList(level.LevelNum)
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(skillImage, op)
	}

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

			}
			heal := component.Heal.Get(playerEntity)
			scaleNum := 16
			{
				//  todo 可以先绘制一个底图，后续将图片放置进去，再做居中处理
				// hp heal
				raw := assetImages.HpLevelImage[heal.GetLevel()]
				img, _, _ := image.Decode(bytes.NewReader(raw))
				fixedY := float64(img.Bounds().Dy())
				opt := &ebiten.DrawImageOptions{}
				opt.GeoM.Scale(float64(scaleNum)/float64(img.Bounds().Dx()), float64(scaleNum)/float64(img.Bounds().Dy()))
				opt.GeoM.Translate(float64(x+149), float64(screen.Bounds().Max.Y)-float64(fixedY)-11)
				screen.DrawImage(ebiten.NewImageFromImage(img), opt)

			}
			{
				// mp heal 180
				raw := assetImages.MpLevelImage[heal.GetLevel()]
				img, _, _ := image.Decode(bytes.NewReader(raw))
				fixedY := float64(img.Bounds().Dy())
				opt := &ebiten.DrawImageOptions{}
				opt.GeoM.Scale(float64(scaleNum)/float64(img.Bounds().Dx()), float64(scaleNum)/float64(img.Bounds().Dy()))
				opt.GeoM.Translate(float64(x+177), float64(screen.Bounds().Max.Y)-float64(fixedY)-11)
				screen.DrawImage(ebiten.NewImageFromImage(img), opt)
			}
		}
	}

	if gameData.IsPlayerStoreOpen {
		if ok {
			backpack := component.Store.Get(playerEntity)
			// draw store
			backpack.DrawBackpackUI(screen)

			x, y := ebiten.CursorPosition()
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				store := component.Store.Get(playerEntity)

				if store.DragItem != nil {
					ops := &ebiten.DrawImageOptions{}
					ops.GeoM.Translate(float64(x-store.DragItem.Image.Bounds().Dx()/2), float64(y-store.DragItem.Image.Bounds().Dy()/2))
					ops.ColorScale.ScaleAlpha(0.5)
					screen.DrawImage(store.DragItem.Image, ops)
				}
			}
		}

	}
}
