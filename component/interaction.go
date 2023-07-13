package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type InteractionFunc int

const (
	InteractionFuncTalk InteractionFunc = iota + 1
	InteractionFuncBuy
	InteractionFuncWeaponLevel
	InteractionFuncTeleport
)

type InteractionData struct {
	IsOpen      bool
	DialogCache *ebiten.Image

	Items    []InteractionFunc
	Rect     *ebiten.Image
	TalkWord *string
	Buy      []*StoreItem
	Teleport []*PositionData
	// todo weapon level
}

func (i *InteractionData) ShowDialog() {

}

var Interaction = donburi.NewComponentType[InteractionData](InteractionData{})
