-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    occupation VARCHAR(100) DEFAULT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_file_name VARCHAR(255) DEFAULT NULL,
    role VARCHAR(255) DEFAULT NULL,
    token VARCHAR(255) DEFAULT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
