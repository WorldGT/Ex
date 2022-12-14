package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type Userclaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

// 验证码长度
var CodeLength = 6

// 验证码过期时间(s)
var CodeExpire = 300

var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 3600
var RefreshTokenExpire = 7200
