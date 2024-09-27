-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.bank_storage(
        user_id INT NOT NULL UNIQUE,
        balance INT not null ,
        date_created date not null ,
        date_updated date not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists public.bank_storage;
-- +goose StatementEnd




