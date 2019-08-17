package api

import (
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// @router /system/ [GET]
func Index(c *gin.Context) {
	util.SuccessResult("Hello Model System").ToJSON(c)
}
