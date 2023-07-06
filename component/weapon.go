package component

import (
	"fmt"
	systemMath "math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
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

var WeaponHandler = donburi.NewComponentType[WeaponHandlerData](WeaponHandlerData{})

func (we *WeaponHandlerData) GetRenderPoint() math.Vec2 {
	width := we.Width / 2
	// height := we.Height
	// c := systemMath.Sqrt(width*width + height*height)
	p1 := RotatePoint(math.Vec2{X: we.Point.X - width, Y: we.Point.Y}, we.Point, we.Angle)
	fmt.Println(p1)
	return p1
}

func (we *WeaponHandlerData) GetRenderImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(-we.Width/2, -we.Height/2)
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

func RotatePoint(p, center math.Vec2, angle float64) math.Vec2 {
	// Convert the angle to radians
	angleRad := angle * systemMath.Pi / 180.0

	// Translate the point to the origin
	translatedPoint := math.Vec2{
		X: p.X - center.X,
		Y: p.Y - center.Y,
	}

	// Rotate the translated point around the origin
	rotatedPoint := math.Vec2{
		X: translatedPoint.X*systemMath.Cos(angleRad) - translatedPoint.Y*systemMath.Sin(angleRad),
		Y: translatedPoint.X*systemMath.Sin(angleRad) + translatedPoint.Y*systemMath.Cos(angleRad),
	}

	// Translate the rotated point back to its original position
	finalPoint := math.Vec2{
		X: rotatedPoint.X + center.X,
		Y: rotatedPoint.Y + center.Y,
	}

	return finalPoint
}
