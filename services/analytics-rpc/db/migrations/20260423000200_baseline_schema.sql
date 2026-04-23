-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS analytics (
    id UUID PRIMARY KEY,
    url_id UUID NOT NULL,
    user_id UUID NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS analytics;
-- +goose StatementEnd
