package main

import (
	"net/http"
	"net/http/httputil"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.Write([]byte("Error reading the content"))
	}
	w.Write([]byte(requestDump))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
