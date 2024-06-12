package tool

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtConfig struct {
	Secret string `json:"secret,omitempty" yaml:"secret"`
	Expire int64  `json:"expire,omitempty" yaml:"expire"`
}

type Payload struct {
	Id                   int64  `json:"number,omitempty"`
	Username             string `json:"string,omitempty"`
	jwt.RegisteredClaims `json:"-"`
}

// GenToken 生成Token 如果 expiration <= 0 则为永不过期
func GenToken(secret string, userId int64, username string, expiration int64) (token string, err error) {
	var expires *jwt.NumericDate
	if expiration <= 0 {
		expires = jwt.NewNumericDate(time.Now().AddDate(100, 0, 0))
	} else {
		expires = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expiration)))
	}
	claims := Payload{
		userId, username,
		jwt.RegisteredClaims{
			ExpiresAt: expires,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}

func ParseJwt(secret string, token string) (payload *Payload, err error) {
	t, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := t.Claims.(*Payload); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
