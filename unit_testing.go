package mscore

import (
	"github.com/newsquid/gorm"
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
