package repository

import (
	"github.com/f1rdavsi/reporter/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) ExistsAccount(id int) (models.Account, error) {
	var a models.Account
	err := r.db.Model(&models.Account{}).Where("id = ? AND is_active = ?", id, true).First(&a).Error
	if err != nil {
		return a, err
	}

	return a, nil
}

func (r *AccountRepository) GetAccounts(userId int) ([]models.Account, error) {
	var accounts []models.Account

	err := r.db.Where("user_id = ? AND is_active = ?", userId, true).Find(&accounts).Error
	if err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (r *AccountRepository) GetAccount(id, userId int) (models.Account, error) {
	var account models.Account
	err := r.db.Where("id = ? AND user_id = ? AND is_active = ?", id, userId, true).First(&account).Error
	if err != nil {
		return account, err
	}

	return account, nil
}

func (r *AccountRepository) CreateAccount(account models.Account) (int, error) {
	if err := r.db.Create(&account).Error; err != nil {
		return -1, err
	}

	return account.ID, nil
}

func (r *AccountRepository) UpdateAccount(account models.Account) (models.Account, error) {
	if err := r.db.Save(&account).Error; err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func (r *AccountRepository) DeleteAccount(id, userId int) error {
	return r.db.
		Model(&models.Account{}).
		Where("id = ? AND user_id = ? AND is_active = ?", id, userId, false).
		Update("is_active", false).
		Error
}

func (r *AccountRepository) RestoreAccount(id, userId int) error {
	return r.db.
		Model(&models.Account{}).
		Where("id = ? AND user_id = ? AND is_active = ?", id, userId, false).
		Update("is_active", true).
		Error
}
