package entity

import (
	"fmt"
	"liangminghaoangus/guaiguaizhu/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var EnemyEntity = []donburi.IComponentType{
	component.Enemy,
	transform.Transform,
	component.Position,
	component.Health,
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

	// todo get enemy info from data

	return entitys
}
