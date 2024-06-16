package emailx

import (
	"regexp"
)

func CheckEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(email) {
		return true
	}
	return false
}
