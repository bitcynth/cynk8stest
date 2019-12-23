package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var buildVersion string
var buildTime string

func main() {
	r := mux.NewRouter()
	hostname, _ := os.Hostname()

	r.HandleFunc("/info", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "version: %s\nbuild_time: %s\nnode_hostname: %s\n", buildVersion, buildTime, hostname)
	})

	http.ListenAndServe(":8080", r)
}
