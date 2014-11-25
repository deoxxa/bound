package bound

import (
	"net/http"
)

type Anything interface{}
type Handler struct {
	Anything
	Fn http.HandlerFunc
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Fn(w, r)
}
