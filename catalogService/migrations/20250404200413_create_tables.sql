-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS products
(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    description TEXT,
    price DECIMAL NOT NULL,
    category_id INTEGER REFERENCES categories(id) NOT NULL,
    created_at TIMESTAMP NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
