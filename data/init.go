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

	if !db.Migrator().HasTable("ability") {
		db.Migrator().CreateTable(Ability{})
	}

}

func getDb() *gorm.DB {
	return db.Debug()
}
