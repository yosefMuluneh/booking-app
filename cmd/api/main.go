package main

import (
	"log"
	"net/http"

	"booking-app/internal/bookings"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Welcome to the Booking App!")); err != nil {
		log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func main() {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := bookings.NewDBStore(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	handler := bookings.NewHandler(store)

	r := mux.NewRouter()
	r.Use(bookings.LoggingMiddleware)
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	r.HandleFunc("/bookings", handler.ListBookings).Methods(http.MethodGet)
	r.HandleFunc("/bookings", handler.CreateBookingHandler).Methods(http.MethodPost)
	r.HandleFunc("/bookings/{id}", handler.GetBookingHandler).Methods(http.MethodGet)
	r.HandleFunc("/bookings/{id}", handler.UpdateBookingHandler).Methods(http.MethodPut)
	r.HandleFunc("/bookings/{id}", handler.DeleteBookingHandler).Methods(http.MethodDelete)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
