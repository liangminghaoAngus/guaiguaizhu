package entity

import (
	"fmt"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/data"
	"liangminghaoangus/guaiguaizhu/enums"
	"strconv"
	"strings"
)

var NPCEntity = []donburi.IComponentType{
	transform.Transform,
	component.Position,
	component.SpriteStand,
	component.Intro,
}

func NewNPCs(w donburi.World, npcIDs []int) []*donburi.Entry {
	npcEntitys := make([]*donburi.Entry, len(npcIDs))
	for i, npcID := range npcIDs {
		npcEntity := w.Create(NPCEntity...)
		npc := w.Entry(npcEntity)
		intro := component.MustGetIntro(npc)
		position := component.Position.Get(npc)
		//image := component.SpriteStand.Get(npc)
		intro.ID = fmt.Sprintf("npc_%d", npcID)
		// search npc from data
		if dbNpc := data.GetNpc(npcID); dbNpc != nil {
			intro.Name = dbNpc.Name
			intro.Type = dbNpc.Type
			intro.Intro = dbNpc.Intro
			if p := strings.Split(dbNpc.Position, ","); len(p) == 0 {
				x, _ := strconv.Atoi(p[0])
				y, _ := strconv.Atoi(p[1])
				position.X = float64(x)
				position.Y = float64(y)
				position.Map = enums.Map(dbNpc.Map)
			}
			// todo image
			//ebiten.NewImage()
		}
		npcEntitys[i] = npc
	}
	return npcEntitys
}
