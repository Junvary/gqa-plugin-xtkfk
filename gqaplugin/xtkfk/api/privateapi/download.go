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

func GetDownloadList(c *gin.Context) {
	var getDownloadList model.RequestGetDownloadList
	if err := gqaModel.RequestShouldBindJSON(c, &getDownloadList); err != nil {
		return
	}
	if err, download, total := privateservice.GetDownloadList(getDownloadList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取资源列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取资源列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  download,
			Page:     getDownloadList.Page,
			PageSize: getDownloadList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditDownload(c *gin.Context) {
	var toEditDownload model.GqaPluginXtkfkDownload
	if err := gqaModel.RequestShouldBindJSON(c, &toEditDownload); err != nil {
		return
	}
	toEditDownload.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditDownload(toEditDownload, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑资源失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑资源失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑资源成功！", c)
	}
}

func AddDownload(c *gin.Context) {
	var toAddDownload model.RequestAddDownload
	if err := gqaModel.RequestShouldBindJSON(c, &toAddDownload); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddDownload.Status,
			Sort:      toAddDownload.Sort,
			Memo:      toAddDownload.Memo,
		},
	}
	addDownload := &model.GqaPluginXtkfkDownload{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Title:                             toAddDownload.Title,
		Content:                           toAddDownload.Content,
		Attachment:                        toAddDownload.Attachment,
	}
	if err := privateservice.AddDownload(*addDownload, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加资源失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加资源失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加资源成功！", c)
	}
}

func DeleteDownloadById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteDownloadById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除资源失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除资源失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除资源成功！", c)
	}
}

func QueryDownloadById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryDownloadById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找资源失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找资源失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找资源成功！", c)
	}
}
