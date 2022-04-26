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

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.UserRouter(r)

	r.Run(":9999")
}
