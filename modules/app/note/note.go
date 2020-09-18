package note

import (
	"biligo/modules/common"
	"biligo/mysql"
	"biligo/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/////////////////////////////////////////////////////////////////
// 也可以将 api、model、service 写在一个文件中，看你怎么取舍了 //
/////////////////////////////////////////////////////////////////

// API

func NoteList(c *gin.Context) {
	pagination := common.NewPageFromGin(c)
	util.SuccessResult(queryNoteList(pagination)).ToJSON(c)
}

func LabelList(c *gin.Context) {
	util.SuccessResult(queryLabelList()).ToJSON(c)
}

// SERVICE

// NoteList - 分页查询 note
func queryNoteList(pagination *common.Pagination) *[]Note {
	notes := []Note{}
	mysql.Conn.Find(&notes)
	return &notes
}

// LabelList - 返回所有的 Label
func queryLabelList() *[]Label {
	labels := []Label{}
	mysql.Conn.Find(&labels)
	return &labels
}

// MODEL

type Note struct {
	gorm.Model
	NoteTitle   string `json:"noteTitle"`
	NoteType    string `json:"noteType"`
	NoteContent string `json:"noteContent"`
	SourceURL   string `json:"sourceUrl"`
	Author      string `json:"author"`
	ShareToken  string `json:"shareToken"`
	CategoryID  uint   `json:"categoryId"`
	OrgID       uint   `json:"orgId"`
}

func (Note) TableName() string {
	return "seed_note"
}

type Label struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (Label) TableName() string {
	return "seed_label"
}

type NoteLabel struct {
	NoteID  uint `json:"noteId"`
	LabelID uint `json:"labelId"`
}
