package api

import (
	"biligo/modules/app/service"
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// @router /api/app/test [GET]
func Test(c *gin.Context) {
	util.SuccessResult(service.TestService()).ToJSON(c)
}
