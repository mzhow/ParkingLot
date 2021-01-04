package controller

import (
	"ParkingLot/dao"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mojocn/base64Captcha"
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
	SECRETKEY = "3bf84hr2g1xr4i96pe7v5y" // 私钥
)

type ResData struct {
	Valid    int    `json:"valid"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Message  string `json:"message"`
}

//生成token
func CreateToken(username string) string {
	maxAge := 60 * 60
	claims := jwt.StandardClaims{
		Audience:  username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
		NotBefore: time.Now().Unix(),
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

func checkToken(token string) bool {
	if token == "null" || token == "" {
		return false
	}
	parseToken, err := ParseToken(token)
	if err != nil {
		return false
	}
	username := parseToken.Audience
	// token错误
	if dao.CheckUsernameValid(username) {
		return true
	}
	return false
}

var store = base64Captcha.DefaultMemStore

type Captcha struct {
	Id   string `json:"id"`
	B64s string `json:"b64s"`
}

//  获取验证码
func GetCaptcha() (string, string) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		checkErr(err)
		return "", ""
	}
	return id, b64s
}

// 验证验证码
func VerifyCaptcha(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}
