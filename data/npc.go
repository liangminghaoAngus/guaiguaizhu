package data

import "fmt"

type Npc struct {
	ID       int    `json:"id" gorm:"column:id;"`
	Type     int    `json:"type" gorm:"column:type;"`
	Name     string `json:"name" gorm:"column:name;"`
	Intro    string `json:"intro" gorm:"column:intro;"`
	Position string `json:"position" gorm:"column:position;"`
	Map      int    `json:"map" gorm:"column:map;"`
	Image    string `json:"image" gorm:"column:image;"`
}

func (n *Npc) TableName() string {
	return "npc"
}

func GetNpc(id int) *Npc {
	if l := GetNpcByID([]int{id}); len(l) > 0 {
		return &l[0]
	}
	return nil
}

func GetNpcByID(ids []int) []Npc {
	r := make([]Npc, 0)
	if err := getDb().Model(Npc{}).Where("id in ?", ids).Find(&r).Error; err != nil {
		fmt.Println(err)
	}
	return r
}

func GetNpcByMapID(mapID int) []int {
	r := make([]int, 0)
	if err := getDb().Model(Npc{}).Where("map = ?", mapID).Select("id").Find(&r).Error; err != nil {
		fmt.Println(err)
	}
	return r
}
