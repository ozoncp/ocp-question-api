-- +goose Up
-- +goose StatementBegin
CREATE TABLE questions (
    id BIGSERIAL,
    user_id BIGINT NOT NULL,
    text TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE questions;
-- +goose StatementEnd
