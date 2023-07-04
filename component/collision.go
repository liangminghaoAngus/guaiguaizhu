package component

import (
	"github.com/yohamta/donburi"
	"liangminghaoangus/guaiguaizhu/engine"
)

type CollisionData struct {
	Items     []*engine.Object // 碰撞检测组件
	TagsOrder []string         // 组件按标签叠层
	Debug     bool             // 渲染边界值
}

var Collision = donburi.NewComponentType[CollisionData](CollisionData{})

type CollisionSpaceData struct {
	Space *engine.Space
}

var CollisionSpace = donburi.NewComponentType[CollisionSpaceData](CollisionSpaceData{})
