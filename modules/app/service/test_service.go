package service

import "biligo/modules/app/model"

func TestService() *model.TestModel {
	return &model.TestModel{Name: "Test"}
}
