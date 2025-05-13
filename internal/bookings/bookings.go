package bookings

import (
	"errors"
	"sync"
	"time"
)

type Booking struct {
	ID        int       `json:"id" db:"id"`
	UserName  string    `json:"user_name" db:"user_name"`
	Event     string    `json:"event" db:"event"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	IsActive  bool      `json:"is_active" db:"is_active"`
}

// Store manages bookings in memory
type Store struct {
	bookings map[int]Booking
	mu       sync.RWMutex // For thread-safe access
}

var store = &Store{
	bookings: make(map[int]Booking),
}

// CreateBooking creates a new booking
func CreateBooking(id int, user, event string) (Booking, error) {
	if id <= 0 {
		return Booking{}, errors.New("ID must be positive")
	}
	if user == "" || event == "" {
		return Booking{}, errors.New("user and event cannot be empty")
	}
	b := Booking{
		ID:        id,
		UserName:  user,
		Event:     event,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	store.mu.Lock()
	if _, exists := store.bookings[id]; exists {
		store.mu.Unlock()
		return Booking{}, errors.New("booking ID already exists")
	}
	store.bookings[id] = b
	store.mu.Unlock()
	return b, nil
}

// GetBooking retrieves a booking by ID
func GetBooking(id int) (Booking, error) {
	store.mu.RLock()
	b, exists := store.bookings[id]
	store.mu.RUnlock()
	if !exists {
		return Booking{}, errors.New("booking not found")
	}
	return b, nil
}

// GetAllBookings returns all bookings
func GetAllBookings() []Booking {
	store.mu.RLock()
	bookings := make([]Booking, 0, len(store.bookings))
	for _, b := range store.bookings {
		bookings = append(bookings, b)
	}
	store.mu.RUnlock()
	return bookings
}

// UpdateBooking
func UpdateBooking(id int, user_name string, event string, is_active bool) (Booking, error) {
	store.mu.Lock()
	_, exists := store.bookings[id]
	if !exists {
		store.mu.Unlock()
		return Booking{}, errors.New("Booking not found")
	}
	booking := store.bookings[id]
	booking.Event = event
	booking.UserName = user_name
	booking.IsActive = is_active
	store.bookings[id] = booking

	store.mu.Unlock()

	return booking, nil
}

// DeleteBooking deletes a booking by ID
func DeleteBooking(id int) error {
	store.mu.Lock()
	if _, exists := store.bookings[id]; !exists {
		store.mu.Unlock()
		return errors.New("booking not found")
	}
	delete(store.bookings, id)
	store.mu.Unlock()
	return nil
}
