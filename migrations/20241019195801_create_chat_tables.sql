-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chats (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    created_at timestamptz default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chat_users (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT REFERENCES chats(id) ON DELETE CASCADE,
    user_email varchar(255) NOT NULL,
    CREATED_AT timestamptz default CURRENT_TIMESTAMP,
    CONSTRAINT chat_users_chat_id_user_email_unique_key UNIQUE(chat_id, user_email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chat_users;
DROP TABLE chats;
-- +goose StatementEnd
