package tool

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword 单向hash加密，生成加密密文
func GenerateFromPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if nil != err {
		return ""
	}
	return string(hash)
}

// CompareHashAndPassword hash验证是否匹配
func CompareHashAndPassword(hash string, rawStr string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawStr))
	return nil == err
}
