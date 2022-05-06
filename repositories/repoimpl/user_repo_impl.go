package repoimpl

import (
	"golang-crud/models"
	repo "golang-crud/repositories"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}

func (u *UserRepoImpl) GetList() ([]models.User, error) {
	users := []models.User{}
	if err := u.Db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (u *UserRepoImpl) Create(user models.User) (models.User, error) {
	if err := u.Db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) Find(id string) (models.User, error) {
	user := models.User{}
	if err := u.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) Update(id string, userInput models.User) (models.User, error) {
	user, err := u.Find(id)
	if err != nil {
		return user, err
	}
	if err := u.Db.Model(&user).Updates(userInput).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) Delete(id string) (bool, error) {
	user, err := u.Find(id)
	if err != nil {
		return false, err
	}

	if err := u.Db.Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
