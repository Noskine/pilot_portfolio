package models

import (
	"time"

	"github.com/google/uuid"
)

type Publications struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title     string    `gorm:"type:varchar(255)"`
	SubTitle  string    `gorm:"type:varchar(100)"`
	Text      string    `gorm:"type:text"`
	Author    string    `gorm:"type:varchar(100)"`
	Image     string    `gorm:"type:varchar(255)"`
	Like      int       `gorm:"type:integer"`
	Create_at time.Time
}
