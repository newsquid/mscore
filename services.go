package mscore

import (
	"github.com/go-martini/martini"
	"net/http"
)

/*
QueryParameters service to get query parameters from requests
*/
func QueryParameters() func(martini.Context, *http.Request) {
	return func(c martini.Context, req *http.Request) {
		c.Map(req.URL.Query())
	}
}

/*
UserIP IP address of a user.
*/
type UserIP struct {
	IP string
}

/*
UserIP service gets the user ip from the request
*/
func UserIPService() func(martini.Context, *http.Request) {
	return func(c martini.Context, req *http.Request) {
		c.Map(UserIP{req.RemoteAddr})
	}
}
