package component

import "github.com/yohamta/donburi"

type StoreData struct {
	CapNum       int
	CapItemIDMap map[int][]byte
	Cap          [][]*StoreItem
}

type StoreItem struct {
	ID    int
	Type  int
	Count int
}

var defaultStore = func() StoreData {
	col, row := 6, 8
	c := make([][]*StoreItem, col)
	for i := 0; i < col; i++ {
		l := make([]*StoreItem, row)
		c = append(c, l)
	}
	return StoreData{
		CapNum:       row * col,
		CapItemIDMap: make(map[int][]byte),
		Cap:          c,
	}
}()

var Store = donburi.NewComponentType[StoreData](defaultStore)
