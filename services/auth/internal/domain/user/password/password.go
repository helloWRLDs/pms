package password

import (
	"golang.org/x/crypto/bcrypt"
	"pms.pkg/errs"
)

type Password string

func New(password string) (Password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errs.ErrInvalidInput{
			Object: "password",
			Reason: err.Error(),
		}
	}
	return Password(string(hashed)), nil
}

func (p Password) Matches(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
	return err == nil
}
