package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name;type:varchar(255);uniqueIndex;not null;"`
	Password string `gorm:"type:varchar(255);not null"`
}

//在post的时候上传body必须按照json字段来填写
//之前好多次都是因为没匹配导致409

type UserLogin struct {
	UserName string `json:"username,userName,user_name"`
	Password string `json:"password,passWord,pass_word"`
}
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
