用户注册时将密码加密后存储在数据库中，使用Bcrypt算法，对于同一个密码，每次生成的密文都不同，无法通过直接比对密文来反推明文，因此可以有效抵御彩虹表攻击：
```go
package controller

import (
	"golang.org/x/crypto/bcrypt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

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
```
注册时提供：用户名、密码、车牌号
