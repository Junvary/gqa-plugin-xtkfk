package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetNewsList(getNewsList model.RequestGetNewsList) (err error, news []model.GqaPluginXkNews, total int64) {
	pageSize := getNewsList.PageSize
	offset := getNewsList.PageSize * (getNewsList.Page - 1)
	db := gqaGlobal.GqaDb.Model(&model.GqaPluginXkNews{})
	var newsList []model.GqaPluginXkNews
	//配置搜索
	if getNewsList.Title != "" {
		db = db.Where("title like ?", "%"+getNewsList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getNewsList.SortBy, getNewsList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}
