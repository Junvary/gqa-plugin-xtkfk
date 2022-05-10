package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetProjectList(getProjectList model.RequestGetProjectList) (err error, project []model.GqaPluginXtkfkProject, total int64) {
	pageSize := getProjectList.PageSize
	offset := getProjectList.PageSize * (getProjectList.Page - 1)
	db := gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkProject{})
	var projectList []model.GqaPluginXtkfkProject
	//配置搜索
	if getProjectList.ProjectName != "" {
		db = db.Where("project_name like ?", "%"+getProjectList.ProjectName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getProjectList.SortBy, getProjectList.Desc)).
		Preload("Leader").Find(&projectList).Error
	return err, projectList, total
}
