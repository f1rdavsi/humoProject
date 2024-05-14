package repository

import (
	"github.com/f1rdavsi/reporter/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (p *AuthPostgres) CreateUser(user models.User) (int, error) {
	if err := p.db.Create(&user).Error; err != nil {
		return -1, err
	}

	return user.ID, nil
}

func (p *AuthPostgres) GetUser(email, username string) (models.User, error) {
	var u models.User
	err := p.db.Where("email = ? AND username = ?", email, username).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}
