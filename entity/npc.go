package entity

import (
	"fmt"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"liangminghaoangus/guaiguaizhu/component"
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
		data := component.MustGetIntro(npc)
		data.ID = fmt.Sprintf("npc_%d", npcID)
		// todo search npc from data
		//data.Name =
		//data.Type =
		//data.Intro =
		npcEntitys[i] = npc
	}
	return npcEntitys
}
