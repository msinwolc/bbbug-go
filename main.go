package main

import (
	"log"

	"github.com/msinwolc/config"
	"github.com/msinwolc/routers"

	"github.com/msinwolc/models"
)

func main() {
	models.GetDB()

	models.OpenDB()

	config.ConnectRDB()

	defer models.CloseDB()
	r := routers.Routers()
	if err := r.Run(); err != nil {
		log.Fatalf("startup server failed: %v", err)
	}
}
