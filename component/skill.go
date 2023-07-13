package component

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/enums"
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
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
	itemCeil := 30
	margin := 10
	l := s.ListOrderByLevel()
	lineImg := ebiten.NewImage(len(l)*itemCeil+(len(l)+1)*margin, itemCeil+2*margin)
	for ind, item := range l {
		grid := ebiten.NewImage(itemCeil, itemCeil)
		x := (ind+1)*margin + ind*itemCeil
		y := margin
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
	raceName := enums.GetRaceStr(raceInt)
	for _, val := range l {
		width, height := 30, 30
		imgBox := ebiten.NewImage(width, height)
		imgBox.Fill(color.Black)
		raw, _ := assetsImage.SkillImageDir.ReadFile(fmt.Sprintf("skill/%s/%d.png", raceName, val.ID))
		i, _, _ := image.Decode(bytes.NewReader(raw))
		scale := float64(width) / float64(i.Bounds().Dx())
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Scale(scale, scale)
		imgBox.DrawImage(ebiten.NewImageFromImage(i), ops)
		m[val.ID] = SkillItem{
			Image:     imgBox,
			Type:      0,
			CoolDown:  0,
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
