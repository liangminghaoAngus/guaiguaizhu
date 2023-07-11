package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"
)

type NodeItem struct {
	ZIndex int

	Point         math.Vec2
	Width, Height int
	RenderPoint   math.Vec2
	Image         *ebiten.Image
	Angle         float64
}

type PlayerNode struct {
	Head      *NodeItem
	Body      *NodeItem
	Hand      *NodeItem
	LeftFoot  *NodeItem
	RightFoot *NodeItem
}
