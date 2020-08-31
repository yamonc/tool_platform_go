package api

import (
	"biligo/modules/common/page"
	"biligo/modules/seed/note/service"
	"biligo/util"

	"github.com/gin-gonic/gin"
)

func NoteList(c *gin.Context) {
	pagination := page.FromGin(c)
	util.SuccessResult(service.NoteList(pagination)).ToJSON(c)
}

func LabelList(c *gin.Context) {
	util.SuccessResult(service.LabelList()).ToJSON(c)
}
