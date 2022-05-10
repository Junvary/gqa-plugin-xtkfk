package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetNewsList(c *gin.Context) {
	var getNewsList model.RequestGetNewsList
	if err := gqaModel.RequestShouldBindJSON(c, &getNewsList); err != nil {
		return
	}
	if err, news, total := publicservice.GetNewsList(getNewsList); err != nil {
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
