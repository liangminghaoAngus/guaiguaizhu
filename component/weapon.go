package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"image/color"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/log"
)

type WeaponData struct {
	Level     int
	AttackNum int

	Image *ebiten.Image

	Angle         float64
	Width, Height float64
}

var Weapon = donburi.NewComponentType[WeaponData](WeaponData{})

func NewWeapon(weaponID int, renderPoint math.Vec2) WeaponData {
	// 获取数据
	weaponData := data.GetWeaponByID(weaponID)
	log.Info("init weapon %s", weaponData.Name)
	// debug test
	weaponImage := ebiten.NewImage(20, 50)
	weaponImage.Fill(color.Black)
	d := WeaponData{
		Level:     0,
		AttackNum: 1,
		Image:     weaponImage,
		Angle:     0,
		Width:     20,
		Height:    50,
	}

	return d
}
