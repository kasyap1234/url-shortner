-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS url_map (
    short_code TEXT PRIMARY KEY,
    url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS url_map;
-- +goose StatementEnd
