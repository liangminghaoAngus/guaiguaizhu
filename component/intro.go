package component

import "github.com/yohamta/donburi"

type IntroData struct {
	ID    string
	Type  int
	Name  string
	Intro string
}

var Intro = donburi.NewComponentType[IntroData](IntroData{})

func MustGetIntro(entry *donburi.Entry) *IntroData {
	return Intro.Get(entry)
}
