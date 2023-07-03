package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/enums"
	"sort"
)

type AbilityData struct {
	Items map[int]SkillItem
}

type SkillItem struct {
	Name      string
	Info      string
	CostHP    int
	CostMP    int
	NeedLevel int
	ID        int
}

var Ability = donburi.NewComponentType[AbilityData](AbilityData{})

func (s *AbilityData) ListByLevel() []SkillItem {
	res := make([]SkillItem, 0)
	for _, v := range s.Items {
		i := v
		res = append(res, i)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].NeedLevel < res[j].NeedLevel
	})
	return res
}

func NewAbility(raceInt enums.Race) AbilityData {
	m := make(map[int]SkillItem)
	l := data.ListAbilityByRace(raceInt)
	for _, val := range l {
		m[val.ID] = SkillItem{
			Name:      val.Name,
			Info:      val.Intro,
			CostHP:    val.CostHP,
			CostMP:    val.CostMP,
			NeedLevel: val.NeedLevel,
			ID:        val.ID,
		}
	}
	return AbilityData{Items: m}
}
