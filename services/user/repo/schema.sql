-- Create the table
CREATE TABLE users (
    id SERIAL PRIMARY KEY, -- Unique identifier for each user
    username VARCHAR(50) UNIQUE NOT NULL, -- Unique username
    email VARCHAR(255) UNIQUE NOT NULL, -- Unique email
    password_hash TEXT NOT NULL, -- Hashed password for security
    first_name VARCHAR(50), -- Optional first name
    last_name VARCHAR(50), -- Optional last name
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Account creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    is_active BOOLEAN DEFAULT TRUE, -- Whether the account is active
    is_admin BOOLEAN DEFAULT FALSE, -- Whether the user has admin privileges
    last_login TIMESTAMP, -- Last login timestamp
    profile_picture_url TEXT, -- URL to the user's profile picture
    bio TEXT, -- User bio or description
    date_of_birth DATE, -- Optional date of birth
    phone_number VARCHAR(15), -- Optional phone number
    address JSONB -- JSON object to store address details if needed
);

-- Create a function to update the `updated_at` field
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to call the function before updating a row
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

