package model

import (
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type GqaPluginXkDocument struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	Title      string `json:"title" gorm:"comment:文档标题;not null;index"`
	Content    string `json:"content" gorm:"comment:文档内容;type:text;"`
	Attachment string `json:"attachment" gorm:"comment:附件;type:text;"`
}

type RequestAddDocument struct {
	gqaModel.RequestAdd
	Title      string `json:"title"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

type RequestGetDocumentList struct {
	gqaModel.RequestPageAndSort
	Title string `json:"title"`
}
