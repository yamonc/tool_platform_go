package api

import (
	"biligo/modules/app/service"
	"biligo/mysql"
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /api/app/test
//
// app 模块 test 接口
//
//     Responses:
//       200: Result
func Test(c *gin.Context) {
	util.SuccessResult(service.TestService()).ToJSON(c)
}

// 测试 QueryForMap 函数
func TestQueryForMap(c *gin.Context) {
	maps, err := mysql.QueryForMap("select * from sys_user")
	if err != nil {
		util.FailResultWithMessage(err.Error(), nil).ToJSON(c)
	} else {
		util.SuccessResult(maps).ToJSON(c)
	}
}
