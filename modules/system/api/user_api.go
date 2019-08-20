package api

import (
	"biligo/modules/auth"
	"biligo/modules/system/service"
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /api/system/user/current
//
// 获取当前登录用户接口
//     Responses:
//       200: Result
func UserCurrent(c *gin.Context) {
	util.SuccessResult(auth.CurrentUser(c)).ToJSON(c)
}

// swagger:route GET /api/system/user/list
//
// 获取用户列表接口
//
//     Responses:
//       200: Result
func UserList(c *gin.Context) {
	util.SuccessResult(service.GetUserList()).ToJSON(c)
}
