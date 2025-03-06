package validators

import (
	"regexp"
	"unicode"

	"pms.pkg/errs"
)

var passwordLengthRegex = regexp.MustCompile(`^[A-Za-z\d]{8,}$`)

func ValidatePassword(password string) error {
	if !passwordLengthRegex.MatchString(password) {
		return errs.ErrInvalidInput{
			Object: "password",
			Reason: "password should contain minimum 8 characters",
		}
	}
	hasLetter := false
	hasNumber := false
	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasNumber = true
		}
		if hasLetter && hasNumber {
			break
		}
	}
	if !hasLetter {
		return errs.ErrInvalidInput{
			Object: "password",
			Reason: "Password should contain at least 1 letter",
		}
	}
	if !hasNumber {
		return errs.ErrInvalidInput{
			Object: "password",
			Reason: "Password should contain at least 1 number",
		}
	}
	return nil
}
