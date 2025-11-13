CREATE TABLE profiles (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT UNIQUE REFERENCES users(id),
    fullname VARCHAR(50),
    image VARCHAR(100),
    phone VARCHAR(20) UNIQUE,
    address VARCHAR
);