package data

import (
	"os"
	"path"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	work, _ := os.Getwd()
	dbPath := path.Join(work, "data", "ggz.db")
	if con, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}); err != nil {
		panic("connect db err")
	} else {
		db = con
	}

	tableList := map[string]interface{}{
		"ability":     Ability{},
		"npc":         Npc{},
		"save_game":   SaveGame{},
		"item":        Item{},
		"store_item":  StoreItem{},
		"teleport":    Teleport{},
		"weapon":      Weapon{},
		"enemy":       Enemy{},
		"player_task": PlayerTask{},
		"task":        Task{},
	}
	for table, structTable := range tableList {
		if !db.Migrator().HasTable(table) {
			_ = db.Migrator().CreateTable(structTable)
		} else {
			_ = db.AutoMigrate(structTable)
		}
	}
}

func getDb() *gorm.DB {
	return db.Debug()
}
