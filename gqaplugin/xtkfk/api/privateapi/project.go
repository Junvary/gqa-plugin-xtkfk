package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetProjectList(c *gin.Context) {
	var getProjectList model.RequestGetProjectList
	if err := gqaModel.RequestShouldBindJSON(c, &getProjectList); err != nil {
		return
	}
	if err, project, total := privateservice.GetProjectList(getProjectList, gqaUtils.GetUsername(c)); err != nil {
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

func EditProject(c *gin.Context) {
	var toEditProject model.GqaPluginXtkfkProject
	if err := gqaModel.RequestShouldBindJSON(c, &toEditProject); err != nil {
		return
	}
	toEditProject.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditProject(toEditProject, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑项目失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑项目失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑项目成功！", c)
	}
}

func EditProjectDetail(c *gin.Context) {
	var toEditProjectDetail model.RequestEditProjectDetail
	if err := gqaModel.RequestShouldBindJSON(c, &toEditProjectDetail); err != nil {
		return
	}
	if err := privateservice.EditProjectDetail(toEditProjectDetail); err != nil {
		gqaGlobal.GqaLogger.Error("编辑项目失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑项目失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑项目成功！", c)
	}
}

func AddProject(c *gin.Context) {
	var toAddProject model.RequestAddProject
	if err := gqaModel.RequestShouldBindJSON(c, &toAddProject); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddProject.Status,
			Sort:      toAddProject.Sort,
			Memo:      toAddProject.Memo,
		},
	}
	addProject := &model.GqaPluginXtkfkProject{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ProjectId:                         uuid.New(),
		ProjectName:                       toAddProject.ProjectName,
		Demand:                            toAddProject.Demand,
		LeaderUsername:                    toAddProject.LeaderUsername,
		Player:                            toAddProject.Player,
		Language:                          toAddProject.Language,
		Node:                              toAddProject.Node,
	}
	if err := privateservice.AddProject(*addProject, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加项目失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加项目失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加项目成功！", c)
	}
}

func DeleteProjectById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteProjectById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除项目失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除项目失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除项目成功！", c)
	}
}

func QueryProjectById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryProjectById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找项目失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找项目失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找项目成功！", c)
	}
}
