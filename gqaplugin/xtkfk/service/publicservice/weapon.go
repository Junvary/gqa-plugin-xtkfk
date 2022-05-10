package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetWeaponLanguageList() (err error, language []interface{}) {
	var languageList []gqaModel.SysDict
	if err = gqaGlobal.GqaDb.Where("parent_code = ?", "codeLanguage").Find(&languageList).Error; err != nil {
		return err, nil
	}
	type record struct {
		Value int64  `json:"value"`
		Name  string `json:"name"`
	}
	for _, v := range languageList {
		var total int64
		projectDb := gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkProject{})
		projectDb.Where("language like ?", "%"+v.DictCode+"%").Count(&total)
		if total != 0 {
			language = append(language, record{
				Value: total,
				Name:  v.DictCode,
			})
		}
	}
	return nil, language
}

func GetWeaponNodeList() (err error, node []interface{}) {
	var nodeList []gqaModel.SysDict
	if err = gqaGlobal.GqaDb.Where("parent_code = ?", "projectNode").Find(&nodeList).Error; err != nil {
		return err, nil
	}
	type record struct {
		Value int64  `json:"value"`
		Name  string `json:"name"`
	}
	for _, v := range nodeList {
		var total int64
		projectDb := gqaGlobal.GqaDb.Model(&model.GqaPluginXtkfkProject{})
		projectDb.Where("node like ?", "%"+v.DictCode+"%").Count(&total)
		if total != 0 {
			node = append(node, record{
				Value: total,
				Name:  v.DictLabel,
			})
		}
	}
	return nil, node
}
