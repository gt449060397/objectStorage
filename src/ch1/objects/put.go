package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {

	log.Println(strings.Split(r.URL.EscapedPath(), "/"))
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	log.Println("put test")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)

}
