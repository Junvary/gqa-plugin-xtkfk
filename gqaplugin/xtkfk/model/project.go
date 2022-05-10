package model

import (
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/google/uuid"
)

type GqaPluginXtkfkProject struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	ProjectId      uuid.UUID                     `json:"project_id" gorm:"comment:项目ID;index;not null"`
	ProjectName    string                        `json:"project_name" gorm:"comment:项目名称;not null;index"`
	Demand         string                        `json:"demand" gorm:"comment:需求单位;"`
	LeaderUsername string                        `json:"leader_username" gorm:"comment:牵头人username;"`
	Leader         gqaModel.SysUser              `json:"leader" gorm:"comment:牵头人;foreignKey:LeaderUsername;references:Username"`
	Player         string                        `json:"player" gorm:"comment:参与人;"`
	Language       string                        `json:"language" gorm:"comment:项目语言"`
	Node           string                        `json:"node" gorm:"comment:项目节点"`
	ProjectDetail  []GqaPluginXtkfkProjectDetail `json:"project_detail" gorm:"foreignKey:ProjectId;references:ProjectId"`
}

type RequestAddProject struct {
	gqaModel.RequestAdd
	ProjectName    string                        `json:"project_name"`
	Demand         string                        `json:"demand"`
	LeaderUsername string                        `json:"leader_username"`
	Player         string                        `json:"player"`
	Language       string                        `json:"language"`
	Node           string                        `json:"node"`
	ProjectDetail  []GqaPluginXtkfkProjectDetail `json:"project_detail"`
}

type RequestGetProjectList struct {
	gqaModel.RequestPageAndSort
	ProjectName string `json:"project_name"`
}
