package seed

import (
	"biligo/modules/seed/note/api"

	"github.com/gin-gonic/gin"
)

// Route - Seed 模块路由设置文件
func Route(r *gin.RouterGroup) {
	r.GET("/note/", api.NoteList)
	r.GET("/label/", api.LabelList)
}
