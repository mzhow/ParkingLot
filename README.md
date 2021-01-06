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
注册时提供：用户名、密码、车牌号，一个用户对应一个车牌号，用户名不能与其他用户相同。每人只能存在一个未完成订单。

查booking时要看是否valid，21:00以后不能选今天的停车场，只能预约明天的（0:00-21:00显示今天和明天的车位信息，但只能预约今天的，21:00-23:59只显示明天的车位信息，21:00-22:00不能预约，22:00-23:59可以预约第二天的）

只能选目前为空的车位，不然有可能下单了进不去。

创建订单时将费用添加到user的fee中，取消订单则减去相应的费用

停车场开放时间：早8点-晚21:00，8:00开始可以进停车场，21:00以后不允许进、只允许出，超过22:00没有出停车场算第二天，直接扣一天的费用

每天22:00抢第二天车位



|        | user valid                       | user fee                                         | bookingvalid                     |
| ------ | -------------------------------- | ------------------------------------------------ | -------------------------------- |
| entry  | \                                | \                                                | \                                |
| out    | fee==0且时间超过endtime后置1放出 | 超过endtime后更新fee若fee不为0不让出，未超时让出 | fee==0且时间超过endtime后置0放出 |
| pay    | 如果在停车场内置0，不在置1       | 置0                                              | 在停车场内置1，不在置0           |
| cancel | 时间未超过starttime置1           | 时间未超过starttime置0                           | 时间未超过starttime置0           |

先更新uservalid和bookingvalid->找spot（empty, valid）->找carid->insert booking->update user

订单还未开始时（未到start_time）才允许取消订单：booking->valid=0，fee->0；

out->若fee=0-> booking valid=0 uservalid=1

fee只管更新费用

想要预约第二天的停车场，必须调用out且fee为0，只有out会更新uservalid和bookingvalid

支付后，fee清零，valid

booking valid变0,user valid才变1



entry和out只进和出

付费：如果时间大于结束时间，就更新费用，user valid置1，book valid置0，否则直接更新fee=0

下单时看fee，fee为0则允许选，fee不为0则提示要付掉费用再选，