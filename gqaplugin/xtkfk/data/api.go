package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginXtkfkSysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-xtkfk").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中xtkfk插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_api 表中xtkfk插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> xtkfk插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 801, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：获取news-list",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-news-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 802, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：编辑news信息",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-news"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 803, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：新增news",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-news"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 804, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：删除news",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-news-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 805, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：根据ID查找news",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-news-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 806, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：获取project-list",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-project-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 807, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：编辑project信息",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-project"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 808, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：编辑project详情",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-project-detail"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 809, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：新增project",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-project"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 810, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：删除project",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-project-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 811, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：根据ID查找project",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-project-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 812, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：获取honour-list",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-honour-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 813, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：编辑honour信息",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-honour"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 814, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：新增honour",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-honour"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 815, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：删除honour",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-honour-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 816, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：根据ID查找honour",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-honour-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 817, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：获取document-list",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-document-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 818, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：编辑document信息",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-document"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 819, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：新增document",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-document"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 820, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：删除document",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-document-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 821, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：根据ID查找document",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-document-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 822, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：获取download-list",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/get-download-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 823, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：编辑download信息",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/edit-download"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 824, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：新增download",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/add-download"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 825, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：删除download",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/delete-download-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 826, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件xtkfk：根据ID查找download",
	}}, ApiGroup: "plugin-xtkfk", ApiMethod: "POST", ApiPath: "/plugin-xtkfk/query-download-by-id"},
}
