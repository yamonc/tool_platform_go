package api

import (
	"biligo/modules/auth"
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// 系统根路径
// @router /system/ [GET]
func Index(c *gin.Context) {
	util.SuccessResult("Hello Model System").ToJSON(c)
}

// 获取当前登录用户
// @router /system/current [GET]
func Current(c *gin.Context) {
	util.SuccessResult(auth.CurrentUser(c)).ToJSON(c)
}
