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

func GetNewsList(c *gin.Context) {
	var getNewsList model.RequestGetNewsList
	if err := gqaModel.RequestShouldBindJSON(c, &getNewsList); err != nil {
		return
	}
	if err, news, total := privateservice.GetNewsList(getNewsList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取最新要闻列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取最新要闻列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  news,
			Page:     getNewsList.Page,
			PageSize: getNewsList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditNews(c *gin.Context) {
	var toEditNews model.GqaPluginXtkfkNews
	if err := gqaModel.RequestShouldBindJSON(c, &toEditNews); err != nil {
		return
	}
	toEditNews.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditNews(toEditNews, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑最新要闻失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑最新要闻失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑最新要闻成功！", c)
	}
}

func AddNews(c *gin.Context) {
	var toAddNews model.RequestAddNews
	if err := gqaModel.RequestShouldBindJSON(c, &toAddNews); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddNews.Status,
			Sort:      toAddNews.Sort,
			Memo:      toAddNews.Memo,
		},
	}
	addNews := &model.GqaPluginXtkfkNews{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Title:                             toAddNews.Title,
		Content:                           toAddNews.Content,
		Attachment:                        toAddNews.Attachment,
	}
	if err := privateservice.AddNews(*addNews, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加最新要闻失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加最新要闻失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加最新要闻成功！", c)
	}
}

func DeleteNewsById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteNewsById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除最新要闻失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除最新要闻失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除最新要闻成功！", c)
	}
}

func QueryNewsById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryNewsById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找最新要闻失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找最新要闻失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找最新要闻成功！", c)
	}
}
