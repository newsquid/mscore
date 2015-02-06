package core

import (
	"bytes"
	ejson "encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"io"
	"io/ioutil"
)

type RouteAdder func(martini.Router)

func SetupIntegrationTestSuite(ra RouteAdder) (*gorm.DB, *martini.Martini) {
	DB := InitTestDB()

	server, router := InitServer(DB)
	ra(router)

	go StartServer(server, router)

	return DB, server
}

func TearDownIntegrationTestSuite() {
	RemoveTestDB()
}

func SetupIntegrationTest(DB *gorm.DB, m *martini.Martini) *gorm.DB {
	tx := DB.Begin()
	m.Map(tx)
	return tx
}

func TearDownIntegrationTest(DB *gorm.DB) {
	DB.Rollback()
}

/***
 * Helper functions for Migration tests
 **/

const (
	JsonType = "application/json"
)

/*
Creates the request url with the local server as prefix expects url of form
"/url" etc.
*/
func TestUrl(s string) string {
	return fmt.Sprintf("http://localhost:3000%s", s)
}

/*
Parses string from json to go type
example:
	p := Post{}
	json(s, &p)
*/
func FromJson(r io.ReadCloser, i interface{}) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	s := string(data)

	err = ejson.Unmarshal([]byte(s), i)
	if err != nil {
		fmt.Println("Json in: " + s)
		panic(err)
	}
}

/*
Parses go struct to json
*/
func ToJson(i interface{}) *bytes.Reader {
	b, err := ejson.Marshal(i)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(b)
}
