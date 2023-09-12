package service

import (
	"biligo/modules/app/model"
	"biligo/mysql"
	"fmt"
)

func TestService() *model.TestModel {
	fmt.Println("11111")
	return &model.TestModel{Name: "Test"}
}

func TestQueryForMap() *[]map[string]interface{} {
	fmt.Println("11111")
	var m []map[string]interface{}
	mysql.Conn.Table("sys_user").Find(&m, "username=?", "admin")
	return &m
}
