package component

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

type CollisionData struct {
	Items     []*resolv.Object // 碰撞检测组件
	TagsOrder []string         // 组件按标签叠层
	Debug     bool             // 渲染边界值
}

var Collision = donburi.NewComponentType[CollisionData](CollisionData{})

type CollisionSpaceData struct {
	Space *resolv.Space
}

var CollisionSpace = donburi.NewComponentType[CollisionSpaceData](CollisionSpaceData{})
