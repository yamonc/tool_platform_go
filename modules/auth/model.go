package auth

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Password    string    `json:"password"`
	UserStatus  uint      `json:"userStatus"`
	DeptId      uint      `json:"deptId"`
	LoginAt     time.Time `json:"loginAt"`
}

func (User) TableName() string {
	return "sys_user"
}

type UserToken struct {
	Token     string    `json:"token" gorm:"PRIMARY_KEY"`
	UserId    uint      `json:"userId"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (UserToken) TableName() string {
	return "sys_user_token"
}
