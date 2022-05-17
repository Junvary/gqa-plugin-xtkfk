package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetAbbreviateList(getAbbreviateList model.RequestGetAbbreviateList) (err error, news []model.GqaPluginXtkfkAbbreviate, total int64) {
	pageSize := getAbbreviateList.PageSize
	offset := getAbbreviateList.PageSize * (getAbbreviateList.Page - 1)
	db := gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkAbbreviate{})
	var newsList []model.GqaPluginXtkfkAbbreviate
	//配置搜索
	if getAbbreviateList.Title != "" {
		db = db.Where("title like ?", "%"+getAbbreviateList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getAbbreviateList.SortBy, getAbbreviateList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}
