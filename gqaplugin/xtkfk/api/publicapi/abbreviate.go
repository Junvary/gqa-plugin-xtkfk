package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAbbreviateList(c *gin.Context) {
	var getAbbreviateList model.RequestGetAbbreviateList
	if err := gqaModel.RequestShouldBindJSON(c, &getAbbreviateList); err != nil {
		return
	}
	if err, news, total := publicservice.GetAbbreviateList(getAbbreviateList); err != nil {
		gqaGlobal.GqaLogger.Error("获取缩写短语列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取缩写短语列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  news,
			Page:     getAbbreviateList.Page,
			PageSize: getAbbreviateList.PageSize,
			Total:    total,
		}, c)
	}
}
