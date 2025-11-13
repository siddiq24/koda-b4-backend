CREATE TABLE variants (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) not NULL
);

CREATE TABLE orders_products (
    invoice VARCHAR(50),
    product_id INT REFERENCES products(id),
    size_id INT NOT NULL REFERENCES sizes(id),
    varian_id INT REFERENCES variants(id),
    qty INT,
    subTotal NUMERIC,
    name VARCHAR(100)
);