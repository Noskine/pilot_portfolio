package models

import (
	"time"

	"github.com/google/uuid"
)

type Publications struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string
	Text      string
	Author    string
	Like      int
	Create_at time.Time
}
