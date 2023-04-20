package main

import (
	boostrap "app-go/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := boostrap.Run(); err != nil {
		log.Fatal(err)
	}
}
