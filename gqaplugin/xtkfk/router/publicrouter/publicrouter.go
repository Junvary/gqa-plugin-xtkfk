package publicrouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/api/publicapi"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		//获取news列表
		publicGroup.POST("get-news-list", publicapi.GetNewsList)
		//获取project列表
		publicGroup.POST("get-project-list", publicapi.GetProjectList)
		//获取武器库语言饼状图
		publicGroup.POST("get-weapon-language-list", publicapi.GetWeaponLanguageList)
		//获取武器库节点饼状图
		publicGroup.POST("get-weapon-node-list", publicapi.GetWeaponNodeList)
		//获取honour列表
		publicGroup.POST("get-honour-list", publicapi.GetHonourList)
		//获取document列表
		publicGroup.POST("get-document-list", publicapi.GetDocumentList)
		//获取download列表
		publicGroup.POST("get-download-list", publicapi.GetDownloadList)
		//获取Abbreviate列表
		publicGroup.POST("get-abbreviate-list", publicapi.GetAbbreviateList)
	}
}
