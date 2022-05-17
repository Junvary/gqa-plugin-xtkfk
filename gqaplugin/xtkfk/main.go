package xtkfk

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/router/privaterouter"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/router/publicrouter"
	"github.com/gin-gonic/gin"
)

var PluginXtkfk = new(xtkfk)

type xtkfk struct{}

func (*xtkfk) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-xtkfk"
}

func (*xtkfk) PluginName() string { //实现接口方法，插件名称
	return "系统开发科"
}

func (*xtkfk) PluginVersion() string { //实现接口方法，插件版本
	return "v2.0.0"
}

func (*xtkfk) PluginMemo() string { //实现接口方法，插件描述
	return "这是系统开发科插件。"
}

func (p *xtkfk) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
	publicrouter.InitPublicRouter(publicGroup)
}

func (p *xtkfk) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
	privaterouter.InitPrivateRouter(privateGroup)
}

func (p *xtkfk) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	var ModelList = []interface{}{
		model.GqaPluginXtkfkNews{},
		model.GqaPluginXtkfkProject{},
		model.GqaPluginXtkfkProjectDetail{},
		model.GqaPluginXtkfkHonour{},
		model.GqaPluginXtkfkDocument{},
		model.GqaPluginXtkfkDownload{},
		model.GqaPluginXtkfkAbbreviate{},
	}
	return ModelList
}

func (p *xtkfk) PluginData() []interface{ LoadData() (err error) } { //实现接口方法，初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginXtkfkSysApi,
		data.PluginXtkfkSysRoleApi,
		data.PluginXtkfkSysMenu,
		data.PluginXtkfkSysRoleMenu,
		data.PluginXtkfkSysDict,
	}
	return DataList
}
