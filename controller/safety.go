package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 加密密码
func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	checkErr(err)
	return string(hash)
}

// 验证密码
func ComparePasswords(encodePassword string, loginPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(loginPassword))
	if err != nil {
		return false
	}
	return true
}

const (
	SECRETKEY = "3bf84u4hr83x7ru84i73he737y4u" // 私钥
)

type JwtRes struct {
	Valid    int    `json:"valid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

//生成token
func CreateToken(username string) string {
	maxAge := 60 * 60 * 2
	claims := jwt.StandardClaims{
		Audience:  username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
		NotBefore: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

//解析token
func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
