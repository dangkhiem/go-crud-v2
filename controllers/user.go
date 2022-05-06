package controllers

import (
	"golang-crud/db"
	"golang-crud/helper"
	"golang-crud/models"
	"golang-crud/repositories"
	"golang-crud/repositories/repoimpl"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	validate     *validator.Validate
	dbConnection *gorm.DB              = db.InitDb()
	userRepo     repositories.UserRepo = repoimpl.NewUserRepo(dbConnection)
)

func GetListUser(c *gin.Context) {
	users, err := userRepo.GetList()
	if err != nil {
		helper.ErrorResponse(c, err, http.StatusNotFound)
		return
	}
	helper.SuccessResponse(c, users)
}

func CreateUser(c *gin.Context) {
	input := models.User{}
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	input.Password, _ = helper.HashPassword(input.Password)
	user, err := userRepo.Create(input)
	if err != nil {
		helper.ErrorResponse(c, err, http.StatusNotFound)
		return
	}
	helper.SuccessResponse(c, user)
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := userRepo.Find(id)

	if err != nil {
		helper.ErrorResponse(c, err, http.StatusNotFound)
		return
	}
	helper.SuccessResponse(c, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	input := models.User{}

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	user, err := userRepo.Update(id, input)
	if err != nil {
		helper.ErrorResponse(c, err, http.StatusNotFound)
		return
	}
	helper.SuccessResponse(c, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	status, err := userRepo.Delete(id)
	if err != nil {
		helper.ErrorResponse(c, err, http.StatusNotFound)
		return
	}

	helper.SuccessResponse(c, status)
}
