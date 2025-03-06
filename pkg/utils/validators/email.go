package validators

import (
	"regexp"

	"pms.pkg/errs"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidateEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errs.ErrInvalidInput{
			Object: "email",
			Reason: "invalid",
		}
	}
	return nil
}
