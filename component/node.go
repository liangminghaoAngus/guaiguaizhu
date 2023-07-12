package component

import (
	"bytes"
	"github.com/fishtailstudio/imgo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"image"
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"sort"
)

type NodeItem struct {
	ZIndex int

	Point         math.Vec2
	Width, Height float64
	RenderPoint   math.Vec2
	ImageLeft     *ebiten.Image
	ImageRight    *ebiten.Image
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

func (p *PlayerNodeData) Draw(screen *ebiten.Image) {
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
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Scale((node.Width)/float64(node.ImageLeft.Bounds().Dx()), (node.Height)/float64(node.ImageLeft.Bounds().Dx()))
		ops.GeoM.Translate(node.RenderPoint.X, node.RenderPoint.Y)
		screen.DrawImage(node.ImageLeft, ops)
	}
}

func getNodeImage(raw []byte) (left, right *ebiten.Image) {
	head, _, _ := image.Decode(bytes.NewReader(raw))
	left = ebiten.NewImageFromImage(head)
	right = ebiten.NewImageFromImage(imgo.LoadFromImage(head).Flip(imgo.Horizontal).ToImage())
	return
}

func NewPlayerNodeData() PlayerNodeData {

	//head, _, _ := image.Decode(bytes.NewReader(assetsImage.HumanHead))
	//headRightImg := ebiten.NewImageFromImage(head)
	//headLeftImg := ebiten.NewImageFromImage(imgo.LoadFromImage(head).Flip(imgo.Horizontal).ToImage())

	headLeftImg, headRightImg := getNodeImage(assetsImage.HumanHead)
	bodyLeftImg, bodyRightImg := getNodeImage(assetsImage.HumanBody)
	handLeftImg, handRightImg := getNodeImage(assetsImage.HumanHand)
	leftFootImg, rightFootImg := getNodeImage(assetsImage.HumanFoot)

	headNode := &NodeItem{
		ZIndex:      2,
		Point:       math.Vec2{},
		Width:       40.25,
		Height:      38.5,
		RenderPoint: math.Vec2{X: 0, Y: 0},
		ImageLeft:   headLeftImg,
		ImageRight:  headRightImg,
		Angle:       0,
	}
	BodyNode := &NodeItem{
		ZIndex:      1,
		Point:       math.Vec2{},
		Width:       29,
		Height:      27,
		RenderPoint: math.Vec2{X: 6.8, Y: 29.5},
		ImageLeft:   bodyLeftImg,
		ImageRight:  bodyRightImg,
		Angle:       0,
	}
	HandNode := &NodeItem{
		ZIndex:      3,
		Point:       math.Vec2{},
		Width:       11.6,
		Height:      14.2,
		RenderPoint: math.Vec2{X: 15.4, Y: 40},
		ImageLeft:   handLeftImg,
		ImageRight:  handRightImg,
		Angle:       0,
	}
	LeftFootNode := &NodeItem{
		ZIndex:      0,
		Point:       math.Vec2{},
		Width:       8.7,
		Height:      7.8,
		RenderPoint: math.Vec2{X: 13.65, Y: 53.5},
		ImageLeft:   leftFootImg,
		ImageRight:  leftFootImg,
		Angle:       0,
	}
	RightFootNode := &NodeItem{
		ZIndex:      0,
		Point:       math.Vec2{},
		Width:       8.7,
		Height:      7.8,
		RenderPoint: math.Vec2{X: 20.25, Y: 52.1},
		ImageLeft:   rightFootImg,
		ImageRight:  rightFootImg,
		Angle:       0,
	}

	nodeData := PlayerNodeData{
		Head:      headNode,
		Body:      BodyNode,
		Hand:      HandNode,
		LeftFoot:  LeftFootNode,
		RightFoot: RightFootNode,
	}

	return nodeData
}
