package images

import (
	"embed"
	"liangminghaoangus/guaiguaizhu/enums"
	"path"
	"strconv"
	"strings"
)

//go:embed race
var raceImageDir embed.FS

//go:embed scene
var sceneImageDir embed.FS

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
}
