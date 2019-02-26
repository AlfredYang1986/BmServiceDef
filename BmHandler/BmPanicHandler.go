package BmHandler

import "net/http"

type BmPanicHandler interface {
	HandlePanic(rw http.ResponseWriter, r *http.Request, p interface{})
}
