package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDownloadList(c *gin.Context) {
	var getDownloadList model.RequestGetDownloadList
	if err := gqaModel.RequestShouldBindJSON(c, &getDownloadList); err != nil {
		return
	}
	if err, news, total := publicservice.GetDownloadList(getDownloadList); err != nil {
		gqaGlobal.GqaLogger.Error("获取资源列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取资源列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  news,
			Page:     getDownloadList.Page,
			PageSize: getDownloadList.PageSize,
			Total:    total,
		}, c)
	}
}
