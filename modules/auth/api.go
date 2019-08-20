package auth

import (
	"biligo/mysql"
	"biligo/util"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

// swagger:route POST /api/auth/login
//
// 登录接口
//
// 目前登录采用自定义 token 形式完成，登录后调用接口时需要在请求头里加入 Authorization: [token]
//
//     Responses:
//       200: Result
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if user, e := validateLogin(c, username, password); e == nil {
		token := UserToken{UserId: user.ID}
		token.Token = util.UUID()
		token.ExpiredAt = util.AddTime(time.Now(), "168h")

		mysql.Conn.Create(&token)

		util.SuccessResult(token).ToJSON(c)
	}
}

// 验证用户
func validateLogin(c *gin.Context, username string, password string) (*User, error) {
	if strings.TrimSpace(username) == "" ||
		strings.TrimSpace(password) == "" {
		util.FailResultWithCodeAndMessage(http.StatusUnauthorized, "请输入用户名和密码",
			nil).ToJSONWithHttpStatus(c)

		return nil, errors.New("请输入用户名和密码")
	}

	user := User{}
	mysql.Conn.Where("username = ? and user_status=1", username).First(&user)

	if user.ID == 0 {
		util.FailResultWithCodeAndMessage(http.StatusUnauthorized,
			"用户名行或密码错误", nil).ToJSONWithHttpStatus(c)
		return nil, errors.New("请输入用户名和密码")

	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password),
			[]byte(password))

		if err != nil {
			util.FailResultWithCodeAndMessage(http.StatusUnauthorized,
				"用户名行或密码错误", nil).ToJSONWithHttpStatus(c)
			return nil, errors.New("请输入用户名和密码")

		} else {
			return &user, nil
		}
	}
}
