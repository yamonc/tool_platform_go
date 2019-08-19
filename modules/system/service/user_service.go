package service

import (
	"biligo/modules/auth"
	"biligo/mysql"
)

func GetUserList() *[]auth.User {
	users := []auth.User{}
	mysql.Conn.Find(&users)
	return &users
}
