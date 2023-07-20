-- +goose Up
-- +goose StatementBegin
CREATE TABLE meters (
  id SERIAL PRIMARY KEY,
  address VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE meters;
-- +goose StatementEnd
