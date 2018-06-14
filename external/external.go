package external

import (
	"net/http"
	"fmt"
	"log"
)

const AppName = "awesome"

func StartServer(port string, handlerFunc http.HandlerFunc) {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting web server on port " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
