package model

import (
	"biligo/modules/common/model"
)

type Note struct {
	model.Base
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
