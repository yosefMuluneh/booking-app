package main

import (
	"booking-app/internal/bookings"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(bookings.LoggingMiddleware) // Apply to all routes
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Booking App!"))
	}).Methods(http.MethodGet)
	r.HandleFunc("/bookings", bookings.ListBookings).Methods(http.MethodGet)
	r.HandleFunc("/bookings", bookings.CreateBookingHandler).Methods(http.MethodPost)
	r.HandleFunc("/bookings/{id}", bookings.GetBookingHandler).Methods(http.MethodGet)
	r.HandleFunc("/bookings/{id}", bookings.UpdateBookingHandler).Methods(http.MethodPut)
	r.HandleFunc("/bookings/{id}", bookings.DeleteBookingHandler).Methods(http.MethodDelete)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
