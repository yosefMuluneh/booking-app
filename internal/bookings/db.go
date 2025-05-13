package bookings

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DBStore manages bookings in PostgreSQL
type DBStore struct {
	db *sqlx.DB
}

func NewDBStore(dsn string) (*DBStore, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DBStore{db: db}, nil
}

func (s *DBStore) CreateBooking(ctx context.Context, user, event string) (Booking, error) {
	if user == "" || event == "" {
		return Booking{}, errors.New("user and event cannot be empty")
	}
	b := Booking{
		UserName:  user,
		Event:     event,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	query := `INSERT INTO bookings (user_name, event, created_at, updated_at, is_active) 
              VALUES (:user_name, :event, :created_at, :updated_at, :is_active) 
              RETURNING *`
	rows, err := s.db.NamedQueryContext(ctx, query, b)
	if err != nil {
		return Booking{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Failed to close rows in CreateBooking: %v", err)
		}
	}()
	if rows.Next() {
		if err := rows.StructScan(&b); err != nil {
			return Booking{}, err
		}
	} else {
		return Booking{}, errors.New("failed to retrieve inserted booking")
	}
	return b, nil
}

func (s *DBStore) GetBooking(ctx context.Context, id int) (Booking, error) {
	var b Booking
	err := s.db.GetContext(ctx, &b, "SELECT * FROM bookings WHERE id = $1", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return Booking{}, errors.New("booking not found")
		}
		return Booking{}, err
	}
	return b, nil
}

func (s *DBStore) GetAllBookings(ctx context.Context) ([]Booking, error) {
	var bookings []Booking
	err := s.db.SelectContext(ctx, &bookings, "SELECT * FROM bookings")
	return bookings, err
}

func (s *DBStore) UpdateBooking(ctx context.Context, id int, user, event string) (Booking, error) {
	if user == "" || event == "" {
		return Booking{}, errors.New("user and event cannot be empty")
	}
	b := Booking{
		ID:        id,
		UserName:  user,
		Event:     event,
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	query := `UPDATE bookings SET user_name = :user_name, event = :event, updated_at = :updated_at, is_active = :is_active 
              WHERE id = :id RETURNING *`
	rows, err := s.db.NamedQueryContext(ctx, query, b)
	if err != nil {
		return Booking{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Failed to close rows in UpdateBooking: %v", err)
		}
	}()
	if rows.Next() {
		if err := rows.StructScan(&b); err != nil {
			return Booking{}, err
		}
	} else {
		return Booking{}, errors.New("booking not found")
	}
	return b, nil
}

func (s *DBStore) DeleteBooking(ctx context.Context, id int) error {
	result, err := s.db.ExecContext(ctx, "DELETE FROM bookings WHERE id = $1", id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("booking not found")
	}
	return nil
}
