-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS statuses
(
    id SERIAL PRIMARY KEY,
    name varchar(20) unique not null
);

INSERT INTO statuses (name)
VALUES ('CREATED'), ('SHIPPED'), ('DELIVERED');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS statuses;
-- +goose StatementEnd
