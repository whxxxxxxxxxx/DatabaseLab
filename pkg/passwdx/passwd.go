package passwdx

import (
	"DatabaseLab/internal/app/users/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	PassWordCost        = 12
	Active       string = "active"
)

func SetPassword(password string, user *model.Users) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func CheckPassword(password string, user *model.Users) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
