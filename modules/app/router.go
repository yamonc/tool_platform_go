package app

import (
	"biligo/modules/app/api"
	"biligo/modules/app/car"
	"biligo/modules/app/maintain"
	"biligo/modules/app/note"
	"biligo/modules/app/password"
	"github.com/gin-gonic/gin"
)

/// app 模块路由设置文件

// 在这里注册你的 app 所有 api 路由
func Route(r *gin.RouterGroup) {

	r.GET("/test", api.TestIndex)
	r.GET("/test/map", api.TestQueryForMap)

	r.GET("/note/", note.NoteList)
	r.GET("/label/", note.LabelList)

	r.GET("/password/list", password.PasswordList)
	r.GET("/password/get/:id", password.GetPasswordById)
	r.POST("/password/save", password.SavePassword)
	r.PUT("/password/:id", password.UpdatePassword)
	r.DELETE("/password/:id", password.DeletePassword)

	r.POST("/car/save", car.SaveCarInfo)
	r.GET("/car/list", car.CarInfoList)
	r.GET("/car/get/:id", car.GetCarInfoById)
	r.PUT("/car/:id", car.UpdateCarInfoById)
	r.DELETE("/car/:id", car.DeleteCarInfoById)
	r.GET("/car/list/names", car.GetAllCarIdAndName)

	r.POST("/maintain/save", maintain.SaveCarMaintainInfo)
	r.GET("/maintain/list", maintain.CarMaintainInfoList)
	r.GET("/maintain/get/:id", maintain.GetMaintainInfoById)
	r.DELETE("/maintain/:id", maintain.DeleteMaintainInfoById)

	r.GET("/car/test", car.TestCar)

}
