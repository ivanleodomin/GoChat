package postgresql

import (
	user "app-go/internal"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Register(user user.User, hash, salt string) UserModel {
	register := UserModel{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Hash:      hash,
		Salt:      salt,
	}
	r.db.Create(&user)
	return register
}
