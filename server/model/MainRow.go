package model

import (
	"GinServer/server/api"

	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

type MainRow struct {
	gorm.Model

	// 员工号 主键
	StaffID string `json:"staffID,omitempty"`

	// 员工姓名
	Name string `json:"name,omitempty"`

	// 员工姓名简拼
	NameShortcut string `json:"nameShortcut,omitempty"`

	// 员工姓名拼音
	NamePinYin string `json:"namePinYin,omitempty"`

	// 当前行政部门
	CurrentDepartment string `json:"currentDepartment,omitempty"`

	// 在三大队时的最后一个行政部门
	LastDepartment3dd string `json:"lastDepartment3dd,omitempty"`

	// 行政部门历史
	DepartmentHistory string `json:"departmentHistory,omitempty"`

	// 遗忘飞过的机型，用json存储
	AircraftExperiences string `json:"AircraftExperiences,omitempty"`

	// 最高的飞行标准
	HighestLevel api.SpecificLevel `json:"highestLevel,omitempty"`

	// F1升级日期时间戳
	F1 float32 `json:"F1,omitempty"`

	// F2升级日期时间戳
	F2 float32 `json:"F2,omitempty"`

	// F2b升级日期时间戳
	F2b float32 `json:"F2b,omitempty"`

	// F3升级日期时间戳
	F3 float32 `json:"F3,omitempty"`

	// F3b升级日期时间戳
	F3b float32 `json:"F3b,omitempty"`

	// F4升级日期时间戳
	F4 float32 `json:"F4,omitempty"`

	// F5升级日期时间戳
	F5 float32 `json:"F5,omitempty"`

	// M升级日期时间戳
	M float32 `json:"M,omitempty"`

	// J升级日期时间戳
	J float32 `json:"J,omitempty"`

	// A1升级日期时间戳
	A1 float32 `json:"A1,omitempty"`

	// A2升级日期时间戳
	A2 float32 `json:"A2,omitempty"`

	// A2b升级日期时间戳
	A2b float32 `json:"A2b,omitempty"`

	// Ta升级日期时间戳
	Ta float32 `json:"Ta,omitempty"`

	// Tb升级日期时间戳
	Tb float32 `json:"Tb,omitempty"`

	// C升级日期时间戳
	C float32 `json:"C,omitempty"`

	// Tc升级日期时间戳
	Tc float32 `json:"Tc,omitempty"`
}

func (m MainRow) GetAircraftExperiences() ([]api.AircraftType, error) {
	// 用jsoniter解析
	var aircraftExperiences []api.AircraftType
	err := jsoniter.Unmarshal([]byte(m.AircraftExperiences), &aircraftExperiences)
	if err != nil {
		return nil, err
	}
	return aircraftExperiences, nil
}

func (m *MainRow) SetAircraftExperiences(aircraftExperiences []api.AircraftType) error {
	// 用jsoniter序列化
	bytes, err := jsoniter.Marshal(aircraftExperiences)
	if err != nil {
		return err
	}
	m.AircraftExperiences = string(bytes)
	return nil
}
