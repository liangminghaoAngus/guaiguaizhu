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
	RenderPoint   math.Vec2
	MainUI        *ebiten.Image
	Height, Width int
	CapNum        int
	CapItemIDMap  map[int]Index
	Cap           [][]*StoreItem

	selectXY  math.Vec2
	DragIndex [2]int
	DragItem  *StoreItem
}

type StoreItem struct {
	Pos0, Pos1 math.Vec2
	Index      [2]int

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

func (st *StoreItem) IsExist() bool {
	return st != nil && st.Exist
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

func (d *StoreData) SetSelect(pos math.Vec2) *StoreItem {
	// itemCeil := 60
	// 需要减去 bag panel 起始点
	d.selectXY = math.NewVec2(pos.X-d.RenderPoint.X, pos.Y-d.RenderPoint.Y)
	// x := pos.X / float64(itemCeil)
	// y := pos.Y / float64(itemCeil)
	item, index := d.GetItem(d.selectXY.X, d.selectXY.Y)
	if item == nil {
		return nil
	}
	item.Drag = true
	d.DragItem = item
	d.DragIndex = index
	return item
}

func (d *StoreData) GetItem(x, y float64) (*StoreItem, [2]int) {
	for i, row := range d.Cap {
		// y := i*itemCeil + (i+1)*margin
		for j, item := range row {
			if item.Pos0.X <= x && item.Pos1.X >= x && item.Pos0.Y <= y && item.Pos1.Y >= y {
				if item == nil || !item.Exist {
					item = nil
				}
				return item, [2]int{i, j}
			}
		}
	}

	return nil, [2]int{}
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
		y := i*itemCeil + (i+1)*margin
		lineImg := ebiten.NewImage(img.Bounds().Dx(), itemCeil+3*margin)
		lineX := float64(uiMain.Bounds().Dx()/2 - lineImg.Bounds().Dx()/2 + marginBox)
		lineY := float64(y + marginBox)
		for j := 0; j < d.Width; j++ {
			// Calculate the position of each item
			x := j*itemCeil + (j+1)*margin

			d.Cap[i][j] = &StoreItem{
				Exist: false,
				Index: [2]int{i, j},
				Pos0:  math.NewVec2(float64(x)+lineX, float64(y)+lineY),
				Pos1:  math.NewVec2(float64(x+itemCeil)+lineX, float64(y+itemCeil)+lineY),
			}

			ops := &ebiten.DrawImageOptions{}
			ops.GeoM.Scale(gridScale, gridScale)
			ops.GeoM.Translate(float64(x), float64(margin))
			lineImg.DrawImage(gridImg, ops)
		}

		lineOps := &ebiten.DrawImageOptions{}
		lineOps.GeoM.Translate(lineX, lineY)
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
			if item == nil || !item.Exist || item.Drag {
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
	d.RenderPoint = math.NewVec2(x, y)
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

func (d *StoreData) SwitchItems(item1, item2 *StoreItem) bool {
	if len(item1.Index) < 2 || len(item2.Index) < 2 {
		return false
	}

	item1Row, item1Col := item1.Index[0], item1.Index[1]
	item2Row, item2Col := item2.Index[0], item2.Index[1]

	d.Cap[item1Row][item1Col], d.Cap[item2Row][item2Col] = d.Cap[item2Row][item2Col], d.Cap[item1Row][item1Col]

	return true
}
