package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endPointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Successful",
		Body:   "You have hit the toll booth API!",
	}
	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	message := Message{
		Status: "Request Failed!",
		Body:   "The API is at limit",
	}
	jsonMessage, _ := json.Marshal(message)
	tlbthLimiter := tollbooth.NewLimiter(1, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))
	http.Handle("/pingtoll", tollbooth.LimitFuncHandler(tlbthLimiter, endPointHandler))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Error while running the server")
	}
}
