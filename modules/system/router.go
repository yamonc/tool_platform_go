package system

import (
	"biligo/modules/system/api"
	"github.com/gin-gonic/gin"
)

/// system 模块路由设置文件

func Route(r *gin.RouterGroup) {
	r.GET("/", api.Index)

	r.GET("/user/current", api.UserCurrent)
	r.GET("/user/list", api.UserList)
}
