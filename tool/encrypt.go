package tool

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword 单向hash
func GenerateFromPassword(str string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if nil != err {
		return ""
	}
	return string(hash)
}

func CompareHashAndPassword(hash string, rawStr string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawStr))
	return nil == err
}
