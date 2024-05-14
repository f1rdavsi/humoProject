package service

import (
	"time"

	"github.com/f1rdavsi/reporter/internal/repository"
	"github.com/f1rdavsi/reporter/models"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (t *TransactionService) CreateTransaction(tr models.Transaction) (models.Transaction, error) {
	tr.CreatedAt = time.Now()
	return t.repo.CreateTransaction(tr)
}

func (t *TransactionService) GetTransaction(id, userId int) (models.Transaction, error) {
	return t.repo.GetTransaction(id, userId)
}

func (t *TransactionService) UpdateTransaction(id, userId int, comment string) (models.Transaction, error) {
	tr, err := t.GetTransaction(id, userId)
	if err != nil {
		return tr, err
	}

	tr.Comment = comment
	tr.UpdatedAt = time.Now()

	return t.repo.UpdateTransaction(models.Transaction{})
}

func (t *TransactionService) GetTransactions(userId int) ([]models.Transaction, error) {
	return t.repo.GetTransactions(userId)
}
