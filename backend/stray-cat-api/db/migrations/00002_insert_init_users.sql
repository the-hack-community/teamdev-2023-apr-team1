-- +goose Up
INSERT INTO users (user_id, name, email) VALUES ('123', '初期ユーザ', 'initial.user@example.com');
INSERT INTO users (user_id, name, email) VALUES ('456', '山田太郎', 'yamada.taro@example.com');

-- +goose Down
DELETE FROM users WHERE user_id = '123';
DELETE FROM users WHERE user_id = '456';
