package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetDocumentList(getDocumentList model.RequestGetDocumentList) (err error, news []model.GqaPluginXkDocument, total int64) {
	pageSize := getDocumentList.PageSize
	offset := getDocumentList.PageSize * (getDocumentList.Page - 1)
	db := gqaGlobal.GqaDb.Model(&model.GqaPluginXkDocument{})
	var newsList []model.GqaPluginXkDocument
	//配置搜索
	if getDocumentList.Title != "" {
		db = db.Where("title like ?", "%"+getDocumentList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getDocumentList.SortBy, getDocumentList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}
