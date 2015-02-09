package mscore

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/cors"
	"log"
	"net/http"
	"os"
)

/*
Initialize and return a server (martini)
*/
func InitServer(DB *gorm.DB) (*martini.Martini, martini.Router) {
	m := martini.New()

	//Set middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"OPTIONS", "HEAD", "POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	m.Use(QueryParameters())
	m.Use(UserIPService())

	// Map the database connection
	m.Map(DB)

	//Set returnhandler
	m.Map(JSONReturnHandler())

	//Create the router
	r := martini.NewRouter()

	//Options matches all and sends okay
	r.Options(".*", func() (int, string) {
		return 200, "ok"
	})

	return m, r
}

/*
StartServer starts the server from a http handler. Gets port from env variable
PORT or default 3000
*/
func StartServer(m *martini.Martini, r martini.Router) {
	m.Action(r.Handle)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), m)

	log.Fatalf("HTTP ListAndServe Failed: %s", err.Error())
}
