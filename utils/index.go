package utils

import "golang.org/x/crypto/bcrypt"

func MD5(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", err
	}

	return string(hash), err
}
