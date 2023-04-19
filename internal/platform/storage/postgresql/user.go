package postgresql

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model
	ID        int
	Firstname string
	Lastname  string
	Hash      string
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
