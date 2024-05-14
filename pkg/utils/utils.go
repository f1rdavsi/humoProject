package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CheckField(field string) bool {
	if field == "" || field == " " || len(field) > 50 {
		return false
	}

	return true
}

func CheckBalance(balance float64) error {
	if balance < 0 || balance > 5000 {
		return ErrTooMuchBalance
	}

	return nil
}

func IsChangeablePassword(userPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
}

func GeneratePassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pwd), nil
}

func CheckPassword(hashPassword, password string) bool {
	return IsChangeablePassword(hashPassword, password) == nil
}

func GenFilenameWithDir(filename string) (string, error) {
	validExtensions := []string{".jpg", ".jpeg", ".png"}

	fileExt := strings.ToLower(filepath.Ext(filename))

	for _, ext := range validExtensions {
		if ext == fileExt {
			timeSign := fmt.Sprintf("%d", time.Now().UnixNano())
			filePath := fmt.Sprintf("%s_%s", timeSign, filename)
			filePath = strings.Replace(filePath, " ", "", 111)

			return filePath, nil
		}
	}

	return "", ErrUnsupportedMediaType
}

func GetContentType(filePath string) string {
	return fmt.Sprintf("image/%s", strings.Replace(filepath.Ext(filePath), ".", "", 111))
}
