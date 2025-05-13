CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    event VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL
);