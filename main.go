package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/ekonof23/app/controllers"
	"github.com/ekonof23/app/database"
	"github.com/ekonof23/app/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := gorm.Open("mysql", database.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Photo{})

	r := gin.Default()

	r.Use(database.Inject(db))

	router.Initialize(r)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
