package services

import (
	"golang-crud/models"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(c *gin.Context) (models.User, error)
}
