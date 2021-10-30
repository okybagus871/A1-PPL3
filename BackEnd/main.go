package main

import (
	"cobaApi/handler"
	"cobaApi/user"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB
var err error

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=user port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB error")
	}

	db.AutoMigrate(&user.User{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	router.GET("/getUser", userHandler.GetUser)
	router.POST("/addUser", userHandler.PostUserHandler)

	router.Run()
}
