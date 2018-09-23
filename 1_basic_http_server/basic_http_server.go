package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func main() {
	port := 9000
	http.HandleFunc("/helloworld", helloFunc)

	log.Printf("Server Staring on port %v\n", port)
	http.ListenAndServe(":9000", nil)
}
