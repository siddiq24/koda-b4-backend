-- Active: 1763071950390@@ep-bold-cake-a198oos2-pooler.ap-southeast-1.aws.neon.tech@5432@coffee_shop
CREATE TABLE variants (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) not NULL
);

CREATE TABLE orders_products (
    invoice VARCHAR(50),
    product_id BIGINT REFERENCES products(id),
    size_id INT NOT NULL REFERENCES sizes(id),
    varian_id INT REFERENCES variants(id),
    qty INT,
    subTotal NUMERIC,
    name VARCHAR(100)
);

CREATE TABLE products_variants (
    id SERIAL PRIMARY KEY,
    product_id BIGINT REFERENCES products(id),
    variant_id INT NOT NULL REFERENCES variants(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (product_id, variant_id)
);

ALTER TABLE orders_products
ALTER COLUMN product_id TYPE BIGINT;

ALTER TABLE products_images
ALTER COLUMN product_id TYPE BIGINT;

ALTER TABLE products_promos
ALTER COLUMN product_id TYPE BIGINT;

ALTER TABLE products_tags
ALTER COLUMN product_id TYPE BIGINT;
