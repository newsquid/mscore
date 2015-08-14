package mscore

import (
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func Redirect(url string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, 302)
	}
}

/*
Extra logging
*/
func LoggerService() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		start := time.Now()

		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}

		bodydata, err := ioutil.ReadAll(req.Body)
		var body string
		if req.Method == "POST" && err != nil {
			body = "body: " + string(bodydata)
		}

		log.Printf("Started %s %s for %s %s", req.Method, req.URL.String(), addr, body)

		rw := res.(martini.ResponseWriter)
		c.Next()

		log.Printf("Completed %v %s in %v\n", rw.Status(), http.StatusText(rw.Status()), time.Since(start))
	}
}
