package service

import (
	"os"
	"time"

	"github.com/f1rdavsi/reporter/internal/repository"
	"github.com/f1rdavsi/reporter/models"
	"github.com/xuri/excelize/v2"
)

type Authorization interface {
	CreateUser(fistname, lastname, username, email, password string) (int, error)
	GenerateToken(email, username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserById(id int) (models.User, error)
	UpdateUser(id int, firstname, lastname, email, username, password string) (models.User, error)
	DeleteUserById(id int) error
	RestoreUserById(id int) error
	UpdatePictureUser(id int, filepath string) (models.User, error)
	UploadUserPicture(id int, filepath string) (models.User, error)
}

type Account interface {
	ExistsAccount(id int) (models.Account, error)
	GetAccount(id, userId int) (models.Account, error)
	GetAllAccounts(userId int) ([]models.Account, error)
	CreateAccount(userId int, name string, balance float64) (int, error)
	UpdateAccount(id, userId int, name string, balance float64) (models.Account, error)
	DeleteAccount(id, userId int) error
	RestoreAccount(id, userId int) error
	ChangePictureAccount(id, userId int, filepath string) (models.Account, error)
	UploadAccountPicture(id, userId int, filepath string) (models.Account, error)
}

type Category interface {
	CreateCategory(name, description string, price float64) (models.Category, error)
	GetCategories() ([]models.Category, error)
	GetCategory(id int) (models.Category, error)
	UpdateCategory(id int, name, description string, price float64) (models.Category, error)
	UploadPictureCategory(id int, filepath string) (models.Category, error)
	ChangePictureCategory(id int, filepath string) (models.Category, error)
	DeleteCategory(id int) error
	RestoreCategory(id int) error
}

type Transaction interface {
	CreateTransaction(tr models.Transaction) (models.Transaction, error)
	GetTransaction(id, userId int) (models.Transaction, error)
	UpdateTransaction(id, userId int, comment string) (models.Transaction, error)
	GetTransactions(userId int) ([]models.Transaction, error)
}

type Report interface {
	GetCSVReport(
		userFrom, userTo models.User,
		accountFrom, accountTo models.Account,
		tr []models.Transaction,
	) (*os.File, error)
	GetExcelReport(
		userFrom, userTo models.User,
		accountFrom, accountTo models.Account,
		tr []models.Transaction,
	) (*excelize.File, error)
	GetReport(
		FromID, ToID int, ToType string,
		Limit, Page int, Type string,
		From, To time.Time,
	) ([]models.Transaction, error)
}

type Service struct {
	Authorization
	User
	Account
	Category
	Transaction
	Report
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		User:          NewUserService(repository.User),
		Account:       NewAccountService(repository.Account),
		Category:      NewCategoryService(repository.Category),
		Transaction:   NewTransactionService(repository.Transaction),
		Report:        NewReportService(repository.Report),
	}
}
