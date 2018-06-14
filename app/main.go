package main

import (
	"fmt"
	"net/http"
	"github.com/mikemadisonweb/go-debug-example/external"
)

func handle(w http.ResponseWriter, r *http.Request) {
	name := external.AppName + "_1"
	fmt.Fprintf(w, "You have visited %s in `%s`!", r.URL.Path, name)
}

func main() {
	external.StartServer("8080", handle)
}