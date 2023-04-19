package user

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"log"
	"math/rand"
)

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Hash      string
	Salt      string
}

func NewUser(firstname, lastname, password string) (User, error) {

	id := uuid.New().String()

	salt, err := generateSalt()
	if err != nil {
		log.Fatalf("Error a generar Salt")
		return User{}, err
	}

	hash := generateHash(password, salt)

	return User{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
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
	Register(user User, hash, salt string) error
}
