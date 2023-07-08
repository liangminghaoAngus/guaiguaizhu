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

//go:embed ui/teleport.png
var Teleport []byte

//go:embed ui/hp_1.png
var hp1 []byte

//go:embed ui/hp_2.png
var hp2 []byte

//go:embed ui/hp_3.png
var hp3 []byte

//go:embed ui/hp_4.png
var hp4 []byte

//go:embed ui/hp_5.png
var hp5 []byte

var HpLevelImage = map[int][]byte{
	1: hp1,
	2: hp2,
	3: hp3,
	4: hp4,
	5: hp5,
}

//go:embed ui/mp_1.png
var mp1 []byte

//go:embed ui/mp_2.png
var mp2 []byte

//go:embed ui/mp_3.png
var mp3 []byte

//go:embed ui/mp_4.png
var mp4 []byte

//go:embed ui/mp_5.png
var mp5 []byte

var MpLevelImage = map[int][]byte{
	1: mp1,
	2: mp2,
	3: mp3,
	4: mp4,
	5: mp5,
}

//go:embed npc
var NpcImagesDir embed.FS

var NpcImages = make(map[string]*ebiten.Image)

//go:embed power.png
var PowerBackgroud []byte

//go:embed exp.png
var ExpBackground []byte

//go:embed startBg.png
var StartScreenImage []byte

//go:embed logo.png
var StartScreenLogo []byte

//go:embed ui/ui.png
var SystemUI []byte

//go:embed ui/hp.png
var SystemHP []byte

//go:embed ui/mp.png
var SystemMP []byte

//go:embed race
var raceImageDir embed.FS

//go:embed scene
var sceneImageDir embed.FS

//go:embed human_stand
var humanStandImageDir embed.FS
var HumanStandImgs = make([]*ebiten.Image, 0)
var HumanStandImgsLeft = make([]*ebiten.Image, 0)

//go:embed human_movement
var humanMovementImageDir embed.FS
var HumanMovementLeftImgs = make([]*ebiten.Image, 0)
var HumanMovementRightImgs = make([]*ebiten.Image, 0)

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

		if mapInt, ok := enums.MapImages[imageName]; ok {
			MapImage[mapInt] = raw
		}
		//switch imageName {
		//case "rookie_map.png":
		//	MapImage[enums.MapRookie] = raw
		//}
	}
	humanStandDirEntry, err := humanStandImageDir.ReadDir("human_stand")
	if err != nil {
		panic(err)
	}
	HumanStandImgs = make([]*ebiten.Image, len(humanStandDirEntry))
	HumanStandImgsLeft = make([]*ebiten.Image, len(humanStandDirEntry))
	for index, entry := range humanStandDirEntry {
		imageName := entry.Name()
		filePath := path.Join("human_stand", imageName)
		raw, _ := humanStandImageDir.ReadFile(filePath)
		i, _, _ := image.Decode(bytes.NewReader(raw))

		b := i.Bounds()
		newImage := image.NewRGBA(b)
		for x := b.Min.X; x < b.Max.X; x++ {
			for y := b.Min.Y; y < b.Max.Y; y++ {
				newPixel := i.At((b.Max.X-1)-x, y)
				newImage.Set(x, y, newPixel)
			}
		}

		// 生成向左的图片
		HumanStandImgsLeft[index] = ebiten.NewImageFromImage(newImage)

		HumanStandImgs[index] = ebiten.NewImageFromImage(i)
	}

	humanMovementDirEntry, err := humanMovementImageDir.ReadDir("human_movement")
	if err != nil {
		panic(err)
	}
	HumanMovementLeftImgs = make([]*ebiten.Image, len(humanMovementDirEntry))
	HumanMovementRightImgs = make([]*ebiten.Image, len(humanMovementDirEntry))
	for index, entry := range humanMovementDirEntry {
		imageName := entry.Name()
		filePath := path.Join("human_movement", imageName)
		raw, _ := humanMovementImageDir.ReadFile(filePath)
		i, _, _ := image.Decode(bytes.NewReader(raw))

		b := i.Bounds()
		newImage := image.NewRGBA(b)
		for x := b.Min.X; x < b.Max.X; x++ {
			for y := b.Min.Y; y < b.Max.Y; y++ {
				newPixel := i.At((b.Max.X-1)-x, y)
				newImage.Set(x, y, newPixel)
			}
		}

		// 生成向左的图片
		HumanMovementLeftImgs[index] = ebiten.NewImageFromImage(newImage)

		HumanMovementRightImgs[index] = ebiten.NewImageFromImage(i)
	}

	npcDirEntry, err := NpcImagesDir.ReadDir("npc")
	if err != nil {
		panic(err)
	}
	for _, entry := range npcDirEntry {
		imageName := entry.Name()
		split := strings.Split(imageName, ".")
		filePath := path.Join("npc", imageName)
		raw, _ := NpcImagesDir.ReadFile(filePath)
		i, _, _ := image.Decode(bytes.NewReader(raw))
		if len(split) == 2 {
			NpcImages[split[0]] = ebiten.NewImageFromImage(i)
		}
	}
}
