package models

import (
	"time"

	"github.com/google/uuid"
)

type Publications struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string    `gorm:"type:varchar(100)"`
	Text      string    `gorm:"type:text"`
	Author    string    `gorm:"type:varchar(100)"`
	Image     string    `gorm:"type:varchar(255)"`
	Like      int
	Create_at time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
