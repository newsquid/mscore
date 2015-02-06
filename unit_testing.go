package core

import (
	"github.com/jinzhu/gorm"
)

func SetupUnitTestSuite() *gorm.DB {
	return InitTestDB()
}

func TearDownUnitTestSuite() {
	RemoveTestDB()
}

func SetupUnitTest(DB *gorm.DB) *gorm.DB {
	return DB.Begin()
}

func TearDownUnitTest(DB *gorm.DB) {
	DB.Rollback()
}
