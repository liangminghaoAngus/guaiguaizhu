package sound

import (
	_ "embed"
)

//go:embed intro.wav
var Intro []byte

//go:embed boss.wav
var Boss []byte

//go:embed body.wav
var Body []byte
