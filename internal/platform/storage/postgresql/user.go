package postgresql

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Hash      string
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
