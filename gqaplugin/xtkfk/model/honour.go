package model

import (
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type GqaPluginXtkfkHonour struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	Title      string `json:"title" gorm:"comment:新闻标题;not null;index"`
	Attachment string `json:"attachment" gorm:"comment:附件;type:text;"`
}

type RequestAddHonour struct {
	gqaModel.RequestAdd
	Title      string `json:"title"`
	Attachment string `json:"attachment"`
}

type RequestGetHonourList struct {
	gqaModel.RequestPageAndSort
	Title string `json:"title"`
}
