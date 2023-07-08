package component

import (
	"bytes"
	"image"

	assetImages "liangminghaoangus/guaiguaizhu/assets/images"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
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
	selectXY      math.Vec2
}

type StoreItem struct {
	Image    *ebiten.Image
	Exist    bool
	ID       int
	UUID     string
	Type     int
	Drag     bool
	Count    int  // todo
	MaxCount int  // todo
	CanGroup bool // todo
}

var defaultStore = func() StoreData {
	col, row := 4, 7
	c := make([][]*StoreItem, col)
	for i := 0; i < col; i++ {
		l := make([]*StoreItem, row)
		c[i] = l
		// c = append(c, l)
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

func (d *StoreData) SetSelect(pos math.Vec2) {
	d.selectXY = pos
}

func (d *StoreData) DrawUI() {
	itemCeil := 60
	margin := 5
	marginBox := 16
	grid, _, _ := image.Decode(bytes.NewReader(assetImages.BagGrid))
	gridImg := ebiten.NewImageFromImage(grid)
	img, _, _ := image.Decode(bytes.NewReader(assetImages.BagPanel))
	gridScale := float64(itemCeil) / float64(grid.Bounds().Dx())

	uiMain := ebiten.NewImageFromImage(img)
	for i := 0; i < d.Height; i++ {
		lineImg := ebiten.NewImage(img.Bounds().Dx(), itemCeil+3*margin)
		for j := 0; j < d.Width; j++ {
			// Calculate the position of each item
			x := j*itemCeil + (j+1)*margin

			ops := &ebiten.DrawImageOptions{}
			ops.GeoM.Scale(gridScale, gridScale)
			ops.GeoM.Translate(float64(x), float64(margin))
			lineImg.DrawImage(gridImg, ops)
		}
		y := i*itemCeil + (i+1)*margin
		lineOps := &ebiten.DrawImageOptions{}
		lineOps.GeoM.Translate(float64(uiMain.Bounds().Dx()/2-lineImg.Bounds().Dx()/2+marginBox), float64(y+marginBox))
		uiMain.DrawImage(lineImg, lineOps)
	}

	d.MainUI = uiMain
}

func (d *StoreData) DrawBackpackUI(screen *ebiten.Image) {
	// d.DrawUI()
	uiMain := d.MainUI

	itemCeil := 60
	margin := 5
	marginBox := 16

	// draw Items
	for i, row := range d.Cap {
		y := i*itemCeil + (i+1)*margin
		for j, item := range row {
			if item == nil || !item.Exist {
				continue
			}
			x := j*itemCeil + (j+1)*margin
			i := item.Image.Bounds()
			gridScale := float64(itemCeil) / float64(i.Dx())
			ops := &ebiten.DrawImageOptions{}
			ops.GeoM.Scale(gridScale, gridScale)
			ops.GeoM.Translate(float64(x+marginBox), float64(y)+float64(margin+marginBox))
			uiMain.DrawImage(item.Image, ops)
		}
	}

	op := &ebiten.DrawImageOptions{}
	x, y := float64(screen.Bounds().Dx()/2-uiMain.Bounds().Dx()/2), float64(screen.Bounds().Dy()/2-uiMain.Bounds().Dy()/2)
	op.GeoM.Translate(x, y)
	// 居中绘制
	screen.DrawImage(uiMain, op)
}

func (d *StoreData) AddItem(item StoreItem) bool {
	// 判断同种类型是否存在背包中
	itemIndex := NewIndex()
FullLoop:
	for row := 0; row < len(d.Cap); row++ {
		line := d.Cap[row]
		for col := 0; col < len(line); col++ {
			if d.Cap[row][col] == nil || !d.Cap[row][col].Exist {
				itemIndex = &Index{col, row}
				break FullLoop
			}
		}
	}
	// 如果背包满了，无法存放，返回 false
	if !itemIndex.Check() {
		return false
	}

	//
	if d.Cap[itemIndex.X][itemIndex.Y] == nil {
		d.Cap[itemIndex.X][itemIndex.Y] = &StoreItem{}
	}

	d.Cap[itemIndex.X][itemIndex.Y].Image = item.Image
	d.Cap[itemIndex.X][itemIndex.Y].Exist = true
	d.Cap[itemIndex.X][itemIndex.Y].ID = item.ID
	d.Cap[itemIndex.X][itemIndex.Y].UUID = item.UUID
	d.Cap[itemIndex.X][itemIndex.Y].Count += 1

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
