package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDocumentList(c *gin.Context) {
	var getDocumentList model.RequestGetDocumentList
	if err := gqaModel.RequestShouldBindJSON(c, &getDocumentList); err != nil {
		return
	}
	if err, document, total := privateservice.GetDocumentList(getDocumentList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取文档列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取文档列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  document,
			Page:     getDocumentList.Page,
			PageSize: getDocumentList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditDocument(c *gin.Context) {
	var toEditDocument model.GqaPluginXtkfkDocument
	if err := gqaModel.RequestShouldBindJSON(c, &toEditDocument); err != nil {
		return
	}
	toEditDocument.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditDocument(toEditDocument, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑文档失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑文档失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑文档成功！", c)
	}
}

func AddDocument(c *gin.Context) {
	var toAddDocument model.RequestAddDocument
	if err := gqaModel.RequestShouldBindJSON(c, &toAddDocument); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddDocument.Status,
			Sort:      toAddDocument.Sort,
			Memo:      toAddDocument.Memo,
		},
	}
	addDocument := &model.GqaPluginXtkfkDocument{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Title:                             toAddDocument.Title,
		Content:                           toAddDocument.Content,
		Attachment:                        toAddDocument.Attachment,
	}
	if err := privateservice.AddDocument(*addDocument, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加文档失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加文档失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加文档成功！", c)
	}
}

func DeleteDocumentById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteDocumentById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除文档失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除文档失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除文档成功！", c)
	}
}

func QueryDocumentById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryDocumentById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找文档失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找文档失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找文档成功！", c)
	}
}
