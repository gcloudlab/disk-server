package define

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// JwtKey
type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "gcloud-key"
var MailPassword = os.Getenv("MailPassword")

var RedisPassword = os.Getenv("RedisPassword")

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间（s）
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = os.Getenv("TencentSecretKey")
var TencentSecretID = os.Getenv("TencentSecretID")
var CosBucket = "https://gcloud-1303456836.cos.ap-chengdu.myqcloud.com"
var CosFolderName = "gcloud"

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2000-01-01 00:00:01"

var TokenExpire = 36000
var RefreshTokenExpire = 72000
