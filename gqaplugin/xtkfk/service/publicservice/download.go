package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetDownloadList(getDownloadList model.RequestGetDownloadList) (err error, news []model.GqaPluginXtkfkDownload, total int64) {
	pageSize := getDownloadList.PageSize
	offset := getDownloadList.PageSize * (getDownloadList.Page - 1)
	db := gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkDownload{})
	var newsList []model.GqaPluginXtkfkDownload
	//配置搜索
	if getDownloadList.Title != "" {
		db = db.Where("title like ?", "%"+getDownloadList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getDownloadList.SortBy, getDownloadList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}
