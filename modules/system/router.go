package system

import (
	"biligo/modules/system/api"
	"github.com/gin-gonic/gin"
)

func RouteSys(r *gin.RouterGroup) {
	r.GET("/", api.Index)
	r.GET("/current", api.Current)
}
