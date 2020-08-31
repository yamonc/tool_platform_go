package service

import (
	"biligo/modules/common/page"
	"biligo/modules/seed/note/model"
	"biligo/mysql"
)

// NoteList - 分页查询 note
func NoteList(pagination *page.Pagination) *[]model.Note {
	notes := []model.Note{}
	mysql.Conn.Find(&notes)
	return &notes
}

// LabelList - 返回所有的 Label
func LabelList() *[]model.Label {
	labels := []model.Label{}
	mysql.Conn.Find(&labels)
	return &labels
}
