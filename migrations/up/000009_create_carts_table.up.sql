CREATE TABLE carts(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    product_id BIGINT REFERENCES products(id),
    size_id INT REFERENCES sizes(id),
    varian_id INT REFERENCES variants(id),
    qty INT NOT NULL DEFAULT 0,
    subtotal NUMERIC NOT NULL DEFAULT 0,
    product_name VARCHAR(100)
);