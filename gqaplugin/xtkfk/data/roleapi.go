package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginXtkfkSysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysRoleApi{}).Where("api_group = ?", "plugin-xtkfk").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_api 表中xtkfk插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_role_api 表中xtkfk插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_role_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_role_api 表成功！")
		return nil
	})
}

var sysRoleApiData = []gqaModel.SysRoleApi{
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-news-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-news"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-news"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-news-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-news-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-project-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-project"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-project-detail"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-project"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-project-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-project-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-honour-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-honour"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-honour"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-honour-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-honour-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-document-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-document"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-document"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-document-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-document-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-download-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-download"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-download"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-download-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-download-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-abbreviate-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-abbreviate"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-abbreviate"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-abbreviate-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-abbreviate-by-id"},
}
