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

func GetAbbreviateList(c *gin.Context) {
	var getAbbreviateList model.RequestGetAbbreviateList
	if err := gqaModel.RequestShouldBindJSON(c, &getAbbreviateList); err != nil {
		return
	}
	if err, document, total := privateservice.GetAbbreviateList(getAbbreviateList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取简称列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取简称列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  document,
			Page:     getAbbreviateList.Page,
			PageSize: getAbbreviateList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditAbbreviate(c *gin.Context) {
	var toEditAbbreviate model.GqaPluginXtkfkAbbreviate
	if err := gqaModel.RequestShouldBindJSON(c, &toEditAbbreviate); err != nil {
		return
	}
	toEditAbbreviate.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditAbbreviate(toEditAbbreviate, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑简称失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑简称失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑简称成功！", c)
	}
}

func AddAbbreviate(c *gin.Context) {
	var toAddAbbreviate model.RequestAddAbbreviate
	if err := gqaModel.RequestShouldBindJSON(c, &toAddAbbreviate); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddAbbreviate.Status,
			Sort:      toAddAbbreviate.Sort,
			Memo:      toAddAbbreviate.Memo,
		},
	}
	addAbbreviate := &model.GqaPluginXtkfkAbbreviate{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Top:                               toAddAbbreviate.Top,
		From:                              toAddAbbreviate.From,
		Title:                             toAddAbbreviate.Title,
		Content:                           toAddAbbreviate.Content,
	}
	if err := privateservice.AddAbbreviate(*addAbbreviate, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加简称失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加简称失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加简称成功！", c)
	}
}

func DeleteAbbreviateById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteAbbreviateById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除简称失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除简称失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除简称成功！", c)
	}
}

func QueryAbbreviateById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryAbbreviateById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找简称失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找简称失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找简称成功！", c)
	}
}
