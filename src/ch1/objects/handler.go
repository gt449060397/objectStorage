package objects

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	log.Println(m)
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
