package api

import (
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /api/system/
//
// system 模块根接口
//
//     Responses:
//       200: Result
func Index(c *gin.Context) {
	util.SuccessResult("Hello Model System").ToJSON(c)
}
