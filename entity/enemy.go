package entity

import (
	"fmt"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"liangminghaoangus/guaiguaizhu/component"
)

var EnemyEntity = []donburi.IComponentType{
	transform.Transform,
	component.Position,
	component.Health,
	component.Enemy,
	component.Level,
	component.Intro,
	component.Movement,
	component.SpriteStand,
	component.SpriteMovement,
	component.Collision,
}

// todo
func NewEnemyEntity(w donburi.World, enemyID int, num int) []*donburi.Entity {
	entitys := make([]*donburi.Entity, num)

	enemyEntitys := w.CreateMany(num, EnemyEntity...)
	fmt.Println(enemyEntitys)
	//player := w.Entry(playerEntity)

	return entitys
}
