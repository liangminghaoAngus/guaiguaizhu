package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/enums"
	"sort"
	"time"
)

type AbilityData struct {
	Items map[int]SkillItem
}

type SkillItem struct {
	Image     *ebiten.Image
	Name      string
	Info      string
	CoolDown  time.Duration
	CostHP    int
	CostMP    int
	NeedLevel int
	Type      int
	ID        int
}

var Ability = donburi.NewComponentType[AbilityData](AbilityData{})

func (s *AbilityData) ListOrderByLevel() []SkillItem {
	res := make([]SkillItem, 0)
	for _, v := range s.Items {
		i := v
		res = append(res, i)
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i].Type != res[j].Type {
			return res[i].Type < res[j].Type
		} else {
			return res[i].NeedLevel < res[j].NeedLevel
		}
	})
	return res
}

func (s *AbilityData) DrawAbilityList(level int) *ebiten.Image {
	itemCeil := 60
	margin := 10
	l := s.ListOrderByLevel()
	lineImg := ebiten.NewImage(len(l)*itemCeil+(len(l)+1)*margin, itemCeil+2*margin)
	for ind, item := range l {
		grid := ebiten.NewImage(itemCeil, itemCeil)
		x := (ind+1)*margin + ind*itemCeil
		y := margin + itemCeil
		scale := float64(item.Image.Bounds().Dx()) / float64(itemCeil)
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Scale(scale, scale)
		if !(item.NeedLevel > level) {
			// draw
			grid.DrawImage(item.Image, ops)
		}
		opsLine := &ebiten.DrawImageOptions{}
		opsLine.GeoM.Translate(float64(x), float64(y))
		lineImg.DrawImage(grid, opsLine)
	}
	return lineImg
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
