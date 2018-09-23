package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	response := helloResponse{Message: "Message 1"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}

func main() {
	port := 9000
	http.HandleFunc("/helloworld", helloFunc)
	log.Printf("Server Staring on port %v\n", port)
	http.ListenAndServe(":9000", nil)
}
