package main

import (
	"log"

	"github.com/msinwolc/routers"

	"github.com/msinwolc/models"
)

func main() {
	models.GetDB()

	models.OpenDB()

	defer models.CloseDB()
	r := routers.Routers()
	// fmt.Println(r)
	if err := r.Run(); err != nil {
		log.Fatalf("startup server failed: %v", err)
	}
}
