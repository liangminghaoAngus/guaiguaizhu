package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"image/color"
	"liangminghaoangus/guaiguaizhu/config"
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

func (i *InteractionData) ShowDialog(screen *ebiten.Image) *ebiten.Image {
	if i.DialogCache == nil {
		dialog := ebiten.NewImage(screen.Bounds().Dx(), 200)
		dialog.Fill(color.Black)
		font := config.GetSystemFontSize(16)
		if i.TalkWord != nil {
			//textBound := text.BoundString(font,*i.TalkWord)
			text.Draw(dialog, *i.TalkWord, font, 10, 10, color.White)
		}
		i.DialogCache = dialog
	}
	return i.DialogCache
}

var Interaction = donburi.NewComponentType[InteractionData](InteractionData{})
