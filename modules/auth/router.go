package auth

import (
	"github.com/gin-gonic/gin"
)

func RouteAuth(r *gin.Engine) {
	r.POST("/api/auth/login", DoLogin)
}
