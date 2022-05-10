package privateservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"gorm.io/gorm"
)

func GetNewsList(getNewsList model.RequestGetNewsList, username string) (err error, news []model.GqaPluginXtkfkNews, total int64) {
	pageSize := getNewsList.PageSize
	offset := getNewsList.PageSize * (getNewsList.Page - 1)
	var newsList []model.GqaPluginXtkfkNews
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkNews{})); err != nil {
		return err, newsList, 0
	}
	//配置搜索
	if getNewsList.Title != "" {
		db = db.Where("title like ?", "%"+getNewsList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, newsList, 0
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getNewsList.SortBy, getNewsList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}

func EditNews(toEditNews model.GqaPluginXtkfkNews, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkNews{})); err != nil {
		return err
	}
	var news model.GqaPluginXtkfkNews
	if err = db.Where("id = ?", toEditNews.Id).First(&news).Error; err != nil {
		return err
	}
	//err = gqaGlobal.GqaDb.Updates(&toEditNews).Error
	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEditNews.GqaModel),
		map[string]interface{}{
			"title":      toEditNews.Title,
			"content":    toEditNews.Content,
			"attachment": toEditNews.Attachment,
		})).Error
	return err
}

func AddNews(toAddNews model.GqaPluginXtkfkNews, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkNews{})); err != nil {
		return err
	}
	err = db.Create(&toAddNews).Error
	return err
}

func DeleteNewsById(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkNews{})); err != nil {
		return err
	}
	var news model.GqaPluginXtkfkNews
	if err = db.Where("id = ?", id).First(&news).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&news).Error
	return err
}

func QueryNewsById(id uint, username string) (err error, newsInfo model.GqaPluginXtkfkNews) {
	var news model.GqaPluginXtkfkNews
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkNews{})); err != nil {
		return err, news
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&news, "id = ?", id).Error
	return err, news
}
