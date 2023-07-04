package entity

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"liangminghaoangus/guaiguaizhu/component"
)

//type Teleport struct {
//
//}

var Teleport = []donburi.IComponentType{
	transform.Transform,
	component.Position,
	component.Sprite,
	// todo
}
