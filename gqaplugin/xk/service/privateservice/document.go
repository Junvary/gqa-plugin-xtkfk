package privateservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"gorm.io/gorm"
)

func GetDocumentList(getDocumentList model.RequestGetDocumentList, username string) (err error, document []model.GqaPluginXkDocument, total int64) {
	pageSize := getDocumentList.PageSize
	offset := getDocumentList.PageSize * (getDocumentList.Page - 1)
	var documentList []model.GqaPluginXkDocument
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXkDocument{})); err != nil {
		return err, documentList, 0
	}
	//配置搜索
	if getDocumentList.Title != "" {
		db = db.Where("title like ?", "%"+getDocumentList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getDocumentList.SortBy, getDocumentList.Desc)).Preload("CreatedByUser").Find(&documentList).Error
	return err, documentList, total
}

func EditDocument(toEditDocument model.GqaPluginXkDocument, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXkDocument{})); err != nil {
		return err
	}
	var document model.GqaPluginXkDocument
	if err = db.Where("id = ?", toEditDocument.Id).First(&document).Error; err != nil {
		return err
	}
	//err = db.Updates(&toEditDocument).Error
	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEditDocument.GqaModel), map[string]interface{}{
		"title":      toEditDocument.Title,
		"content":    toEditDocument.Content,
		"attachment": toEditDocument.Attachment,
	})).Error
	return err
}

func AddDocument(toAddDocument model.GqaPluginXkDocument, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXkDocument{})); err != nil {
		return err
	}
	err = db.Create(&toAddDocument).Error
	return err
}

func DeleteDocumentById(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXkDocument{})); err != nil {
		return err
	}
	var document model.GqaPluginXkDocument
	if err = db.Where("id = ?", id).First(&document).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&document).Error
	return err
}

func QueryDocumentById(id uint, username string) (err error, documentInfo model.GqaPluginXkDocument) {
	var document model.GqaPluginXkDocument
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginXkDocument{})); err != nil {
		return err, document
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&document, "id = ?", id).Error
	return err, document
}
