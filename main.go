package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello, gitops\n")
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe err: %v", err)
		os.Exit(1)
	}
}
