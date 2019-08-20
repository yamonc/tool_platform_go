package auth

import (
	"github.com/gin-gonic/gin"
)

/// auth 模块路由设置文件

func RouteAuth(r *gin.Engine) {
	r.POST("/api/auth/login", Login)
}
