package app

import (
	"biligo/modules/app/api"
	"github.com/gin-gonic/gin"
)

// 在这里注册你的 app 所有 api 路由
func RouteApp(r *gin.RouterGroup) {

	r.GET("/test", api.Test)
}
