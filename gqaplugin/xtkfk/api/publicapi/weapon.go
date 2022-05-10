package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetWeaponLanguageList(c *gin.Context) {
	if err, language := publicservice.GetWeaponLanguageList(); err != nil {
		gqaGlobal.GqaLogger.Error("获取武器库语言列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取武器库语言列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": language}, "获取武器库语言列表成功！", c)
	}
}

func GetWeaponNodeList(c *gin.Context) {
	if err, node := publicservice.GetWeaponNodeList(); err != nil {
		gqaGlobal.GqaLogger.Error("获取武器库节点列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取武器库节点列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": node}, "获取武器库节点列表成功！", c)
	}
}
