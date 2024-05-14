package service

import (
	"fmt"
	utils2 "github.com/f1rdavsi/reporter/pkg/utils"
	"os"
	"time"

	"github.com/f1rdavsi/reporter/internal/repository"
	"github.com/f1rdavsi/reporter/models"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) ExistsAccount(id int) (models.Account, error) {
	return s.repo.ExistsAccount(id)
}

func (s *AccountService) GetAccount(id, userId int) (models.Account, error) {
	return s.repo.GetAccount(id, userId)
}

func (s *AccountService) GetAllAccounts(userId int) ([]models.Account, error) {
	return s.repo.GetAccounts(userId)
}

func (s *AccountService) CreateAccount(userId int, name string, balance float64) (int, error) {
	var account models.Account

	account.UserID = userId

	if utils2.CheckField(name) {
		account.Name = name
	}

	if err := utils2.CheckBalance(balance); err == nil {
		account.Balance = balance
	} else {
		return -1, err
	}

	account.IsActive = true
	account.CreatedAt = time.Now()

	return s.repo.CreateAccount(account)
}

func (s *AccountService) UpdateAccount(id, userId int, name string, balance float64) (models.Account, error) {
	var account models.Account
	account, err := s.repo.GetAccount(id, userId)
	if err != nil {
		return account, err
	}

	if utils2.CheckField(name) {
		account.Name = name
	} else {
		return account, utils2.ErrInvalidAccountName
	}

	if err := utils2.CheckBalance(balance); err == nil {
		account.Balance = balance
	} else {
		return account, err
	}

	account.UpdatedAt = time.Now()

	return s.repo.UpdateAccount(account)
}

func (s *AccountService) DeleteAccount(id, userId int) error {
	return s.repo.DeleteAccount(id, userId)
}

func (s *AccountService) RestoreAccount(id, userId int) error {
	return s.repo.RestoreAccount(id, userId)
}

func (s *AccountService) ChangePictureAccount(id, userId int, filepath string) (models.Account, error) {
	account, err := s.GetAccount(id, userId)
	if err != nil {
		return account, err
	}

	if err := os.Remove(fmt.Sprintf("./files/layouts/%s", account.Picture)); err != nil {
		return account, err
	}

	account.Picture = filepath
	account.UpdatedAt = time.Now()

	return s.repo.UpdateAccount(account)
}

func (s *AccountService) UploadAccountPicture(id, userId int, filepath string) (models.Account, error) {
	account, err := s.GetAccount(id, userId)
	if err != nil {
		return account, err
	}
	account.Picture = filepath
	account.UpdatedAt = time.Now()

	return s.repo.UpdateAccount(account)
}
