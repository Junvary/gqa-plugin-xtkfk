package privaterouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/api/privateapi"
	gqaMiddleware "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	xkRouter := privateGroup.Use(gqaMiddleware.LogOperationHandler())
	{
		//获取news列表
		xkRouter.POST("get-news-list", privateapi.GetNewsList)
		//编辑news信息
		xkRouter.POST("edit-news", privateapi.EditNews)
		//新增news
		xkRouter.POST("add-news", privateapi.AddNews)
		//删除news
		xkRouter.POST("delete-news-by-id", privateapi.DeleteNewsById)
		//根据ID查找news
		xkRouter.POST("query-news-by-id", privateapi.QueryNewsById)
	}
	{
		//获取project列表
		xkRouter.POST("get-project-list", privateapi.GetProjectList)
		//编辑project信息
		xkRouter.POST("edit-project", privateapi.EditProject)
		//编辑project详情
		xkRouter.POST("edit-project-detail", privateapi.EditProjectDetail)
		//新增project
		xkRouter.POST("add-project", privateapi.AddProject)
		//删除project
		xkRouter.POST("delete-project-by-id", privateapi.DeleteProjectById)
		//根据ID查找project
		xkRouter.POST("query-project-by-id", privateapi.QueryProjectById)
	}
	{
		//获取honour列表
		xkRouter.POST("get-honour-list", privateapi.GetHonourList)
		//编辑honour信息
		xkRouter.POST("edit-honour", privateapi.EditHonour)
		//新增honour
		xkRouter.POST("add-honour", privateapi.AddHonour)
		//删除honour
		xkRouter.POST("delete-honour-by-id", privateapi.DeleteHonourById)
		//根据ID查找honour
		xkRouter.POST("query-honour-by-id", privateapi.QueryHonourById)
	}
	{
		//获取document列表
		xkRouter.POST("get-document-list", privateapi.GetDocumentList)
		//编辑document信息
		xkRouter.POST("edit-document", privateapi.EditDocument)
		//新增document
		xkRouter.POST("add-document", privateapi.AddDocument)
		//删除document
		xkRouter.POST("delete-document-by-id", privateapi.DeleteDocumentById)
		//根据ID查找document
		xkRouter.POST("query-document-by-id", privateapi.QueryDocumentById)
	}
	{
		//获取download列表
		xkRouter.POST("get-download-list", privateapi.GetDownloadList)
		//编辑download信息
		xkRouter.POST("edit-download", privateapi.EditDownload)
		//新增download
		xkRouter.POST("add-download", privateapi.AddDownload)
		//删除download
		xkRouter.POST("delete-download-by-id", privateapi.DeleteDownload)
		//根据ID查找download
		xkRouter.POST("query-download-by-id", privateapi.QueryDownloadById)
	}
}
