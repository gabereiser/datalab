package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationModel struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"index;unique;size:255"`
	Url      string    `gorm:"size:255;not null"`
	Created  time.Time `gorm:"autoCreateTime"`
	Disabled bool

	Accounts []*AccountModel `gorm:"many2many:organizations_accounts"`
}

func (OrganizationModel) TableName() string {
	return "organizations"
}
