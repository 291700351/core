package tool

import "regexp"

var (
	regMobile = regexp.MustCompile(`^1[3-9]\d{9}$`)
	regEmail  = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

func IsMobile(str string) bool {
	return regMobile.MatchString(str)
}
func IsEmail(str string) bool {
	return regEmail.MatchString(str)
}
