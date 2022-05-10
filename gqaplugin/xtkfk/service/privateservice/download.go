package privateservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"gorm.io/gorm"
)

func GetDownloadList(getDownloadList model.RequestGetDownloadList, username string) (err error, download []model.GqaPluginXtkfkDownload, total int64) {
	pageSize := getDownloadList.PageSize
	offset := getDownloadList.PageSize * (getDownloadList.Page - 1)
	var downloadList []model.GqaPluginXtkfkDownload
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkDownload{})); err != nil {
		return err, downloadList, 0
	}
	//配置搜索
	if getDownloadList.Title != "" {
		db = db.Where("title like ?", "%"+getDownloadList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getDownloadList.SortBy, getDownloadList.Desc)).Preload("CreatedByUser").Find(&downloadList).Error
	return err, downloadList, total
}

func EditDownload(toEditDownload model.GqaPluginXtkfkDownload, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkDownload{})); err != nil {
		return err
	}
	var download model.GqaPluginXtkfkDownload
	if err = db.Where("id = ?", toEditDownload.Id).First(&download).Error; err != nil {
		return err
	}
	//err = db.Updates(&toEditDownload).Error
	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEditDownload.GqaModel), map[string]interface{}{
		"title":      toEditDownload.Title,
		"content":    toEditDownload.Content,
		"attachment": toEditDownload.Attachment,
	})).Error
	return err
}

func AddDownload(toAddDownload model.GqaPluginXtkfkDownload, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkDownload{})); err != nil {
		return err
	}
	err = db.Create(&toAddDownload).Error
	return err
}

func DeleteDownloadById(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkDownload{})); err != nil {
		return err
	}
	var download model.GqaPluginXtkfkDownload
	if err = db.Where("id = ?", id).First(&download).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&download).Error
	return err
}

func QueryDownloadById(id uint, username string) (err error, downloadInfo model.GqaPluginXtkfkDownload) {
	var download model.GqaPluginXtkfkDownload
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkDownload{})); err != nil {
		return err, download
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&download, "id = ?", id).Error
	return err, download
}
