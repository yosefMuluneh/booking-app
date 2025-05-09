package bookings

import "time"

// Booking represents a single booking
type Booking struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Event     string    `json:"event"`
	CreatedAt time.Time `json:"created_at"`
}
