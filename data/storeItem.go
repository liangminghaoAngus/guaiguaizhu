package data

import "time"

// item for backpack , could be same weapon with different ability
type StoreItem struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	ItemID     int       `json:"item_id"`
	Data       string    `json:"data"`
}

func (i *StoreItem) TableName() string {
	return "store_item"
}

type Item struct {
	ID    int    `json:"id"`
	Count int    `json:"count"`
	Name  string `json:"name"`
	Data  string `json:"data"`
}

func (i *Item) TableName() string {
	return "item"
}
