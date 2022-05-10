package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginXkSysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysRoleApi{}).Where("api_group = ?", "plugin-xk").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_api 表中xk插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_role_api 表中xk插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_role_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> xk插件初始数据进入 sys_role_api 表成功！")
		return nil
	})
}

var sysRoleApiData = []gqaModel.SysRoleApi{
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/get-news-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/edit-news"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/add-news"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/delete-news-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/query-news-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/get-project-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/edit-project"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/edit-project-detail"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/add-project"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/delete-project-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/query-project-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/get-honour-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/edit-honour"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/add-honour"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/delete-honour-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/query-honour-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/get-document-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/edit-document"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/add-document"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/delete-document-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/query-document-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/get-download-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/edit-download"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/add-download"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/delete-download-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-xk", ApiMethod: "POST", ApiPath: "/plugin-xk/query-download-by-id"},
}
