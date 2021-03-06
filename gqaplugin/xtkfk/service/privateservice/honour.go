package privateservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"gorm.io/gorm"
)

func GetHonourList(getHonourList model.RequestGetHonourList, username string) (err error, honour []model.GqaPluginXtkfkHonour, total int64) {
	pageSize := getHonourList.PageSize
	offset := getHonourList.PageSize * (getHonourList.Page - 1)
	var honourList []model.GqaPluginXtkfkHonour
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkHonour{})); err != nil {
		return err, honourList, 0
	}
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

func EditHonour(toEditHonour model.GqaPluginXtkfkHonour, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkHonour{})); err != nil {
		return err
	}
	var honour model.GqaPluginXtkfkHonour
	if err = db.Where("id = ?", toEditHonour.Id).First(&honour).Error; err != nil {
		return err
	}
	//err = db.Updates(&toEditHonour).Error
	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEditHonour.GqaModel),
		map[string]interface{}{
			"title":      toEditHonour.Title,
			"attachment": toEditHonour.Attachment,
		})).Error
	return err
}

func AddHonour(toAddHonour model.GqaPluginXtkfkHonour, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkHonour{})); err != nil {
		return err
	}
	err = db.Create(&toAddHonour).Error
	return err
}

func DeleteHonourById(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkHonour{})); err != nil {
		return err
	}
	var honour model.GqaPluginXtkfkHonour
	if err = db.Where("id = ?", id).First(&honour).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&honour).Error
	return err
}

func QueryHonourById(id uint, username string) (err error, honourInfo model.GqaPluginXtkfkHonour) {
	var honour model.GqaPluginXtkfkHonour
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkHonour{})); err != nil {
		return err, honour
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&honour, "id = ?", id).Error
	return err, honour
}
