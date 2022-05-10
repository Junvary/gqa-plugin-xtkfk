package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHonourList(c *gin.Context) {
	var getHonourList model.RequestGetHonourList
	if err := gqaModel.RequestShouldBindJSON(c, &getHonourList); err != nil {
		return
	}
	if err, honour, total := publicservice.GetHonourList(getHonourList); err != nil {
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
