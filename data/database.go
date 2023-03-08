package data

import (
	"fmt"

	"github.com/gabereiser/datalab/config"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var DB *Database

func NewDatabase() *Database {
	if DB == nil {
		dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", config.Config.DatabaseUser, config.Config.DatabasePassword, config.Config.DatabaseUrl, config.Config.DatabasePort, config.Config.DatabaseName)
		DB = &Database{
			db: gorm.Open(dsn),
		}
	}
	return DB
}
