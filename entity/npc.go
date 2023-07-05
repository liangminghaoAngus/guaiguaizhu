package entity

import (
	"fmt"
	assetImages "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/enums"
	"strconv"
	"strings"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var NPCEntity = []donburi.IComponentType{
	transform.Transform,
	component.Position,
	component.Sprite,
	component.Intro,
}

func NewNPCs(w donburi.World, npcIDs []int) []*donburi.Entry {
	npcEntitys := make([]*donburi.Entry, len(npcIDs))
	for i, npcID := range npcIDs {
		npcEntity := w.Create(NPCEntity...)
		npc := w.Entry(npcEntity)
		intro := component.MustGetIntro(npc)
		position := component.Position.Get(npc)
		image := component.Sprite.Get(npc)
		intro.ID = fmt.Sprintf("npc_%d", npcID)
		// search npc from data
		if dbNpc := data.GetNpc(npcID); dbNpc != nil {
			intro.Name = dbNpc.Name
			intro.Type = dbNpc.Type
			intro.Intro = dbNpc.Intro
			if p := strings.Split(dbNpc.Position, ","); len(p) == 2 {
				x, _ := strconv.Atoi(p[0])
				y, _ := strconv.Atoi(p[1])
				position.X = float64(x)
				position.Y = float64(y)
				position.Map = enums.Map(dbNpc.Map)
			}
			// l := []*ebiten.Image{assetImages.NpcImages[fmt.Sprintf("%d", npcID)]}
			image.Image = assetImages.NpcImages[fmt.Sprintf("%d", npcID)]
			// image.Images = l
			// image.ImagesRight = l
		}
		npcEntitys[i] = npc
	}
	return npcEntitys
}
