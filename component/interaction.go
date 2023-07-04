package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type InteractionData struct {
	Rect *ebiten.Image
}

var Interaction = donburi.NewComponentType[InteractionData](InteractionData{})
