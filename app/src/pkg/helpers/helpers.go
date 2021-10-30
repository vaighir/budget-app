package helpers

import (
	"fmt"
	"regexp"
)

const (
	minUsernameLength = 5
	maxUsernameLength = 50
	minPasswordLength = 10
)

func CheckUsername(username string) (bool, string) {
	check := false
	msg := "ok"

	if len(username) == 0 {
		msg = "Username cannot be empty"
	} else if len(username) < minUsernameLength {
		msg = fmt.Sprintf("Username has to be at least %d characters long", minUsernameLength)
	} else if len(username) > maxUsernameLength {
		msg = fmt.Sprintf("Username can't be longer than %d", maxUsernameLength)
	} else if regexp.MustCompile(`\s`).MatchString(username) {
		msg = "Username can't contain whitespace"
	} else {
		check = true
	}

	return check, msg
}

func CheckPassword(password string, repeatPassword string) (bool, string) {
	check := false
	msg := "ok"

	if len(password) == 0 {
		msg = "Password cannot be empty"
	} else if len(password) < minPasswordLength {
		msg = fmt.Sprintf("Password has to be at least %d characters long", minPasswordLength)
	} else if password != repeatPassword {
		msg = "Passwords have to match"
	} else if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		msg = "Password must container lower and upper case characters"
	} else if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		msg = "Password must container lower and upper case characters"
	} else if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		msg = "Password must container numbers"
	} else if regexp.MustCompile(`\s`).MatchString(password) {
		msg = "Password can't contain whitespace"
	} else {
		check = true
	}

	return check, msg
}
