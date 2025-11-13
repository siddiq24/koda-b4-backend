CREATE TABLE deliveries (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) UNIQUE
);

CREATE TABLE status (
    id int generated always as identity PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE payment_methods(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    image VARCHAR(100) NOT NULL,
    no_va VARCHAR(50) NOT NULL
);

CREATE TABLE orders(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    email VARCHAR(100),
    fullname VARCHAR(100),
    phone VARCHAR(20),
    address TEXT,
    payment_method_id INT REFERENCES payment_methods(id),
    delivery_id INT NOT NULL REFERENCES deliveries(id) DEFAULT 1,
    total_order NUMERIC NOT NULL DEFAULT 0,
    invoice VARCHAR(50),
    status_id INT NOT NULL REFERENCES status(id) DEFAULT 1,
    promo_id INT REFERENCES promos(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);