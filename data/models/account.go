package models

import (
	"time"

	"github.com/google/uuid"
)

type AccountModel struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Email    string    `gorm:"index;unique;size:255"`
	Password string    `gorm:"size:255;not null"`
	VToken   string    `gorm:"not null"`
	Joined   time.Time `gorm:"autoCreateTime"`
	Disabled bool
}
