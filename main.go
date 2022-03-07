package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", HandleRoot)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	// Get Client IP Address
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	fmt.Fprintf(w, "\nIP Address: Port of Client: %s\n\n", IPAddress)
	fmt.Printf("%s %s, ", time.Now().Format("2006-01-02 15:04:05"), IPAddress)

	// sort r.Header to make it more easy to track
	keys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Fprintf(w, "Headers:\n")
	for _, k := range keys {
		fmt.Fprintf(w, "%q: %q\n", k, r.Header[k])
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "\nGET Request: %q\n", html.EscapeString(r.URL.Path))
		fmt.Printf("GET: %q \n", html.EscapeString(r.URL.Path))
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "\nPOST Request: %q\n", html.EscapeString(r.URL.Path))
		fmt.Printf("POST: %q \n", html.EscapeString(r.URL.Path))
	} else {
		http.Error(w, "Invalid request method.\n", 405)
	}

	fmt.Fprintf(w, "\nAll Environment Vars:\n")
	for _, env := range os.Environ() {
		// env is
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]

		fmt.Fprintf(w, "%s: %s\n", key, value)
	}

}
