package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetHonourList(getHonourList model.RequestGetHonourList) (err error, honour []model.GqaPluginXtkfkHonour, total int64) {
	pageSize := getHonourList.PageSize
	offset := getHonourList.PageSize * (getHonourList.Page - 1)
	db := gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkHonour{})
	var honourList []model.GqaPluginXtkfkHonour
	//配置搜索
	if getHonourList.Title != "" {
		db = db.Where("title like ?", "%"+getHonourList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getHonourList.SortBy, getHonourList.Desc)).Preload("CreatedByUser").Find(&honourList).Error
	return err, honourList, total
}
