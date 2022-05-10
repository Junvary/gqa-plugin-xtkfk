package model

import (
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type GqaPluginXkDownload struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	Title      string `json:"title" gorm:"comment:资源标题;not null;index"`
	Content    string `json:"content" gorm:"comment:资源内容;type:text;"`
	Attachment string `json:"attachment" gorm:"comment:附件;type:text;"`
}

type RequestAddDownload struct {
	gqaModel.RequestAdd
	Title      string `json:"title"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

type RequestGetDownloadList struct {
	gqaModel.RequestPageAndSort
	Title string `json:"title"`
}
