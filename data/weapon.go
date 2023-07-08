package data

import (
	"fmt"
	"time"
)

type Weapon struct {
	ID         int       `json:"id" gorm:"column:id;"`
	Name       string    `json:"name" gorm:"column:name;"`
	Image      string    `json:"image" gorm:"column:image;"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;"`
}

func (w *Weapon) TableName() string {
	return "weapon"
}

func GetWeaponByID(id int) *Weapon {
	r := Weapon{}
	if err := getDb().Model(Weapon{}).Where("id = ?", id).First(&r).Error; err != nil {
		fmt.Println(err)
	}
	return &r
}
