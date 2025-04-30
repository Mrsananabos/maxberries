-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS statuses
(
    id SERIAL PRIMARY KEY,
    name varchar(20) unique not null
);

CREATE TABLE IF NOT EXISTS orders
(
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    total_items_price decimal not null,
    delivery_price decimal not null,
    total_price decimal,
    currency text not null,
    distance int,
    status_id INTEGER REFERENCES statuses(id) not NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items
(
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE cascade not NULL,
    product_id int NOT NULL,
    quantity int not null,
    unit_price decimal not null
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS statuses;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
