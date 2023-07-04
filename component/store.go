package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Index struct {
	X, Y int
}

func (i *Index) Check() bool {
	return i.X != -1 && i.Y != -1
}

func NewIndex() *Index {
	return &Index{-1, -1}
}

type StoreData struct {
	MainUI        *ebiten.Image
	Height, Width int
	CapNum        int
	CapItemIDMap  map[int]Index
	Cap           [][]*StoreItem
}

type StoreItem struct {
	ID       int
	UUID     string
	Type     int
	Count    int  // todo
	MaxCount int  // todo
	CanGroup bool // todo
}

var defaultStore = func() StoreData {
	col, row := 6, 8
	c := make([][]*StoreItem, col)
	for i := 0; i < col; i++ {
		l := make([]*StoreItem, row)
		c = append(c, l)
	}
	return StoreData{
		Width:        row,
		Height:       col,
		CapNum:       row * col,
		CapItemIDMap: make(map[int]Index),
		Cap:          c,
	}
}()

var Store = donburi.NewComponentType[StoreData](defaultStore)

func MustFindStore(w donburi.World) *StoreData {
	entry, ok := query.NewQuery(filter.Contains(Store, Player)).First(w)
	if !ok {
		panic("not store player")
	}
	return Store.Get(entry)
}

func (d *StoreData) DrawUI() {
	itemCeil := 50
	borderSize := 2
	wpx, hpx := itemCeil*d.Width, itemCeil*d.Height
	uiMain := ebiten.NewImage(wpx, hpx)
	for i := 0; i < d.Height; i++ {
		for j := 0; j < d.Width; j++ {
			// Calculate the position of each item
			x := i * itemCeil
			y := j * itemCeil

			// todo draw item with drag & drop

			// Draw a border around each item
			vector.DrawFilledRect(uiMain, float32(x), float32(y), float32(itemCeil), float32(borderSize), color.Black, false)
			vector.DrawFilledRect(uiMain, float32(x), float32(y)+float32(itemCeil-borderSize), float32(itemCeil), float32(borderSize), color.Black, false) // Bottom border
			vector.DrawFilledRect(uiMain, float32(x), float32(y), float32(borderSize), float32(itemCeil), color.Black, false)
			// Left border
			vector.DrawFilledRect(uiMain, float32(x)+float32(itemCeil-borderSize), float32(y), float32(borderSize), float32(itemCeil), color.Black, false) // Right border
		}
	}
	d.MainUI = uiMain
}

func (d *StoreData) DrawBackpackUI(screen *ebiten.Image) {
	uiMain := d.MainUI

	op := &ebiten.DrawImageOptions{}
	x, y := float64(screen.Bounds().Dx()/2-uiMain.Bounds().Dx()/2), float64(screen.Bounds().Dy()/2-uiMain.Bounds().Dy()/2)
	op.GeoM.Translate(x, y)
	// 居中绘制
	screen.DrawImage(uiMain, op)
}

func (d *StoreData) AddItem(item StoreItem) bool {
	// 判断同种类型是否存在背包中
	itemIndex := NewIndex()
	for row := 0; row < len(d.Cap); row++ {
		line := d.Cap[row]
		for col := 0; col < len(line); col++ {
			if d.Cap[row][col] == nil {
				itemIndex = &Index{col, row}
			}
		}
	}
	// 如果背包满了，无法存放，返回 false
	if !itemIndex.Check() {
		return false
	}

	//
	d.Cap[itemIndex.X][itemIndex.Y] = &item

	return true
}

func (d *StoreData) RemoveItem(item StoreItem) bool {

	for row := 0; row < d.Height; row++ {
		for col := 0; col < d.Width; col++ {
			if d.Cap[row][col].UUID == item.UUID {
				d.Cap[row][col] = nil
				return true
			}
		}
	}
	return false

}

func (d *StoreData) SwitchItems(item1, item2 StoreItem) bool {
	item1Found := false
	item2Found := false
	item1Row, item1Col := -1, -1
	item2Row, item2Col := -1, -1
	for row := 0; row < d.Height; row++ {
		for col := 0; col < d.Width; col++ {
			if d.Cap[row][col].UUID == item1.UUID {
				item1Found = true
				item1Row, item1Col = row, col
			} else if d.Cap[row][col].UUID == item2.UUID {
				item2Found = true
				item2Row, item2Col = row, col
			}
		}
	}
	if item1Found && item2Found {
		d.Cap[item1Row][item1Col], d.Cap[item2Row][item2Col] = d.Cap[item2Row][item2Col], d.Cap[item1Row][item1Col]
		return true
	}
	return false
}
