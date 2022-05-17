package privateservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"gorm.io/gorm"
)

func GetAbbreviateList(getAbbreviateList model.RequestGetAbbreviateList, username string) (err error, document []model.GqaPluginXtkfkAbbreviate, total int64) {
	pageSize := getAbbreviateList.PageSize
	offset := getAbbreviateList.PageSize * (getAbbreviateList.Page - 1)
	var documentList []model.GqaPluginXtkfkAbbreviate
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkAbbreviate{})); err != nil {
		return err, documentList, 0
	}
	//配置搜索
	if getAbbreviateList.Top != "" {
		db = db.Where("top like ?", "%"+getAbbreviateList.Top+"%")
	}
	if getAbbreviateList.Title != "" {
		db = db.Where("title like ?", "%"+getAbbreviateList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getAbbreviateList.SortBy, getAbbreviateList.Desc)).Preload("CreatedByUser").Find(&documentList).Error
	return err, documentList, total
}

func EditAbbreviate(toEditAbbreviate model.GqaPluginXtkfkAbbreviate, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkAbbreviate{})); err != nil {
		return err
	}
	var document model.GqaPluginXtkfkAbbreviate
	if err = db.Where("id = ?", toEditAbbreviate.Id).First(&document).Error; err != nil {
		return err
	}
	//err = db.Updates(&toEditAbbreviate).Error
	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEditAbbreviate.GqaModel), map[string]interface{}{
		"top":     toEditAbbreviate.Top,
		"from":    toEditAbbreviate.From,
		"title":   toEditAbbreviate.Title,
		"content": toEditAbbreviate.Content,
	})).Error
	return err
}

func AddAbbreviate(toAddAbbreviate model.GqaPluginXtkfkAbbreviate, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkAbbreviate{})); err != nil {
		return err
	}
	err = db.Create(&toAddAbbreviate).Error
	return err
}

func DeleteAbbreviateById(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkAbbreviate{})); err != nil {
		return err
	}
	var document model.GqaPluginXtkfkAbbreviate
	if err = db.Where("id = ?", id).First(&document).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&document).Error
	return err
}

func QueryAbbreviateById(id uint, username string) (err error, documentInfo model.GqaPluginXtkfkAbbreviate) {
	var document model.GqaPluginXtkfkAbbreviate
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkAbbreviate{})); err != nil {
		return err, document
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&document, "id = ?", id).Error
	return err, document
}
