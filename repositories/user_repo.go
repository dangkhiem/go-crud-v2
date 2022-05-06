package repositories

import "golang-crud/models"

type UserRepo interface {
	GetList() ([]models.User, error)
	Create(u models.User) (models.User, error)
	Find(id string) (models.User, error)
	Update(id string, u models.User) (models.User, error)
	Delete(id string) (bool, error)
}
