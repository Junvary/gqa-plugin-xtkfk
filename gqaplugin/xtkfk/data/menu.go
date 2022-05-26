package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginXtkfkSysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPluginXtkfk").Or("name = ?", "GqaPluginXtkfk").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中xtkfk插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_menu 表中xtkfk插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 801, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是系统开发科插件",
	}}, IsPlugin: "yes", Name: "GqaPluginXtkfk", Title: "系统开发科", Icon: "home", Path: "", Component: ""},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是最新要闻",
	}}, IsPlugin: "yes", Name: "plugin-xtkfk-news", Title: "最新要闻", Icon: "newspaper", Path: "/plugin-xtkfk/xtkfk/news", Component: "plugins/xtkfk/News/index", ParentCode: "GqaPluginXtkfk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是项目管理",
	}}, IsPlugin: "yes", Name: "plugin-xtkfk-project", Title: "项目管理", Icon: "beenhere", Path: "/plugin-xtkfk/xtkfk/project", Component: "plugins/xtkfk/Project/index", ParentCode: "GqaPluginXtkfk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是荣誉认证",
	}}, IsPlugin: "yes", Name: "plugin-xtkfk-honour", Title: "荣誉认证", Icon: "stars", Path: "/plugin-xtkfk/xtkfk/honour", Component: "plugins/xtkfk/Honour/index", ParentCode: "GqaPluginXtkfk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 4, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是文档管理",
	}}, IsPlugin: "yes", Name: "plugin-xtkfk-document", Title: "文档管理", Icon: "edit_note", Path: "/plugin-xtkfk/xtkfk/document", Component: "plugins/xtkfk/Document/index", ParentCode: "GqaPluginXtkfk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 5, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是资源管理",
	}}, IsPlugin: "yes", Name: "plugin-xtkfk-download", Title: "资源管理", Icon: "download", Path: "/plugin-xtkfk/xtkfk/download", Component: "plugins/xtkfk/Download/index", ParentCode: "GqaPluginXtkfk"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 6, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是简称管理",
	}}, IsPlugin: "yes", Name: "plugin-xtkfk-abbreviate", Title: "简称管理", Icon: "download", Path: "/plugin-xtkfk/xtkfk/abbreviate", Component: "plugins/xtkfk/Abbreviate/index", ParentCode: "GqaPluginXtkfk"},
}
