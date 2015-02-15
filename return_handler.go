package mscore

import (
	"encoding/json"
	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"net/http"
	"reflect"
)

/*
JSONReturnHandler is a Martini returnhandler that automatically makes returned values into JSON.
The return handler handles 4 different cases:
- One value and one integer returned. Here the integer becomes the http status
  code and the value as JSON is set as body.
- One value and one mscore.Error is returned. If the mscore.Error != nil the
  http status code is set to mscore.Error.StatusCode() and the body is
  mscore.Error.Error() as text. Otherwise the value as JSON is set as body and
  the http status code is 200.
- One value is returned. The value as JSON is set as body and the http status
  code is 200.
- One mscore.Error is returned. If the mscore.Error != nil the
  http status code is set to mscore.Error.StatusCode() and the body is
  mscore.Error.Error() as text. Otherwise an empty body with status code 200.
*/
func JSONReturnHandler() martini.ReturnHandler {
	return func(ctx martini.Context, vals []reflect.Value) {
		rv := ctx.Get(inject.InterfaceOf((*http.ResponseWriter)(nil)))
		w := rv.Interface().(http.ResponseWriter)

		if len(vals) > 1 && vals[0].Kind() == reflect.Int {
			//One value and one integer.
			result, err := json.Marshal(vals[1].Interface())

			if err != nil {
				http.Error(w, err.Error(), 500)
			}

			//Set the integer as status code
			w.WriteHeader(int(vals[0].Int()))

			//Write the JSON
			w.Write(result)
		} else if len(vals) > 1 {
			//One value and one mscore.Error
			if !vals[1].IsNil() {
				//If the mscore.Error != nil we use it as http error.
				empty := []reflect.Value{}
				errmsg := vals[1].MethodByName("Error").Call(empty)[0]
				errcode := vals[1].MethodByName("StatusCode").Call(empty)[0]
				http.Error(w, errmsg.String(), int(errcode.Int()))
			} else {
				//Marshall the value
				result, err := json.Marshal(vals[0].Interface())

				if err != nil {
					http.Error(w, err.Error(), 500)
				}

				w.WriteHeader(200)
				w.Write(result)
			}

		} else if len(vals) > 0 {
			value := vals[0].Interface()
			switch t := value.(type) {
			case Error:
				//One error is returned.
				if t != nil {
					http.Error(w, t.Error(), t.StatusCode())
				}
			default:
				//Some value write it as JSON.
				result, err := json.Marshal(t)

				if err != nil {
					http.Error(w, err.Error(), 500)
				}

				w.Write(result)
			}

		}
	}
}
