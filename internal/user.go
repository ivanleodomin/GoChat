package user

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/rand"
)

type User struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Hash      string
	Salt      string
}

func NewUser(firstname, lastname, email, password string) (User, error) {

	salt, err := generateSalt()
	if err != nil {
		log.Fatalf("Error a generar Salt")
		return User{}, err
	}

	hash := generateHash(password, salt)

	return User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Hash:      hash,
		Salt:      salt,
	}, nil
}

func generateSalt() (string, error) {
	saltBytes := make([]byte, 32)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(saltBytes), nil
}

func generateHash(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write([]byte(salt))
	passwordHash := hash.Sum(nil)
	return hex.EncodeToString(passwordHash)
}

func Authenticate(password, salt, hash string) bool {
	passHash := generateHash(password, salt)
	return hash == passHash
}

type Repository interface {
	Register(user User)
	Login(email, password string) (*User, error)
	GetById(id string) (*User, error)
	GetAll(page int) ([]User, error)
}
