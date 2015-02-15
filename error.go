package mscore

import (
	"log"
	"net/http"
	"runtime"
	"strings"
)

/*
Error, a special error used to contain an error message and a status code
The message is to be passed
*/
type Error interface {
	StatusCode() int
	Error() string
}

/*
httpError, an implementation of the Error interface
*/
type httpError struct {
	statusCode   int
	errorMessage string
}

/*
Error get the error message of the error
*/
func (t *httpError) Error() string {
	return t.errorMessage
}

/*
StatusCode get the statuscode of the error
*/
func (t *httpError) StatusCode() int {
	return t.statusCode
}

/*
NewError creates an error with statusCode and errorMessage
*/
func NewError(statusCode int, errorMessage string) Error {
	return &httpError{statusCode, errorMessage}
}

/*
NewErrorFromStatusWithMessage creates an error with statusCode and errorMessage
*/
func NewErrorFromStatusWithMessage(code int, msg string) Error {
	return &httpError{code, msg}
}

/*
NewErrorFromStatus creates an error from a statuscode with the default message for the status code.
*/
func NewErrorFromStatus(code int) Error {
	return &httpError{code, http.StatusText(code)}
}

/*
InternalServerError creates an internal server error with the errorMessage and
logs the error struct for future investigation.
*/
func InternalServerError(errorMessage string, err error) Error {
	logError(err)
	return &httpError{http.StatusInternalServerError, errorMessage}
}

/*
InternalServerErr creates an internal server error with the default StatusText and
logs the error struct for future investigation.
*/
func InternalServerErr(err error) Error {
	logError(err)
	return &httpError{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)}
}

/*
Log an error
*/
func LogError(err error) {
	logError(err)
}

/*
LogError, logs an error to the console
*/
func logError(err error) {
	//Get stack trace 2 up.
	_, file, line, ok := runtime.Caller(2)
	if ok {
		//Parse the filename out after the ttserver folder
		fileSplit := strings.Split(file, "ttserver/")
		fileName := fileSplit[len(fileSplit)-1]

		log.Printf("[Internal Error] %s:%d (%s)", fileName, line, err.Error())
	} else {
		log.Printf("[Internal Error] %s", err.Error())
	}
}
