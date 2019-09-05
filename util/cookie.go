package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

// 设置 Cookie
func SetCookie(c *gin.Context, name string, value string, expires int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   expires,
		Path:     "/",
		HttpOnly: true,
	})
}
