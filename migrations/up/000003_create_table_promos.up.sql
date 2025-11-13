CREATE TABLE promos (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    discount FLOAT NOT NULL,
    start TIMESTAMP,
    "end" TIMESTAMP
);

CREATE TABLE products_promos (
    product_id INT REFERENCES products(id),
    promo_id INT REFERENCES promos(id)
);