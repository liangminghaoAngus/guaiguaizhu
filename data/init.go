package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
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
}

func getDb() *gorm.DB {
	return db
}
