package component

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"
	"image/color"
	"testing"
)

type TestGame struct {
	item       *WeaponHandlerData
	Angle      float64
	Count      int
	CountDelay int
}

func (tt *TestGame) Update() error {
	tt.Count++
	if tt.Count >= tt.CountDelay {
		tt.item.Angle += 5
		tt.Count = 0
		tt.Angle += 5
	}

	return nil
}

func (tt *TestGame) Draw(screen *ebiten.Image) {
	parentImage := ebiten.NewImage(50, 80)
	parentImage.Fill(color.White)

	handImg, op := tt.item.GetRenderImage()
	parentImage.DrawImage(handImg, op)

	parentOP := ebiten.DrawImageOptions{}
	parentOP.GeoM.Translate(50, 50)
	screen.DrawImage(parentImage, &parentOP)
	//angleOp := ebiten.DrawImageOptions{}
	//angleOp.GeoM.Rotate(tt.Angle)
	//angleOp.GeoM.Translate(50, 50)
	//screen.DrawImage(handImg, &angleOp)
}

func (tt *TestGame) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func TestWeaponHandlerData_GetRenderPoint(t *testing.T) {
	ebiten.SetWindowSize(640, 320)

	hand := ebiten.NewImage(40, 20)
	hand.Fill(color.Black)
	item := &WeaponHandlerData{
		Image:  hand,
		Point:  math.Vec2{X: 20, Y: 20},
		Angle:  0,
		Width:  float64(hand.Bounds().Dx()),
		Height: float64(hand.Bounds().Dy()),
	}
	handImg, op := item.GetRenderImage()
	fmt.Println(handImg, op)
	item.Angle += 5
	handImg, op = item.GetRenderImage()
	fmt.Println(handImg, op)
	item.Angle += 5
	handImg, op = item.GetRenderImage()
	fmt.Println(handImg, op)
	fmt.Println()
	//if err := ebiten.RunGame(&TestGame{item: item, CountDelay: 10}); err != nil {
	//	log.Fatal(err)
	//}
	//t.Log("")
}
