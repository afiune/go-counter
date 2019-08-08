package main

import (
	"fmt"
	"net/http"
)

var REQUESTS int

func main() {
	url := ":1234"
	fmt.Printf("Starting app listening on '%s'\n", url)
	http.HandleFunc("/", processRequest)
	if err := http.ListenAndServe(url, nil); err != nil {
		panic(err)
	}
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	// TODO(afiune) make it atomic
	REQUESTS++
	msg := fmt.Sprintf("{\"req_num\":%d,\"status\":\"processed\"}", REQUESTS)
	fmt.Fprintf(w, msg+"\n")
	fmt.Println(msg)
}
