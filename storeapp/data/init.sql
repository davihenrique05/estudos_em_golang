CREATE TABLE products (
    id serial primary key,
    name varchar,
    description TEXT,
    price decimal,
    quantity integer
);

INSERT INTO products (name, description, price, quantity) values 
('T-shirt', 'Black', 19.90, 10),
('Headphone', 'Better than avarage', 99.00, 5);