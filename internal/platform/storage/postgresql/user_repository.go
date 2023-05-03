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

func (u *UserRepository) GetAll(page int) ([]user.User, error) {
	var users []UserModel
	limit := 15
	result := db.Limit(limit).Offset((page - 1) * limit).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	var userList []user.User
	for _, userDb := range users {
		userList = append(userList, userDb.toUser())
	}

	return userList, nil
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
