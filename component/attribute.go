package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/enums"
)

type AttributeData struct {
	AvailablePoint int

	Power    int
	Strength int
	Quick    int
	Magic    int
	Energy   int
}

func (a *AttributeData) LevelUp(item enums.Attribute, count int) {
	if count > a.AvailablePoint {
		count = a.AvailablePoint
	}
	switch item {
	case enums.AttributePower:
		a.Power += count
	case enums.AttributeStrength:
		a.Strength += count
	case enums.AttributeQuick:
		a.Quick += count
	case enums.AttributeMagic:
		a.Magic += count
	case enums.AttributeEnergy:
		a.Energy += count
	}
}

func (a *AttributeData) GetAdditionNum(item enums.Attribute) int {
	i := 0
	switch item {
	case enums.AttributePower:
		i = a.Power * 1
	case enums.AttributeStrength:
		i = a.Strength * 5
	case enums.AttributeQuick:
		i = a.Quick * 2
	case enums.AttributeMagic:
		i = a.Magic * 1
	case enums.AttributeEnergy:
		i = a.Energy * 5
	}
	return i
}

var defaultAttribute = AttributeData{
	AvailablePoint: 0,
	Power:          0,
	Strength:       0,
	Quick:          0,
	Magic:          0,
	Energy:         0,
}

var Attribute = donburi.NewComponentType[AttributeData](defaultAttribute)
