-- +goose Up
-- +goose StatementBegin
CREATE TABLE consumptions (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  meter_id int,
  active FLOAT,
  reactive_inductive FLOAT,
  reactive_capacitive FLOAT,
  exported FLOAT,
  timestamp TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE consumptions;
-- +goose StatementEnd
