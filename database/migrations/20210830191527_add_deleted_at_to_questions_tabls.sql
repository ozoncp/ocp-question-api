-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.questions
    ADD COLUMN deleted_at TIMESTAMP WITHOUT TIME ZONE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.questions
    DROP COLUMN deleted_at;
-- +goose StatementEnd
