package model

import (
	"GinServer/server/api"

	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

// 级别查询导出，无需导入数据库
type LevelsRawData struct {
	gorm.Model
	// 序号
	IDNumber string `mapstructure:"序号" json:"id_number" gorm:"column:id_number"`
	// 姓名
	Name string `mapstructure:"姓名" json:"name" gorm:"column:name"`
	// 工号
	EmployeeID string `mapstructure:"工号" json:"employee_id" gorm:"column:employee_id"`
	// 性别
	Gender string `mapstructure:"性别" json:"gender" gorm:"column:gender"`
	// 出生日期
	BirthDate string `mapstructure:"出生日期" json:"birth_date" gorm:"column:birth_date"`
	// 标准
	Standard string `mapstructure:"标准" json:"standard" gorm:"column:standard"`
	// 技术机型
	TechModel string `mapstructure:"技术机型" json:"tech_model" gorm:"column:tech_model"`
	// 级别
	Level string `mapstructure:"级别" json:"level" gorm:"column:level"`
	// 符号
	Symbol string `mapstructure:"符号" json:"symbol" gorm:"column:symbol"`
	// 中文名称
	ChineseName string `mapstructure:"中文名称" json:"chinese_name" gorm:"column:chinese_name"`
	// 分类
	Category string `mapstructure:"分类" json:"category" gorm:"column:category"`
	// 生效日期
	EffectiveDate string `mapstructure:"生效日期" json:"effective_date" gorm:"column:effective_date"`
	// 有效日期
	ExpiryDate string `mapstructure:"有效日期" json:"expiry_date" gorm:"column:expiry_date"`
}

func (v LevelsRawData) TimeStamp() float32 {
	effectiveDate := carbon.Parse(v.EffectiveDate)
	return float32(effectiveDate.Timestamp())
}

func (v LevelsRawData) RelevantLevel() api.SpecificLevel {
	for _, level := range AllSpecificLevels {
		if level == api.SpecificLevel(v.Level) {
			return level
		}
	}
	if v.Symbol == "T" {
		return "Ta"
	}
	return api.SpecificLevel(v.Level)
}
