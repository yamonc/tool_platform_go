package service

import (
	"biligo/modules/app/model"
	"biligo/mysql"
)

func TestService() *model.TestModel {
	return &model.TestModel{Name: "Test"}
}

func TestQueryForMap() *[]map[string]interface{} {
	var m []map[string]interface{}
	mysql.Conn.Table("sys_user").Find(&m, "username=?", "admin")
	return &m
}
