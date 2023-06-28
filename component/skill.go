package component

import "github.com/yohamta/donburi"

type SkillData struct {
	Name   string
	Info   string
	CostHP int
	CostMP int
	ID     int
}

var Skill = donburi.NewComponentType[SkillData](SkillData{})
