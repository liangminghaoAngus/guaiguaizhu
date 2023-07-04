package component

import (
	"github.com/yohamta/donburi"
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
