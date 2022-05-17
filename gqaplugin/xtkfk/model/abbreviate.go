package model

import gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"

type GqaPluginXtkfkAbbreviate struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	Top     string `json:"top" gorm:"comment:大分类;not null;index"`
	From    string `json:"from" gorm:"comment:来源;"`
	Title   string `json:"title" gorm:"comment:标题;index"`
	Content string `json:"content" gorm:"comment:内容; type:text"`
}

type RequestAddAbbreviate struct {
	gqaModel.RequestAdd
	Top     string `json:"top"`
	From    string `json:"from"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type RequestGetAbbreviateList struct {
	gqaModel.RequestPageAndSort
	Top   string `json:"top"`
	Title string `json:"title"`
}
