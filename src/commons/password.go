package commons

import (
	"log"
	"math/rand"
	"regexp"
	"unicode"

	config "github.com/soheilsirousi/golang-web-api/src/configs"
)

var (
	passwordConfig = config.GetConfig().Password
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = []string{lowerCharSet, upperCharSet, specialCharSet, numberSet}
	specialRegex   = `[^((0-9)|(a-z)|(A-Z)|\s)]`
)

func CheckPassword(password string) bool {
	if !(passwordConfig.InputNum && HasNumber(password)) {
		return false
	}
	if !(passwordConfig.InputSpecials && HasSpecial(password)) {
		return false
	}
	if !(passwordConfig.LowerChars && HasLower(password)) {
		return false
	}
	if !(passwordConfig.UpperChars && HasUpper(password)) {
		return false
	}

	return true
}

func GeneratePassword() string {
	passwordCheck := false
	var password string
	for !passwordCheck {
		for i := 0; i < passwordConfig.MaxLength; i++ {
			indexCharSet := rand.Intn(len(allCharSet))
			charSet := allCharSet[indexCharSet]
			indexChar := rand.Intn(len(charSet))
			password += string(charSet[indexChar])
		}
		if CheckPassword(password) {
			passwordCheck = true
		}
	}

	return password
}

func HasNumber(password string) bool {
	for _, letter := range password {
		if unicode.IsDigit(letter) {
			return true
		}
	}
	return false
}

func HasSpecial(password string) bool {
	ok, err := regexp.MatchString(specialRegex, password)
	if err != nil {
		log.Printf("can not match the regex")
	}
	return ok
}

func HasLower(password string) bool {
	for _, letter := range password {
		if unicode.IsLower(letter) {
			return true
		}
	}
	return false
}

func HasUpper(password string) bool {
	for _, letter := range password {
		if unicode.IsUpper(letter) {
			return true
		}
	}
	return false
}
