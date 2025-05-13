package bookings

import (
	"context"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestCreateBooking(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	_, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	booking, err := store.CreateBooking(context.Background(), "alice", "Concert")
	if err != nil {
		t.Fatalf("Failed to create booking: %v", err)
	}
	if booking.UserName != "alice" || booking.Event != "Concert" {
		t.Errorf("Expected booking with user_name=alice, event=Concert, got %+v", booking)
	}
}

func TestGetBooking(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	booking, err := store.CreateBooking(context.Background(), "bob", "Theater")
	if err != nil {
		t.Fatalf("Failed to create booking: %v", err)
	}
	retrievedBooking, err := store.GetBooking(context.Background(), booking.ID)
	if err != nil {
		t.Fatalf("Failed to get booking: %v", err)
	}
	if retrievedBooking.UserName != "bob" || retrievedBooking.Event != "Theater" {
		t.Errorf("Expected booking with user_name=bob, event=Theater, got %+v", retrievedBooking)
	}
}
func TestGetAllBookings(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	_, err = store.CreateBooking(context.Background(), "charlie", "Exhibition")
	if err != nil {
		t.Fatalf("Failed to create booking: %v", err)
	}
	bookings, err := store.GetAllBookings(context.Background())
	if err != nil {
		t.Fatalf("Failed to get all bookings: %v", err)
	}
	if len(bookings) == 0 {
		t.Errorf("Expected at least one booking, got %d", len(bookings))
	}
}
func TestUpdateBooking(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	booking, err := store.CreateBooking(context.Background(), "dave", "Workshop")
	if err != nil {
		t.Fatalf("Failed to create booking: %v", err)
	}
	updatedBooking, err := store.UpdateBooking(context.Background(), booking.ID, "dave", "Updated Workshop")
	if err != nil {
		t.Fatalf("Failed to update booking: %v", err)
	}
	if updatedBooking.Event != "Updated Workshop" {
		t.Errorf("Expected updated booking with event=Updated Workshop, got %+v", updatedBooking)
	}
}
func TestDeleteBooking(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	booking, err := store.CreateBooking(context.Background(), "eve", "Seminar")
	if err != nil {
		t.Fatalf("Failed to create booking: %v", err)
	}
	err = store.DeleteBooking(context.Background(), booking.ID)
	if err != nil {
		t.Fatalf("Failed to delete booking: %v", err)
	}
}
func TestDeleteBookingNotFound(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	err = store.DeleteBooking(context.Background(), 9999) // Assuming this ID does not exist
	if err == nil {
		t.Fatal("Expected error when deleting non-existent booking, got nil")
	}
}
func TestCreateBookingInvalid(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	_, err = store.CreateBooking(context.Background(), "", "") // Invalid booking
	if err == nil {
		t.Fatal("Expected error when creating invalid booking, got nil")
	}
}
func TestGetBookingNotFound(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	_, err = store.GetBooking(context.Background(), 9999) // Assuming this ID does not exist
	if err == nil {
		t.Fatal("Expected error when getting non-existent booking, got nil")
	}
}
func TestUpdateBookingNotFound(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	_, err = store.UpdateBooking(context.Background(), 9999, "nonexistent", "Updated Event") // Assuming this ID does not exist
	if err == nil {
		t.Fatal("Expected error when updating non-existent booking, got nil")
	}
}
func TestUpdateBookingInvalid(t *testing.T) {
	dsn := "postgres://booking_user:securepassword@localhost:5432/booking_app?sslmode=disable"
	store, err := NewDBStore(dsn)
	if err != nil {
		t.Fatalf("Failed to create booking store: %v", err)
	}
	_, err = store.UpdateBooking(context.Background(), 1, "", "") // Invalid booking
	if err == nil {
		t.Fatal("Expected error when updating invalid booking, got nil")
	}
}
