package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func echoHandler(w http.ResponseWriter, r *http.Request) {
	delayInMsec, sizeInBytes := getDelay(r)
	time.Sleep(time.Duration(delayInMsec) * time.Millisecond)

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.Write([]byte("Error reading the content"))
	}

	time.Sleep(time.Duration(delayInMsec) * time.Millisecond)
	if sizeInBytes > len(requestDump) {
		sizeInBytes = len(requestDump)
	}
	w.Write([]byte(requestDump[:sizeInBytes]))
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	delayInMsec, _ := getDelay(r)
	time.Sleep(time.Duration(delayInMsec) * time.Millisecond)
	printSummary(w, r)
}

func printSummary(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test Backend ")
	fmt.Fprintf(w, "RequestURI : %q\r\n", r.URL.RequestURI())
	jsonURI, _ := json.Marshal(r.URL.Query())
	fmt.Fprintln(w, "Query : \r\n"+string(jsonURI))
	fmt.Fprintf(w, "EscapedPath : %q\r\n", r.URL.EscapedPath())
	fmt.Fprintf(w, "URLString : %q\r\n", r.URL.String())
	fmt.Fprintf(w, "Port :  %q\r\n", r.URL.Port())
	fmt.Fprintf(w, "HostName :  %q\r\n", r.URL.Hostname())
	jsonHeader, _ := json.Marshal(r.Header)
	fmt.Fprintln(w, "Header : \r\n"+string(jsonHeader))
}

func getDelay(r *http.Request) (delayInMsec, sizeInBytes int) {
	escapedPath := r.URL.EscapedPath()
	splitBySlash := strings.Split(escapedPath, "/")
	readDelay, readSize := false, false
	delayInMsec = 10
	sizeInBytes = 1000
	for _, element := range splitBySlash {
		if readDelay {
			number, err := strconv.Atoi(element)
			if err == nil {
				delayInMsec = number
			}
			readDelay = false
			continue
		}

		if readSize {
			number, err := strconv.Atoi(element)
			if err == nil {
				sizeInBytes = number
			}
			readSize = false
			continue
		}

		if element == "size" {
			readSize = true
		}

		if element == "delay" {
			readDelay = true
		}
	}

	return delayInMsec, sizeInBytes
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	delayInMsec, sizeInBytes := getDelay(r)
	time.Sleep(time.Duration(delayInMsec) * time.Millisecond)
	printSummary(w, r)
	fmt.Fprintln(w, "RandomContent: ")
	if sizeInBytes < 5 {
		fmt.Fprintf(w, "%s", randSeq(5))
	} else {
		for i := 0; i < sizeInBytes/5; i++ {
			fmt.Fprintf(w, "%s ", randSeq(5))
		}
	}
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/         - Echo + return random content for the specified size ")
	fmt.Fprintln(w, "/echo     - Echo the request  ")
	fmt.Fprintln(w, "/summary  - Summarize the request ")
	fmt.Fprintln(w, "Optional Endpoints:")
	fmt.Fprintln(w, "/delay/<integer>   - Generate a server delay of <integer> msec ")
	fmt.Fprintln(w, "/size/<integer>    - Specify the size of the response in bytes ")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8088"
	}
	return ":" + port
}

func main() {
	fmt.Println(" Server started on PORT ", port())
	http.HandleFunc("/help/", helpHandler)
	http.HandleFunc("/echo/", echoHandler)
	http.HandleFunc("/summary/", summaryHandler)
	http.HandleFunc("/", httpRequestHandler)
	if err := http.ListenAndServe(port(), nil); err != nil {
		panic(err)
	}
}
