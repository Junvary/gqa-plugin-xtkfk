package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDocumentList(c *gin.Context) {
	var getDocumentList model.RequestGetDocumentList
	if err := gqaModel.RequestShouldBindJSON(c, &getDocumentList); err != nil {
		return
	}
	if err, news, total := publicservice.GetDocumentList(getDocumentList); err != nil {
		gqaGlobal.GqaLogger.Error("获取文档列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取文档列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  news,
			Page:     getDocumentList.Page,
			PageSize: getDocumentList.PageSize,
			Total:    total,
		}, c)
	}
}
