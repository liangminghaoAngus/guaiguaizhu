package entity

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"liangminghaoangus/guaiguaizhu/component"
	"liangminghaoangus/guaiguaizhu/enums"
)

//	var PlayerEntity = map[string]donburi.IComponentType{
//		"HealthData":   component.Health,
//		"RaceData":     component.Race,
//		"LevelData":    component.Level,
//		"MovementData": component.Movement,
//		"PositionData": component.Position,
//		"SpriteData":   component.Sprite,
//	}

var PlayerEntity = []donburi.IComponentType{
	transform.Transform,
	component.Health,
	component.Race,
	component.Level,
	component.Movement,
	component.Position,
	//component.Sprite,
	component.Control,
}

func NewPlayer(w donburi.World, raceInt enums.Race) *donburi.Entry {
	if name := enums.GetRaceName(raceInt); name == "" {
		panic("unknow race")
	}

	playerEntity := w.Create(PlayerEntity...)
	player := w.Entry(playerEntity)
	component.Health.SetValue(player, component.NewPlayerHealthData())
	component.Race.SetValue(player, component.NewRaceData(raceInt))
	component.Level.SetValue(player, component.NewLevelData())
	component.Movement.SetValue(player, component.NewMovementData())
	component.Position.SetValue(player, component.NewPositionData())
	//component.Sprite.SetValue(player, component.SpriteData{})
	component.Control.SetValue(player, component.NewPlayerControl())

	return nil
}

//func NewPlayer(w donburi.World, playerNumber int, faction component.PlayerFaction) *donburi.Entry {
//	_, ok := Players[playerNumber]
//	if !ok {
//		panic(fmt.Sprintf("unknown player number: %v", playerNumber))
//	}
//
//	player := component.PlayerData{
//		PlayerNumber:  playerNumber,
//		PlayerFaction: faction,
//		Lives:         3,
//		RespawnTimer:  engine.NewTimer(time.Second * 3),
//		WeaponLevel:   component.WeaponLevelSingle,
//	}
//
//	// TODO It looks like a constructor would fit here
//	player.ShootTimer = engine.NewTimer(player.WeaponCooldown())
//
//	return NewPlayerFromPlayerData(w, player)
//}
//
//func NewPlayerFromPlayerData(w donburi.World, playerData component.PlayerData) *donburi.Entry {
//	player := w.Entry(w.Create(component.Player))
//	component.Player.SetValue(player, playerData)
//	return player
//}
