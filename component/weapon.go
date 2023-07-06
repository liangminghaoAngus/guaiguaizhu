package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"
	systemMath "math"
)

type WeaponHandlerData struct {
	Image *ebiten.Image

	Point       math.Vec2
	WeaponPoint math.Vec2
	Weapon      *WeaponData

	Angle         float64
	Width, Height float64
}

type WeaponData struct {
	Image *ebiten.Image

	Angle         float64
	Width, Height float64
}

func (we *WeaponHandlerData) GetRenderPoint() math.Vec2 {
	return calculatePoint(we.Point, 0, we.Height/2)
}

func (we *WeaponHandlerData) GetRenderImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	imgPoint := we.GetRenderPoint()
	tranX, tranY := imgPoint.X, imgPoint.Y
	op.GeoM.Rotate(we.Angle)
	op.GeoM.Translate(tranX, tranY)
	return we.Image, op
}

func (we *WeaponHandlerData) GetWeaponBox() (point math.Vec2, angle float64, w, h float64) {
	// todo

	return math.Vec2{}, 0, 0, 0
}

func calculatePoint(start math.Vec2, angle float64, length float64) math.Vec2 {
	radians := angle * (systemMath.Pi / 180.0)
	x := start.X + length*systemMath.Cos(radians)
	y := start.Y + length*systemMath.Sin(radians)
	return math.Vec2{X: x, Y: y}
}
