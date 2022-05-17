package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginXtkfkSysRoleMenu = new(sysRoleMenu)

type sysRoleMenu struct{}

func (s *sysRoleMenu) LoadData() error {
	return gqaGlobal.GqaDb.Table("sys_role_menu").Transaction(func(tx *gorm.DB) error {
		var count int64
		var menuName []string
		for _, v := range sysRoleMenuData {
			menuName = append(menuName, v.SysMenuName)
		}
		tx.Model(&gqaModel.SysRoleMenu{}).Where("sys_menu_name in ?", menuName).Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_menu 表中xtkfk插件数已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_role_menu 表中xtkfk插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Save(&sysRoleMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_role_menu 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_role_menu 表成功！")
		return nil
	})
}

var sysRoleMenuData = []gqaModel.SysRoleMenu{
	{"super-admin", "GqaPluginXtkfk"},
	{"super-admin", "plugin-xtkfk-news"},
	{"super-admin", "plugin-xtkfk-project"},
	{"super-admin", "plugin-xtkfk-honour"},
	{"super-admin", "plugin-xtkfk-document"},
	{"super-admin", "plugin-xtkfk-download"},
	{"super-admin", "plugin-xtkfk-abbreviate"},
}
