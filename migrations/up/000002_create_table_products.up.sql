CREATE TABLE categories(
    id int generated always as identity PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE products (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(255),
    base_price NUMERIC,
    stock INT,
    category_id INT REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);






CREATE TABLE sizes (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(10) NOT NULL UNIQUE,
    additional_price NUMERIC NOT NULL
);

CREATE TABLE products_sizes (
    product_id BIGINT REFERENCES products(id),
    size_id INT REFERENCES sizes(id),
    CONSTRAINT products_sizes_unique UNIQUE(product_id, size_id)
);






CREATE TABLE products_images(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id INT REFERENCES products(id),
    image VARCHAR(100)
);