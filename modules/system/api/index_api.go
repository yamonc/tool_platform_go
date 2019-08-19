package api

import (
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// 系统根路径
// @router /api/system/ [GET]
func Index(c *gin.Context) {
	util.SuccessResult("Hello Model System").ToJSON(c)
}
