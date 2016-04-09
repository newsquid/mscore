package main

import (
	"github.com/go-martini/martini"
	"github.com/newsquid/gorm"
	"github.com/newsquid/mscore"
	"os"
)

/*
An example type for storage in the database
*/
type Example struct {
	Name string
}

/*
Application entry point
*/
func main() {
	args := mscore.ParseArguments()
	DB := mscore.InitDB()

	if args.Migrate {
		Migrate(DB)
		os.Exit(0)
	}

	server, router := mscore.InitServer(DB)

	SetupRoutes(router)

	mscore.StartServer(server, router)
}

/*
Migrate the database
*/
func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&Example{})
}

/*
Setup the routes
*/
func SetupRoutes(router martini.Router) {
	router.Get("/example", ExampleEndpoint)
}

/*
Example of an endpoint
*/
func ExampleEndpoint(DB *gorm.DB) (Example, mscore.Error) {
	var example Example
	if err := DB.Find(&example).Error; err != nil {
		if err == gorm.RecordNotFound {
			return example, mscore.NewError(404,
				"Example entry not found")
		}
		return example, mscore.InternalServerErr(err)
	}

	return example, nil
}
