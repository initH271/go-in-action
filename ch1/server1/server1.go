package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", rootHandle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
