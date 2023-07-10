package data

import (
	"encoding/json"
	"fmt"
	"liangminghaoangus/guaiguaizhu/enums"
)

type Enemy struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Intro   string `json:"intro"`
	Map     int    `json:"map"`
	Health  string `json:"health"`
	Level   int    `json:"level"`
	BeatExp int    `json:"beat_exp"`
	Box     string `json:"box"`
	Image   string `json:"image"`
	// todo hitbox
}

type EnemyItem struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Intro   string      `json:"intro"`
	Map     int         `json:"map"`
	Health  EnemyHealth `json:"health"`
	Level   int         `json:"level"`
	BeatExp int         `json:"beat_exp"`
	Box     EnemyBox    `json:"box"`
	Image   string      `json:"image"`
}

type EnemyBox struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type EnemyHealth struct {
	Hp    int `json:"hp"`
	Mp    int `json:"mp"`
	HpMax int `json:"hp_max"`
	MpMax int `json:"mp_max"`
}

func (e *Enemy) TransLate2EnemyItem() *EnemyItem {
	tmp := &EnemyItem{
		ID:    e.ID,
		Name:  e.Name,
		Intro: e.Intro,
		Map:   e.Map,
		//Health:  data.EnemyHealth{},
		Level:   e.Level,
		BeatExp: e.BeatExp,
		//Box:     data.EnemyBox{},
		Image: e.Image,
	}
	if e.Health != "" {
		if raw, err := json.Marshal(e.Health); err == nil {
			h := EnemyHealth{}
			_ = json.Unmarshal(raw, &h)
			tmp.Health = h
		}
	}
	if e.Box != "" {
		if raw, err := json.Marshal(e.Box); err == nil {
			h := EnemyBox{}
			_ = json.Unmarshal(raw, &h)
			tmp.Box = h
		}
	}
	return tmp
}

func (e *Enemy) TableName() string {
	return "enemy"
}

func GetEnemyByID(id int) *Enemy {
	res := Enemy{}
	if err := getDb().Model(Enemy{}).Where("id = ?", id).First(&res).Error; err != nil {
		fmt.Println(err)
	}
	return &res
}

func GetEnemyByMap(mapInt enums.Map) []*Enemy {
	res := make([]*Enemy, 0)
	if err := getDb().Model(Enemy{}).Where("map = ?", mapInt).Find(&res).Error; err != nil {
		fmt.Println(err)
	}
	return res
}
