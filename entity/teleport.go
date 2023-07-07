package entity

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"image/color"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/enums"
	"strconv"
	"strings"
)

//type Teleport struct {
//
//}

var Teleport = []donburi.IComponentType{
	component.Teleport,
	component.Intro, // just for information
	transform.Transform,
	component.Position,
	component.Sprite, // todo animate
	component.Box,
}

func NewTeleports(world donburi.World) []*donburi.Entry {

	teleports := data.ListAllTeleports()
	res := make([]*donburi.Entry, len(teleports))
	teleportEntitys := world.CreateMany(len(teleports), Teleport...)
	for ind, entity := range teleportEntitys {
		entry := world.Entry(entity)

		info := teleports[ind]
		Pos := math.Vec2{}
		ToPos := math.Vec2{}
		if l := strings.Split(info.ToPosition, ","); len(l) == 2 {
			x, _ := strconv.Atoi(l[0])
			y, _ := strconv.Atoi(l[1])
			ToPos = math.NewVec2(float64(x), float64(y))
		}
		if l := strings.Split(info.Position, ","); len(l) == 2 {
			x, _ := strconv.Atoi(l[0])
			y, _ := strconv.Atoi(l[1])
			Pos = math.NewVec2(float64(x), float64(y))
		}

		component.Teleport.SetValue(entry, component.TeleportData{
			ToMap:      enums.Map(info.ToMap),
			ToPosition: ToPos,
		})
		component.Intro.SetValue(entry, component.IntroData{
			ID:    fmt.Sprintf("teleport_%d", info.ID),
			Type:  0,
			Name:  info.Name,
			Intro: "",
		})
		component.Position.SetValue(entry, component.PositionData{
			X:             Pos.X,
			Y:             Pos.Y,
			Map:           enums.Map(info.Map),
			FaceDirection: 0,
		})
		component.Box.SetValue(entry, component.NewTeleportBox())
		box := component.Box.Get(entry)

		// todo animate image
		teleportImage := ebiten.NewImage(box.Width, box.Height)
		teleportImage.Fill(color.White)
		component.Sprite.SetValue(entry, component.SpriteData{
			Image: teleportImage,
			//ColorOver:        component.ColorOverride{},
			//HiddenED:         false,
			//Layer:            0,
			//Pivot:            0,
			//OriginalRotation: 0,
		})

		res[ind] = entry
	}
	return res
}

func InTeleport(world donburi.World, mainItem *donburi.Entry, mapInt enums.Map) *donburi.Entry {
	//player := MustFindPlayerEntry(world)
	playerPos := component.Position.Get(mainItem)
	playerBox := component.Box.Get(mainItem)

	teleports := ListMapTeleport(world, mapInt)
	for _, v := range teleports {
		pos := component.Position.Get(v)
		box := component.Box.Get(v)
		if playerPos.X >= pos.X && playerPos.X+float64(playerBox.Width) <= pos.X+float64(box.Width) {
			return v
		}
	}

	return nil
}

func ListMapTeleport(world donburi.World, mapInt enums.Map) []*donburi.Entry {
	r := make([]*donburi.Entry, 0)
	component.Teleport.Each(world, func(entry *donburi.Entry) {
		if entry.HasComponent(component.Position) {
			pos := component.Position.Get(entry)
			if pos.Map == mapInt {
				r = append(r, entry)
			}
		}
	})
	return r
}
