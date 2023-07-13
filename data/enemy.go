package data

import (
	"encoding/json"
	"liangminghaoangus/guaiguaizhu/enums"
	"liangminghaoangus/guaiguaizhu/log"
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

		h := EnemyHealth{}
		if err := json.Unmarshal([]byte(e.Health), &h); err == nil {
			tmp.Health = h
		} else {
			log.Error("%s", err.Error())
		}

	}
	if e.Box != "" {

		h := EnemyBox{}
		if err := json.Unmarshal([]byte(e.Box), &h); err == nil {
			tmp.Box = h
		} else {
			log.Error("%s", err.Error())
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
		log.Error("%s", err.Error())
	}
	return &res
}

func GetEnemyByMap(mapInt enums.Map) []*Enemy {
	res := make([]*Enemy, 0)
	if err := getDb().Model(Enemy{}).Where("map = ?", mapInt).Find(&res).Error; err != nil {
		log.Error("%s", err.Error())
	}
	return res
}
