CREATE TABLE IF NOT EXISTS customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    unit VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS transaction (
    id SERIAL PRIMARY KEY,
    bill_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    entry_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    finish_date TIMESTAMP,
    employee_id VARCHAR(50) NOT NULL,
    customer_id INT REFERENCES customer(id) ON DELETE CASCADE,
    total_bill INT NOT NULL
);

CREATE TABLE IF NOT EXISTS bill_detail (
    id SERIAL PRIMARY KEY,
    bill_id INT REFERENCES transaction(id) ON DELETE CASCADE,
    product_id INT REFERENCES product(id) ON DELETE CASCADE,
    product_price INT NOT NULL,
    qty INT NOT NULL
);
