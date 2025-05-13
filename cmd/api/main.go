package main

import (
	"log"
	"net/http"

	"booking-app/internal/bookings"
	"booking-app/internal/middleware"
	"booking-app/internal/users"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Welcome to the Booking App!")); err != nil {
		log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func main() {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	bookingStore, err := bookings.NewDBStore(dsn)
	if err != nil {
		log.Fatalf("Failed to create booking store: %v", err)
	}
	userStore := users.NewDBStore(db)
	jwtSecret := "your-secret-key" // Use env variable in production

	bookingHandler := bookings.NewHandler(bookingStore)
	userHandler := users.NewHandler(userStore, jwtSecret)

	r := mux.NewRouter()
	r.Use(bookings.LoggingMiddleware)
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	r.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)

	// Protected routes
	protected := r.PathPrefix("/bookings").Subrouter()
	protected.Use(middleware.Auth(jwtSecret))
	protected.HandleFunc("", bookingHandler.ListBookings).Methods(http.MethodGet)
	protected.HandleFunc("", bookingHandler.CreateBookingHandler).Methods(http.MethodPost)
	protected.HandleFunc("/{id}", bookingHandler.GetBookingHandler).Methods(http.MethodGet)
	protected.HandleFunc("/{id}", bookingHandler.UpdateBookingHandler).Methods(http.MethodPut)
	protected.HandleFunc("/{id}", bookingHandler.DeleteBookingHandler).Methods(http.MethodDelete)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
