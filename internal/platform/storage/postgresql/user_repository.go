package postgresql

import (
	user "app-go/internal"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := GetDB()
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Register(user user.User) {
	register := UserModel{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Hash:      user.Hash,
		Salt:      user.Salt,
	}
	r.db.Create(&register)
	return
}
