package model

import (
	"github.com/google/uuid"
)

type GqaPluginXtkfkProjectDetail struct {
	ProjectId   uuid.UUID `json:"project_id" gorm:"comment:项目id;index;not null"`
	ProjectNode string    `json:"project_node" gorm:"comment:项目节点"`
	Content     string    `json:"content" gorm:"comment:描述;type:text"`
	Attachment  string    `json:"attachment" gorm:"comment:附件;type:text;"`
	StartTime   string    `json:"start_time" gorm:"comment:开始时间"`
	EndTime     string    `json:"end_time" gorm:"comment:结束时间"`
}

type RequestEditProjectDetail struct {
	ProjectId     uuid.UUID                     `json:"project_id"`
	ProjectDetail []GqaPluginXtkfkProjectDetail `json:"project_detail"`
}
