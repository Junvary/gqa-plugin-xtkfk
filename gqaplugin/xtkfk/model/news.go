package model

import (
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type GqaPluginXkNews struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	Title      string `json:"title" gorm:"comment:新闻标题;not null;index"`
	Content    string `json:"content" gorm:"comment:新闻内容;not null;type:text;"`
	Attachment string `json:"attachment" gorm:"comment:附件;type:text;"`
}

type RequestAddNews struct {
	gqaModel.RequestAdd
	Title      string `json:"title"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

type RequestGetNewsList struct {
	gqaModel.RequestPageAndSort
	Title string `json:"title"`
}
