package mscore

import (
	"encoding/json"
	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"net/http"
	"reflect"
)

/*
JSONReturnHandler converts return values into JSON
*/
func JSONReturnHandler() martini.ReturnHandler {
	return func(ctx martini.Context, vals []reflect.Value) {
		rv := ctx.Get(inject.InterfaceOf((*http.ResponseWriter)(nil)))
		w := rv.Interface().(http.ResponseWriter)

		if len(vals) > 1 && vals[0].Kind() == reflect.Int {
			result, err := json.Marshal(vals[1].Interface())

			if err != nil {
				http.Error(w, err.Error(), 500)
			}

			w.WriteHeader(int(vals[0].Int()))
			w.Write(result)
		} else if len(vals) > 1 {

			if !vals[1].IsNil() {
				empty := []reflect.Value{}
				errmsg := vals[1].MethodByName("Error").Call(empty)[0]
				errcode := vals[1].MethodByName("StatusCode").Call(empty)[0]
				http.Error(w, errmsg.String(), int(errcode.Int()))
			} else {
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
				if t != nil {
					http.Error(w, t.Error(), t.StatusCode())
				}
			default:
				result, err := json.Marshal(t)

				if err != nil {
					http.Error(w, err.Error(), 500)
				}

				w.Write(result)
			}

		}
	}
}
