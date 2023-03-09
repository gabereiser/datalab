package data

import (
	"fmt"

	"github.com/gabereiser/datalab/config"
	"github.com/gabereiser/datalab/data/models"
	"github.com/gabereiser/datalab/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ID = models.ID

// NewID takes an optional string and returns an UUID compatible ID struct
func NewID(id *string) ID {
	return models.NewID(id)
}

type AccountModel = models.AccountModel
type OrganizationModel = models.OrganizationModel

var DB *gorm.DB

func NewDatabase() {
	if DB == nil {
		dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", config.Config.DatabaseUser, config.Config.DatabasePassword, config.Config.DatabaseUrl, config.Config.DatabasePort, config.Config.DatabaseName)
		d, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if e != nil {
			log.Err("%v", e)
		}
		DB = d
	}
}

func Migrate() {
	if DB == nil {
		NewDatabase()
	}
	DB.AutoMigrate(&AccountModel{})
	DB.AutoMigrate(&OrganizationModel{})
}

func FindAccount(email string) *AccountModel {
	var account *AccountModel
	err := DB.Where("email = ?", email).First(account).Error
	if err != nil {
		log.Err("%v", err)
		return nil
	}
	return account
}

func GetAccount(id ID) *AccountModel {
	return nil
}
