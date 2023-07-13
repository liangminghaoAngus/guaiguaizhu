package data

import (
	"liangminghaoangus/guaiguaizhu/enums"
	"liangminghaoangus/guaiguaizhu/log"
)

type Ability struct {
	ID        int    `json:"id" gorm:"column:id;"`
	Race      int    `json:"race" gorm:"column:race;"`
	NeedLevel int    `json:"need_level" gorm:"column:need_level;"`
	Name      string `json:"name" gorm:"column:name;"`
	Intro     string `json:"intro" gorm:"column:intro;"`
	CostHP    int    `json:"cost_hp" gorm:"column:cost_hp;"`
	CostMP    int    `json:"cost_mp" gorm:"column:cost_mp;"`
}

func (a *Ability) TableName() string {
	return "ability"
}

func ListAbilityByRace(raceInt enums.Race) []Ability {
	res := make([]Ability, 0)
	if err := getDb().Model(Ability{}).Where("race = ?", raceInt).Order("need_level").Find(&res).Error; err != nil {
		log.Error("%s", err.Error())
	}
	return res
}
