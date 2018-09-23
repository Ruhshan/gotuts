package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
}

type helloRequest struct {
	Name string `json:"name"`
}

func helloFunc(w http.ResponseWriter, r *http.Request) {

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	// var request helloRequest
	// err = json.Unmarshal(body, &request)
	// if err != nil {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	var request helloRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	response := helloResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	port := 9000
	http.HandleFunc("/helloworld", helloFunc)
	log.Printf("Server Staring on port %v\n", port)
	http.ListenAndServe(":9000", nil)
}
