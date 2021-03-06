package main

import (
	"golang-crud/db"
	"golang-crud/models"
	"golang-crud/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := db.InitDb()
	db.AutoMigrate(&models.User{})

	router.UserRouter(r)
	router.AuthRouter(r)

	r.Run(":9999")
}
