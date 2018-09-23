package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	responses := []helloResponse{helloResponse{Message: "Message 1"}, helloResponse{Message: "Message 2"}}

	data, err := json.Marshal(responses)

	if err != nil {
		panic("err")
	}

	fmt.Fprintf(w, string(data))

}

func main() {
	port := 9000
	http.HandleFunc("/helloworld", helloFunc)
	log.Printf("Server Staring on port %v\n", port)
	http.ListenAndServe(":9000", nil)
}
