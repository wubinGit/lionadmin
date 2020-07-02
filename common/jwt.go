package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"lionadmin.org/lion/model"
)

var jwtkey = []byte("a_secret_create")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

/**
jwt
 */
func ReleaseToken(admin model.TAdmin) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: admin.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lion.org",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Sprintf("Bearer ")
	tokenString, err := token.SignedString(jwtkey)
	tokenAuth := fmt.Sprintf(tokenString)
	if err != nil {
		return "", err
	}
	return tokenAuth, nil
}

/**
base64
 */
//func (login *Login) createToken() error {
//	loginKey := gostring.StrRandom(40)
//	loginRaw := fmt.Sprintf("%d\t%s", login.UserId, loginKey)
//	var (
//		err error
//		tokenBytes []byte
//	)
//	tokenBytes, err = goaes.Encrypt(syncd.App.CipherKey, []byte(loginRaw))
//	if err != nil {
//		return err
//	}
//	login.Token = gostring.Base64UrlEncode(tokenBytes)
//
//	token := &Token{
//		UserId: login.UserId,
//		Token: loginKey,
//		Expire: int(time.Now().Unix()) + 86400 * 30,
//	}
//	if err := token.CreateOrUpdate(); err != nil {
//		return err
//	}
//	return nil
//}


func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
