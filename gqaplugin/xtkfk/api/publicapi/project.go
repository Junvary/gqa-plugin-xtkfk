package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetProjectList(c *gin.Context) {
	var getProjectList model.RequestGetProjectList
	if err := gqaModel.RequestShouldBindJSON(c, &getProjectList); err != nil {
		return
	}
	if err, project, total := publicservice.GetProjectList(getProjectList); err != nil {
		gqaGlobal.GqaLogger.Error("获取项目列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取项目列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  project,
			Page:     getProjectList.Page,
			PageSize: getProjectList.PageSize,
			Total:    total,
		}, c)
	}
}
