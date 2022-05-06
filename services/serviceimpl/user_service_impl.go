package serviceimpl

import (
	"golang-crud/models"
	"golang-crud/repositories"
	"golang-crud/services"

	"github.com/gin-gonic/gin"
)

type UserServiceImpl struct {
	UserRepo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) services.UserService {
	return &UserServiceImpl{
		UserRepo: repo,
	}
}

func (s *UserServiceImpl) Create(c *gin.Context) (models.User, error) {
	input := models.User{}
	c.ShouldBindJSON(&input)
	user, err := s.UserRepo.Create(input)
	return user, err
}
