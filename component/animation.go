package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"image/color"
)

type AnimationData struct {
	image  *ebiten.Image
	screen *ebiten.Image
	//animateTime    time.Duration
	screenIndex int
	//animationTimer float64
	//lastUpdateTime time.Time
}

func (d *AnimationData) OutOfBound() bool {
	return d.screenIndex >= d.screen.Bounds().Dx()
}

func (d *AnimationData) Update() error {
	d.screenIndex++
	return nil
}

func (d *AnimationData) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(d.screenIndex), 0)
	screen.DrawImage(d.image, op)
}

func NewAnimation(screen *ebiten.Image, width, height int) AnimationData {
	bgImage := ebiten.NewImage(width, height)
	bgImage.Fill(color.Black)
	return AnimationData{
		image:  bgImage,
		screen: screen,
	}
}

var Animation = donburi.NewComponentType[AnimationData](AnimationData{})
