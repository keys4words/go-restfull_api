package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/keys4words/go-restfull_api/ginGormApi/models"
	"github.com/keys4words/go-restfull_api/ginGormApi/routers"
	"github.com/keys4words/go-restfull_api/ginGormApi/storage"
	_ "github.com/lib/pq"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=restapi sslmode=disable")
	if err != nil {
		log.Println("error in processing with DB", err)
	}
	defer storage.DB.Close()
	storage.DB.AutoMigrate(&models.Article{})

	r := routers.SetupRouter()
	r.Run()
}
