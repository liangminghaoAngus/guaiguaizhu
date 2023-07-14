package component

import (
	"bytes"
	"github.com/fishtailstudio/imgo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"image"
	assetsImage "liangminghaoangus/guaiguaizhu/assets/images"
	"liangminghaoangus/guaiguaizhu/log"
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
	IsLeft    bool
	Head      *NodeItem
	Body      *NodeItem
	Hand      *NodeItem
	LeftFoot  *NodeItem
	RightFoot *NodeItem
}

var PlayerNode = donburi.NewComponentType[PlayerNodeData](PlayerNodeData{})

func (p *PlayerNodeData) UpdateStand(frameCount int) {
	// stand animation
	animationStand := []math.Vec2{
		{X: 0, Y: 0.2},
		{X: 0, Y: 0.2},
		{X: 0, Y: 0.15},
		{X: 0, Y: 0.2},
		{X: 0, Y: -0.15},
		{X: 0, Y: -0.15},
		{X: 0, Y: -0.15},
		{X: 0, Y: -0.15},
		{X: 0, Y: -0.15},
	}
	if frameCount > len(animationStand) || frameCount < 0 {
		return
	}
	animationPoint := animationStand[frameCount]
	// body head hand
	p.Body.RenderPoint = p.Body.RenderPoint.Add(animationPoint)
	p.Head.RenderPoint = p.Head.RenderPoint.Add(animationPoint)
	p.Hand.RenderPoint = p.Head.RenderPoint.Add(animationPoint)
}

type MoveRenderPoint struct {
	Head, Body, Hand, LeftFoot, RightFoot struct {
		Angle float64
		Point *math.Vec2
	}
}

func (p *PlayerNodeData) GetMoveRenderPoint(frameCount int) MoveRenderPoint {
	handAngle := []float64{-9.816, -20.02, -30, -20, -10, 0, 19.83, 40.07, 60, 40, 19.8, 0}
	res := MoveRenderPoint{}
	if p.Hand != nil {
		i := p.Hand.RenderPoint.Rotate(handAngle[frameCount])
		res.Hand = struct {
			Angle float64
			Point *math.Vec2
		}{Angle: handAngle[frameCount], Point: &i}
	}

	return res
}

func (p *PlayerNodeData) UpdateMovement(frameCount int) {
	// move animation
	// hand move angle
	handAngle := []float64{-9.816, -20.02, -30, -20, -10, 0, 19.83, 40.07, 60, 40, 19.8, 0}
	log.Info("%s", handAngle)
	p.Hand.RenderPoint.Rotate(handAngle[frameCount])

}

func (p *PlayerNodeData) Draw(screen *ebiten.Image, frameCount int) {
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
		//reflect.TypeOf(node)
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Scale((node.Width)/float64(node.ImageLeft.Bounds().Dx()), (node.Height)/float64(node.ImageLeft.Bounds().Dx()))
		ops.GeoM.Translate(node.RenderPoint.X, node.RenderPoint.Y)
		targetImage := &ebiten.Image{}
		if p.IsLeft {
			targetImage = node.ImageLeft
		} else {
			targetImage = node.ImageRight
		}
		screen.DrawImage(targetImage, ops)
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
		IsLeft:    false,
		Head:      headNode,
		Body:      BodyNode,
		Hand:      HandNode,
		LeftFoot:  LeftFootNode,
		RightFoot: RightFootNode,
	}

	return nodeData
}
