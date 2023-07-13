package data

import (
	"liangminghaoangus/guaiguaizhu/log"
)

type Teleport struct {
	ID       int    `json:"id" gorm:"column:id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Map      int    `json:"map" gorm:"column:map;"`
	Position string `json:"position" gorm:"column:position;"`

	ToMap      int    `json:"to_map" gorm:"column:to_map;"`
	ToPosition string `json:"to_position" gorm:"column:to_position;"`
	Image      string `json:"image" gorm:"column:image;"`
}

func (t *Teleport) TableName() string {
	return "teleport"
}

func ListAllTeleports() []*Teleport {
	l := make([]*Teleport, 0)
	if err := getDb().Model(Teleport{}).Find(&l).Error; err != nil {
		log.Error("%s", err.Error())
	}
	return l
}
