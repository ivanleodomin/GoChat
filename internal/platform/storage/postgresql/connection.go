package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(user, pass, dbname, host, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, pass, dbname, host, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
