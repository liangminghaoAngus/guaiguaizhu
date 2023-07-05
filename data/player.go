package data

import (
	"fmt"
	"time"
)

type SaveGame struct {
	ID         int       `json:"id"`
	SaveData   string    `json:"save_data"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (g *SaveGame) TableName() string {
	return "save_game"
}

func LoadPlayerData(id int) (map[string]interface{}, error) {
	saveData := make(map[string]interface{})
	if err := getDb().Model(SaveGame{}).Where("id = ?", id).First(&saveData).Error; err != nil {
		return nil, fmt.Errorf("读取存档失败")
	}
	return saveData, nil
}

func SaveGameChangeOrigin(id int, data string) error {
	return getDb().Model(SaveGame{}).Where("id = ?", id).Updates(map[string]interface{}{
		"save_data":   data,
		"update_time": time.Now(),
	}).Error
}

func SavePlayerData2Local(data string) (int, error) {
	now := time.Now()
	item := SaveGame{
		SaveData:   data,
		CreateTime: now,
		UpdateTime: now,
	}
	if err := getDb().Create(&item).Error; err != nil {
		return -1, fmt.Errorf("存档失败！")
	}
	return item.ID, nil
}
