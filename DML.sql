
-- Insert sample data into the 'customer' table
INSERT INTO customer (name, phone_number, address) VALUES
    ('John Doe', '123-456-7890', '123 Main St'),
    ('Jane Smith', '987-654-3210', '456 Oak St');

-- Insert sample data into the 'product' table
INSERT INTO product (name, price, unit) VALUES
    ('Product A', 20, 'Piece'),
    ('Product B', 30, 'Piece');

-- Insert sample data into the 'transaction' table
INSERT INTO transaction (bill_date, entry_date, employee_id, customer_id, total_bill) VALUES
    ('2024-03-09 12:00:00', '2024-03-09 12:05:00', 'EMP001', 1, 50),
    ('2024-03-10 14:00:00', '2024-03-10 14:05:00', 'EMP002', 2, 60);

-- Insert sample data into the 'bill_detail' table
INSERT INTO bill_detail (bill_id, product_id, product_price, qty) VALUES
    (1, 1, 20, 2),
    (1, 2, 30, 1),
    (2, 1, 20, 3),
    (2, 2, 30, 2);
