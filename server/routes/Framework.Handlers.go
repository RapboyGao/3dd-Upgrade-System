package routes

import (
	"GinServer/server/api" // 项目API定义和响应模型

	"github.com/gin-gonic/gin" // Gin Web框架
	// 时间处理库，提供更友好的API
)

/**
 * 默认API处理函数集合
 * 定义了系统内置的基础API接口实现
 */
var defaultApis = api.DefaultAPI{
	AircraftExperiencesOverrideDelete: func(*gin.Context) {},
	AircraftExperiencesOverrideGet:    func(*gin.Context) {},
	AircraftExperiencesOverridePost:   func(*gin.Context) {},
	CheckRideDetailDelete:             func(*gin.Context) {},
	CheckRideDetailGet:                func(*gin.Context) {},
	CheckRideDetailPost:               func(*gin.Context) {},
	LevelOverrideDelete:               func(*gin.Context) {},
	LevelOverrideGet:                  func(*gin.Context) {},
	LevelOverridePost:                 func(*gin.Context) {},
	MainSheetGet:                      func(*gin.Context) {},
	PilotInfoGet:                      func(*gin.Context) {},
	UploadExcelPost:                   func(*gin.Context) {},
}
