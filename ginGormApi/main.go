package main

import (
	"log"

	"ginGormApi/models"
	"ginGormApi/routers"
	"ginGormApi/storage"
	"github.com/jinzhu/gorm"
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
