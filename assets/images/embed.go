package images

import (
	"bytes"
	"embed"
	"image"
	"liangminghaoangus/guaiguaizhu/enums"
	"path"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed race
var raceImageDir embed.FS

//go:embed scene
var sceneImageDir embed.FS

//go:embed human_stand
var humanStandImageDir embed.FS
var HumanStandImges = make([]*ebiten.Image, 0)

var RaceImage = map[enums.Race][]byte{
	enums.RaceGod:   nil,
	enums.RaceHuman: nil,
	enums.RaceDevil: nil,
}

var RaceHoverImage = map[enums.Race][]byte{
	enums.RaceGod:   nil,
	enums.RaceHuman: nil,
	enums.RaceDevil: nil,
}

var MapImage = map[enums.Map][]byte{
	enums.MapRookie: nil,
}

func Init() {
	raceDirEntry, err := raceImageDir.ReadDir("race")
	if err != nil {
		panic(err)
	}
	for _, entry := range raceDirEntry {
		imageName := entry.Name()
		filePath := path.Join("race", imageName)
		raw, err := raceImageDir.ReadFile(filePath)
		if err != nil {
			continue
		}
		if strings.Contains(imageName, "_") { // hover image
			if l := strings.Split(imageName, "_"); len(l) > 0 {
				i, e := strconv.Atoi(l[0])
				if e != nil {
					continue
				}
				switch enums.Race(i) {
				case enums.RaceGod:
					RaceHoverImage[enums.RaceGod] = raw
				case enums.RaceHuman:
					RaceHoverImage[enums.RaceHuman] = raw
				case enums.RaceDevil:
					RaceHoverImage[enums.RaceDevil] = raw
				}
			}
		} else { // normal image
			if l := strings.Split(imageName, "."); len(l) > 0 {
				i, e := strconv.Atoi(l[0])
				if e != nil {
					continue
				}
				switch enums.Race(i) {
				case enums.RaceGod:
					RaceImage[enums.RaceGod] = raw
				case enums.RaceHuman:
					RaceImage[enums.RaceHuman] = raw
				case enums.RaceDevil:
					RaceImage[enums.RaceDevil] = raw
				}
			}
		}
	}

	sceneDirEntry, err := sceneImageDir.ReadDir("scene")
	if err != nil {
		panic(err)
	}
	for _, entry := range sceneDirEntry {
		imageName := entry.Name()
		filePath := path.Join("scene", imageName)
		raw, _ := sceneImageDir.ReadFile(filePath)
		switch imageName {
		case "rookie_map.png":
			MapImage[enums.MapRookie] = raw
		}
	}
	humanStandDirEntry, err := humanStandImageDir.ReadDir("human_stand")
	if err != nil {
		panic(err)
	}
	HumanStandImges = make([]*ebiten.Image, len(humanStandDirEntry))
	for index, entry := range humanStandDirEntry {
		imageName := entry.Name()
		filePath := path.Join("human_stand", imageName)
		raw, _ := humanStandImageDir.ReadFile(filePath)
		i, _, _ := image.Decode(bytes.NewReader(raw))

		HumanStandImges[index] = ebiten.NewImageFromImage(i)
	}
}
