package component

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"sort"
)

type NodeItem struct {
	ZIndex int

	Point         math.Vec2
	Width, Height int
	RenderPoint   math.Vec2
	Image         *ebiten.Image
	Angle         float64
}

type PlayerNodeData struct {
	Head      *NodeItem
	Body      *NodeItem
	Hand      *NodeItem
	LeftFoot  *NodeItem
	RightFoot *NodeItem
}

var PlayerNode = donburi.NewComponentType[PlayerNodeData](PlayerNodeData{})

func (p *PlayerNodeData) Draw() {
	l := []*NodeItem{p.Head, p.Body, p.Head, p.LeftFoot, p.RightFoot}
	// 删除不存在的节点
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			l = append(l[:i], l[i+1:]...)
			i--
		}
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].ZIndex < l[j].ZIndex
	})
	for _, node := range l {
		fmt.Println(node)
	}
	// todo
}

func NewPlayerNodeData() *PlayerNodeData {
	nodeData := &PlayerNodeData{}

	return nodeData
}
