package main

import (
	"github.com/jinzhu/gorm"
	"github.com/newsquid/mscore"
	"github.com/stretchr/testify/suite"
	"testing"
)

/*
Test setup
*/

type UnitTests struct {
	suite.Suite
	DB          *gorm.DB
	Transaction *gorm.DB
}

func (suite *UnitTests) SetupSuite() {
	suite.DB = mscore.SetupUnitTestSuite()
	Migrate(suite.DB)
}

func (suite *UnitTests) TearDownSuite() {
	mscore.TearDownUnitTestSuite()
}

func (suite *UnitTests) SetupTest() {
	suite.Transaction = mscore.SetupUnitTest(suite.DB)
}

func (suite *UnitTests) TearDownTest() {
	mscore.TearDownUnitTest(suite.Transaction)
}

func TestUnitTests(t *testing.T) {
	suite.Run(t, new(UnitTests))
}

/*
Actual tests
*/

func (suite *UnitTests) TestExampleSuccess() {
	suite.Nil(suite.Transaction.Create(&Example{Name: "test"}).Error)

	responseExample, err := ExampleEndpoint(suite.Transaction)
	suite.Nil(err)
	suite.Equal("test", responseExample.Name)
}

func (suite *UnitTests) TestExampleNotFound() {
	_, err := ExampleEndpoint(suite.Transaction)
	suite.NotNil(err)
	suite.Equal(404, err.StatusCode())
}
