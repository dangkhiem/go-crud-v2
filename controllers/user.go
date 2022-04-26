package controllers

import (
	"fmt"
	"golang-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var db *gorm.DB
var validate *validator.Validate

func GetListUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"data": users})
	}
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	validate = validator.New()

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	user := models.User{UserName: input.UserName, Email: input.Email}
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
