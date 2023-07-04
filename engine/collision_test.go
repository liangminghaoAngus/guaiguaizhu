package engine

import (
	"fmt"
	"testing"
)

func TestNewObject(t *testing.T) {
	spaceW, spaceH := float64(1280), float64(640)

	space := NewSpace(spaceW, spaceH)

	bounders := []*Object{
		NewObject(0, 0, spaceW, 2),
		NewObject(0, spaceH-2, spaceW, 2),
		NewObject(0, 0, 2, spaceH),
		NewObject(spaceW-2, 0, 2, spaceH),
	}
	space.AddObject(bounders...)

	object := NewObject(1, 1, 50, 80)
	space.AddObject(object)

	//dx := float64(500)
	if c := object.Check(800, 300); c != nil {
		fmt.Println(c)
	} else {
		fmt.Println("no collision")
	}

	return
}
