package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gt449060397/objectStorage/src/ch2/dataServer/heartbeat"
	"github.com/gt449060397/objectStorage/src/ch2/dataServer/locate"
	"github.com/gt449060397/objectStorage/src/ch2/dataServer/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()

	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
