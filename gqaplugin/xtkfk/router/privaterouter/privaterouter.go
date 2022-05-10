package privaterouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/api/privateapi"
	gqaMiddleware "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	xtkfkRouter := privateGroup.Use(gqaMiddleware.LogOperationHandler())
	{
		//获取news列表
		xtkfkRouter.POST("get-news-list", privateapi.GetNewsList)
		//编辑news信息
		xtkfkRouter.POST("edit-news", privateapi.EditNews)
		//新增news
		xtkfkRouter.POST("add-news", privateapi.AddNews)
		//删除news
		xtkfkRouter.POST("delete-news-by-id", privateapi.DeleteNewsById)
		//根据ID查找news
		xtkfkRouter.POST("query-news-by-id", privateapi.QueryNewsById)
	}
	{
		//获取project列表
		xtkfkRouter.POST("get-project-list", privateapi.GetProjectList)
		//编辑project信息
		xtkfkRouter.POST("edit-project", privateapi.EditProject)
		//编辑project详情
		xtkfkRouter.POST("edit-project-detail", privateapi.EditProjectDetail)
		//新增project
		xtkfkRouter.POST("add-project", privateapi.AddProject)
		//删除project
		xtkfkRouter.POST("delete-project-by-id", privateapi.DeleteProjectById)
		//根据ID查找project
		xtkfkRouter.POST("query-project-by-id", privateapi.QueryProjectById)
	}
	{
		//获取honour列表
		xtkfkRouter.POST("get-honour-list", privateapi.GetHonourList)
		//编辑honour信息
		xtkfkRouter.POST("edit-honour", privateapi.EditHonour)
		//新增honour
		xtkfkRouter.POST("add-honour", privateapi.AddHonour)
		//删除honour
		xtkfkRouter.POST("delete-honour-by-id", privateapi.DeleteHonourById)
		//根据ID查找honour
		xtkfkRouter.POST("query-honour-by-id", privateapi.QueryHonourById)
	}
	{
		//获取document列表
		xtkfkRouter.POST("get-document-list", privateapi.GetDocumentList)
		//编辑document信息
		xtkfkRouter.POST("edit-document", privateapi.EditDocument)
		//新增document
		xtkfkRouter.POST("add-document", privateapi.AddDocument)
		//删除document
		xtkfkRouter.POST("delete-document-by-id", privateapi.DeleteDocumentById)
		//根据ID查找document
		xtkfkRouter.POST("query-document-by-id", privateapi.QueryDocumentById)
	}
	{
		//获取download列表
		xtkfkRouter.POST("get-download-list", privateapi.GetDownloadList)
		//编辑download信息
		xtkfkRouter.POST("edit-download", privateapi.EditDownload)
		//新增download
		xtkfkRouter.POST("add-download", privateapi.AddDownload)
		//删除download
		xtkfkRouter.POST("delete-download-by-id", privateapi.DeleteDownload)
		//根据ID查找download
		xtkfkRouter.POST("query-download-by-id", privateapi.QueryDownloadById)
	}
}
