CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    quantity INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

INSERT INTO product (name, quantity,price) VALUES
('MacBook Pro', 3,2999.99),
('GoLang Hoodie',2, 49.90),
('Mechanical Keyboard',1, 89.50);
