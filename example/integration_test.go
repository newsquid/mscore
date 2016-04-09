package main

import (
	"github.com/go-martini/martini"
	"github.com/newsquid/gorm"
	"github.com/newsquid/mscore"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

/*
Test setup
*/

type IntegrationTests struct {
	suite.Suite
	DB          *gorm.DB
	Transaction *gorm.DB
	m           *martini.Martini
}

func (suite *IntegrationTests) SetupSuite() {
	suite.DB, suite.m = mscore.SetupIntegrationTestSuite(SetupRoutes)
	Migrate(suite.DB)
}

func (suite *IntegrationTests) TearDownSuite() {
	mscore.TearDownIntegrationTestSuite()
}

func (suite *IntegrationTests) SetupTest() {
	suite.Transaction = mscore.SetupIntegrationTest(suite.DB, suite.m)
}

func (suite *IntegrationTests) TearDownTest() {
	mscore.TearDownIntegrationTest(suite.Transaction)
}

func TestIntegrationTests(t *testing.T) {
	suite.Run(t, new(IntegrationTests))
}

/*
Actual tests
*/

func (suite *IntegrationTests) TestExampleSuccess() {
	suite.Nil(suite.Transaction.Create(&Example{Name: "test"}).Error)

	resp, err := http.Get(mscore.TestUrl("/example"))
	suite.Nil(err)
	suite.Equal(200, resp.StatusCode)

	var responseExample Example
	mscore.FromJson(resp.Body, &responseExample)
	suite.Equal("test", responseExample.Name)
}

func (suite *IntegrationTests) TestExampleNotFound() {
	resp, err := http.Get(mscore.TestUrl("/example"))
	suite.Nil(err)
	suite.Equal(404, resp.StatusCode)
}
