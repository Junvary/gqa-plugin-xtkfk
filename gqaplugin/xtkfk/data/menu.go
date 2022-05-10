package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginXkSysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPluginXk").Or("name = ?", "GqaPluginXk").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中xk插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_menu 表中xk插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> xk插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 801, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是系统开发科插件",
	}}, Name: "GqaPluginXk", Title: "系统开发科", Icon: "home", Path: "", Component: ""},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是最新要闻",
	}}, Name: "plugin-xk-news", Title: "最新要闻", Icon: "newspaper", Path: "/plugin-xk/xk/news", Component: "plugins/xk/News/index", ParentCode: "GqaPluginXk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是项目管理",
	}}, Name: "plugin-xk-project", Title: "项目管理", Icon: "beenhere", Path: "/plugin-xk/xk/project", Component: "plugins/xk/Project/index", ParentCode: "GqaPluginXk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是荣誉认证",
	}}, Name: "plugin-xk-honour", Title: "荣誉认证", Icon: "stars", Path: "/plugin-xk/xk/honour", Component: "plugins/xk/Honour/index", ParentCode: "GqaPluginXk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 4, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是文档管理",
	}}, Name: "plugin-xk-document", Title: "文档管理", Icon: "edit_note", Path: "/plugin-xk/xk/document", Component: "plugins/xk/Document/index", ParentCode: "GqaPluginXk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 5, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是资源管理",
	}}, Name: "plugin-xk-download", Title: "资源管理", Icon: "download", Path: "/plugin-xk/xk/download", Component: "plugins/xk/Download/index", ParentCode: "GqaPluginXk"},
}
