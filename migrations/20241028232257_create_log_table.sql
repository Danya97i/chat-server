-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chats_logs (
    id BIGSERIAL PRIMARY KEY,
    action TEXT,
    chat_id BIGINT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats_logs;
-- +goose StatementEnd
