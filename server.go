package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.Write([]byte("Error reading the content"))
	}
	w.Write([]byte(requestDump))
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.Write([]byte("Error reading the content"))
	}

	keys, _ := r.URL.Query()["time"]
	t, _ := strconv.Atoi(keys[0])
	time.Sleep(time.Duration(t) * time.Millisecond)
	w.Write([]byte(requestDump))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8085"
	}
	return ":" + port
}

func main() {
	fmt.Println(" Server started on PORT ", port())
	http.HandleFunc("/sleep/", delayHandler)
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(port(), nil); err != nil {
		panic(err)
	}
}
