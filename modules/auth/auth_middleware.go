package auth

import (
	"biligo/config"
	"biligo/constant"
	"biligo/mysql"
	"github.com/gin-gonic/gin"
)

// 登录拦截
// BiliGo 使用请求头中加入 Authorized 参数来获取 token 值
// TODO 先放行，后期加认证 使用 REDIS 优化性能
func AuthMiddleware(c *gin.Context) {
	c.Next()
	//auth := getToken(c)
	//if strings.TrimSpace(auth) == "" {
	//	util.FailResultWithCodeAndMessage(http.StatusUnauthorized,
	//		"请先登录", nil).ToJSONWithHttpStatus(c)
	//	c.Abort()
	//
	//} else {
	//	token := UserToken{}
	//	mysql.Conn.Where("token=? and expired_at > current_timestamp", auth).
	//		First(&token)
	//	if token.Token != "" {
	//		c.Next()
	//
	//	} else {
	//		util.FailResultWithCodeAndMessage(http.StatusUnauthorized,
	//			"请先登录", nil).ToJSONWithHttpStatus(c)
	//		c.Abort()
	//	}
	//}
}

// 获取 Token
func getToken(c *gin.Context) string {
	cookieLoginEnabled := config.GetConfig(constant.CookieLoginEnabled)
	if cookieLoginEnabled == "true" {
		token, _ := c.Cookie(config.GetConfig(constant.CookieLoginName))
		return token

	} else {
		return c.GetHeader(constant.HttpHeaderAuthorized)
	}
}

// 获取当前登录的用户
func CurrentUser(c *gin.Context) *User {
	auth := getToken(c)

	token := UserToken{}
	mysql.Conn.Where("token=?", auth).First(&token)
	user := User{}
	mysql.Conn.Where("id=?", token.UserId).First(&user)
	return &user
}
