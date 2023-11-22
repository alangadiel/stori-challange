CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    external_id INT UNIQUE NOT NULL,
    date DATE NOT NULL,
    amount NUMERIC(10,2) NOT NULL
);