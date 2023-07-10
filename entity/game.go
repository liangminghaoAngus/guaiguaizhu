package entity

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"liangminghaoangus/guaiguaizhu/component"
)

var GameEntity = []donburi.IComponentType{
	component.Game,
	component.Animation,
	transform.Transform,
}
