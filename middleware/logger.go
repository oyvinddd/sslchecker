package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Logger(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println(stringFromRequest(r))
		next(w, r, ps)
	}
}

func stringFromRequest(r *http.Request) string {
	return fmt.Sprintf("[%s] %s", r.Method, r.URL.Path)
}
