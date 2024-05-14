package repository

import (
	"github.com/f1rdavsi/reporter/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (t *TransactionRepository) CreateTransaction(tr models.Transaction) (models.Transaction, error) {
	tx := t.db.Begin()

	if err := tx.Create(&tr).Error; err != nil {
		tx.Rollback()
		return models.Transaction{}, err
	}

	tx.Commit()
	return tr, nil
}

func (t *TransactionRepository) GetTransaction(id, userId int) (models.Transaction, error) {
	var tr models.Transaction
	err := t.db.Model(&models.Transaction{}).Where("id = ? AND from = ?", id, userId).Find(&tr).Error
	if err != nil {
		return tr, err
	}

	return tr, nil
}

func (t *TransactionRepository) UpdateTransaction(tr models.Transaction) (models.Transaction, error) {
	tx := t.db.Begin()

	if err := tx.Save(&tr).Error; err != nil {
		tx.Rollback()
		return models.Transaction{}, err
	}

	tx.Commit()
	return tr, nil
}

func (t *TransactionRepository) GetTransactions(userId int) ([]models.Transaction, error) {
	var tr []models.Transaction

	err := t.db.Model(&models.Transaction{}).Where("from = ?", userId).Find(&tr).Error
	if err != nil {
		return tr, err
	}

	return tr, nil
}
