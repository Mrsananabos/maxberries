-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS delivery_tariff
(
    max_distance int not null,
    price decimal not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS delivery_tariff;
-- +goose StatementEnd
