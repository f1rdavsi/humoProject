package service

import (
	"fmt"
	"github.com/f1rdavsi/reporter/pkg/utils"
	"os"
	"time"

	"github.com/f1rdavsi/reporter/internal/repository"
	"github.com/f1rdavsi/reporter/models"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUserById(id int) (models.User, error) {
	return u.repo.GetUserById(id)
}

func (u *UserService) DeleteUserById(id int) error {
	return u.repo.DeleteUserById(id)
}

func (u *UserService) RestoreUserById(id int) error {
	return u.repo.RestoreUserById(id)
}

func (u *UserService) UpdateUser(
	id int,
	firstname, lastname, email, username, password string,
) (models.User, error) {
	user, err := u.GetUserById(id)
	if err != nil {
		return user, err
	}

	if utils.CheckField(firstname) {
		user.Firstname = firstname
	}

	if utils.CheckField(lastname) {
		user.Lastname = lastname
	}

	if utils.CheckField(email) {
		user.Email = email
	}

	if utils.CheckField(username) {
		user.Username = username
	}

	if utils.IsChangeablePassword(user.Password, password) == nil {
		pwd, err := utils.GeneratePassword(password)
		if err != nil {
			return user, err
		}

		user.Password = pwd
	}

	user.CreatedAt = time.Now()

	return u.repo.UpdateUser(user)
}

func (u *UserService) UpdatePictureUser(id int, filepath string) (models.User, error) {
	user, err := u.GetUserById(id)
	if err != nil {
		return user, err
	}

	if err := os.Remove(fmt.Sprintf("./files/layouts/%s", user.Picture)); err != nil {
		return user, err
	}

	user.Picture = filepath
	user.UpdatedAt = time.Now()

	return u.repo.UpdateUser(user)
}

func (u *UserService) UploadUserPicture(id int, filepath string) (models.User, error) {
	user, err := u.GetUserById(id)
	if err != nil {
		return user, err
	}

	user.Picture = filepath
	user.UpdatedAt = time.Now()

	return u.repo.UpdateUser(user)
}
