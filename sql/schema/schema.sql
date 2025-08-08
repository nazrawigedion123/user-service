CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP ,
    deleted_at TIMESTAMP NULL

);


-- Unique when not deleted
CREATE UNIQUE INDEX IF NOT EXISTS unique_user_name_if_not_deleted
    ON users(user_name)
    WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS unique_phone_if_not_deleted
    ON users(phone)
    WHERE deleted_at IS NULL;