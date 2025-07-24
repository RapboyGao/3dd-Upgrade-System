package model

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	// 确保.build/temp目录存在
	os.MkdirAll(".build/temp", os.ModePerm)
	db, err := gorm.Open(sqlite.Open(".build/temp/gorm.db"), &gorm.Config{})
	// 迁移 schema
	db.AutoMigrate(&LevelsRawData{})
	return db, err
}
