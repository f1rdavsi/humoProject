package repository

import (
	"github.com/f1rdavsi/reporter/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := u.db.Where("id = ? AND is_active = ?", id, true).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(user models.User) (models.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *UserRepository) DeleteUserById(id int) error {
	return u.db.
		Model(&models.User{}).
		Where("id = ? AND is_active = TRUE", id).
		Update("is_active", false).
		Error
}

func (u *UserRepository) RestoreUserById(id int) error {
	return u.db.
		Model(&models.User{}).
		Where("id = ? AND is_active = FALSE", id).
		Update("is_active", true).
		Error
}
