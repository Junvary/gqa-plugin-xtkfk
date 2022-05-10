package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginXkSysDict = new(sysDict)

type sysDict struct{}

func (s *sysDict) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysDict{}).Where("parent_code = ?", "projectNode").Or("dict_code = ?", "projectNode").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_dict 表中xk插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_dict 表中xk插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_dict 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> xk插件初始数据进入 sys_dict 表成功！")
		return nil
	})
}

var sysDictData = []gqaModel.SysDict{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 801, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编程语言/框架",
	}}, DictCode: "codeLanguage", DictLabel: "编程语言/框架"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Java",
	}}, DictCode: "Java", DictLabel: "Java", ParentCode: "codeLanguage"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Python",
	}}, DictCode: "Python", DictLabel: "Python", ParentCode: "codeLanguage"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "C#",
	}}, DictCode: "C#", DictLabel: "C#", ParentCode: "codeLanguage"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 4, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Go",
	}}, DictCode: "Go", DictLabel: "Go", ParentCode: "codeLanguage"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 5, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Vue",
	}}, DictCode: "Vue", DictLabel: "Vue", ParentCode: "codeLanguage"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 6, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "React",
	}}, DictCode: "React", DictLabel: "React", ParentCode: "codeLanguage"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 7, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Bootstrap",
	}}, DictCode: "Bootstrap", DictLabel: "Bootstrap", ParentCode: "codeLanguage"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 802, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "项目节点",
	}}, DictCode: "projectNode", DictLabel: "项目节点"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "需求申请文档、需求规格说明书、需求确认单等",
	}}, DictCode: "p1", DictLabel: "需求申请", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "技术方案、系统设计、数据库设计等",
	}}, DictCode: "p2", DictLabel: "系统设计", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "进度计划、进度汇报等",
	}}, DictCode: "p3", DictLabel: "系统开发", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 4, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "测试计划、问题反馈、测试用例、测试记录等",
	}}, DictCode: "p4", DictLabel: "系统测试", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 5, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上线计划、操作手册、环境配置、代码位置等",
	}}, DictCode: "p5", DictLabel: "交付上线", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 6, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "运维手册、开关机手册、日常点检手册等",
	}}, DictCode: "p6", DictLabel: "产品运维", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 8, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "下线原因、替代产品等",
	}}, DictCode: "p7", DictLabel: "产品下线", ParentCode: "projectNode"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 21, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "会议纪要、培训计划、培训记录、会议签到表等",
	}}, DictCode: "m1", DictLabel: "会议材料", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 22, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "立项申请表、立项会确认表、功能确认表、验收表、代码统计等",
	}}, DictCode: "m2", DictLabel: "立项材料", ParentCode: "projectNode"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 23, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "各类软著申请材料",
	}}, DictCode: "m3", DictLabel: "软著材料", ParentCode: "projectNode"},
}
