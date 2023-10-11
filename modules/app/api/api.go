package api

import (
	"biligo/modules/app/service"
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /api/app/test
//
// app 模块 test 接口
//
//	Responses:
//	  200: Result
func TestIndex(c *gin.Context) {
	util.SuccessResult(service.TestService()).ToJSON(c)
}

func TestQueryForMap(c *gin.Context) {
	util.SuccessResult(service.TestQueryForMap()).ToJSON(c)
}
