package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHonourList(c *gin.Context) {
	var getHonourList model.RequestGetHonourList
	if err := gqaModel.RequestShouldBindJSON(c, &getHonourList); err != nil {
		return
	}
	if err, honour, total := privateservice.GetHonourList(getHonourList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取荣誉认证列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取荣誉认证列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  honour,
			Page:     getHonourList.Page,
			PageSize: getHonourList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditHonour(c *gin.Context) {
	var toEditHonour model.GqaPluginXkHonour
	if err := gqaModel.RequestShouldBindJSON(c, &toEditHonour); err != nil {
		return
	}
	toEditHonour.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditHonour(toEditHonour, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑荣誉认证失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑荣誉认证失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑荣誉认证成功！", c)
	}
}

func AddHonour(c *gin.Context) {
	var toAddHonour model.RequestAddHonour
	if err := gqaModel.RequestShouldBindJSON(c, &toAddHonour); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddHonour.Status,
			Sort:      toAddHonour.Sort,
			Memo:      toAddHonour.Memo,
		},
	}
	addHonour := &model.GqaPluginXkHonour{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Title:                             toAddHonour.Title,
		Attachment:                        toAddHonour.Attachment,
	}
	if err := privateservice.AddHonour(*addHonour, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加荣誉认证失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加荣誉认证失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加荣誉认证成功！", c)
	}
}

func DeleteHonourById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteHonourById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除荣誉认证失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除荣誉认证失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除荣誉认证成功！", c)
	}
}

func QueryHonourById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryHonourById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找荣誉认证失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找荣誉认证失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找荣誉认证成功！", c)
	}
}
