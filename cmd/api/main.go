package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"booking-app/internal/bookings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Booking App!")
}

func bookingsHandler(w http.ResponseWriter, r *http.Request) {
	booking := bookings.Booking{
		ID:        1,
		UserName:  "John Doe",
		Event:     "Concert",
		CreatedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(booking); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/bookings", bookingsHandler)
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
