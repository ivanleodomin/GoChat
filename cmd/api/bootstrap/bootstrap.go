package boostrap

import (
	"app-go/internal/platform/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Run() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	srv := server.New(host, port)
	return srv.Run()
}
