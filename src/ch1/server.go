package main

import (
	"ch1/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)

	log.Println(os.Getenv("LISTEN_ADDRESS"))
	log.Fatal(http.ListenAndServe(":12345", nil))
}
