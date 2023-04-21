package postgresql

import (
	user "app-go/internal"
	"errors"
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
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Hash:      user.Hash,
		Salt:      user.Salt,
	}
	r.db.Create(&register)
	return
}

func (r *UserRepository) Login(email, password string) (*user.User, error) {

	var userDb UserModel
	result := db.First(&userDb, "email = ?", email)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	if !user.Authenticate(password, userDb.Salt, userDb.Hash) {
		return nil, errors.New("incorrect password")
	}

	res := userDb.toUser()

	return &res, nil
}

func (r *UserRepository) GetById(id string) (*user.User, error) {
	var userDb UserModel
	result := db.First(&userDb, "id = ?", id)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	res := userDb.toUser()

	return &res, nil
}

func (r *UserRepository) GetAll() []user.User {
	var usersDb []UserModel
	db.Find(&usersDb)

	users := make([]user.User, len(usersDb))
	for i, u := range usersDb {
		users[i] = u.toUser()
	}

	return users
}

func (u *UserModel) toUser() user.User {
	return user.User{
		ID:        u.ID,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Hash:      u.Hash,
		Salt:      u.Salt,
	}
}
